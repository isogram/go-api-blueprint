package main

import (
    "log"
    "runtime"
    "fmt"

    "github.com/spf13/viper"

    "app/controllers"
    "app/routes"
    "app/commons/email"
    "app/commons/mysql"
    "app/commons/mongodb"
    "app/commons/server"
)

func init() {
    // Verbose logging with file name and line number
    log.SetFlags(log.Lshortfile)

    // Use all CPU cores
    runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {

    // Load configuration file
    fmt.Println("\n==========================")
    fmt.Println("Loading configuration...")

    // Load configuration file
    viper.SetConfigName("config")
    viper.AddConfigPath("config")
    err := viper.ReadInConfig()
    if err != nil {
        panic(fmt.Errorf("Fatal error config file: %s \n", err))
    }

    // Load email service
    ConfigureEmail()

    // Load MySQL service
    ConfigureMySQL()

    // Load controller routes
    controllers.Load()

    // Load MongoDB service
    ConfigureMongoDB()

    // Load Server
    ConfigureServer()

}

// Configure email
func ConfigureEmail() {
    c := email.SMTPInfo{
        Host            : viper.GetString("email.host"),
        Port            : viper.GetInt("email.port"),
        FromName        : viper.GetString("email.from_name"),
        FromAddress     : viper.GetString("email.from_address"),
        Encryption      : viper.GetString("email.encryption"),
        Username        : viper.GetString("email.username"),
        Password        : viper.GetString("email.Password"),
    }

    email.Configure(c)
    fmt.Println("Email configured...")
}

// Configure MySQl
func ConfigureMySQL() {
    c := mysql.MySQLInfo{
        Username  : viper.GetString("database.mysql.username"),
        Password  : viper.GetString("database.mysql.password"),
        Name      : viper.GetString("database.mysql.database"),
        Hostname  : viper.GetString("database.mysql.host"),
        Port      : viper.GetInt("database.mysql.port"),
        Parameter : viper.GetString("database.mysql.parameter"),
    }

    mysql.Connect(c)
    fmt.Println("MySQL configured...")
}

// Configure MongoDB
func ConfigureMongoDB() {
    c := mongodb.DialInfo{
        Addrs       : viper.GetString("database.mongodb.host") + ":" + viper.GetString("database.mongodb.port"),
        Database    : viper.GetString("database.mongodb.database"),
        Username    : viper.GetString("database.mongodb.username"),
        Password    : viper.GetString("database.mongodb.password"),
    }

    mongodb.Connect(c)
    fmt.Println("MongoDB configured...")
}

// Configure Server
func ConfigureServer() {
    c := server.Server{
        Hostname  : viper.GetString("server.host"),
        UseHTTP   : viper.GetBool("server.use_http"),
        UseHTTPS  : viper.GetBool("server.use_https"),
        HTTPPort  : viper.GetInt("server.http_port"),
        HTTPSPort : viper.GetInt("server.https_port"),
        CertFile  : viper.GetString("server.cert_file"),
        KeyFile   : viper.GetString("server.key_file"),
    }

    fmt.Println("Server configured ...")
    fmt.Println("==========================")

    server.Run(routes.LoadHTTP(), routes.LoadHTTPS(), c)
}