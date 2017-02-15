package mysql

import (
    "crypto/rand"
    "database/sql"
    "fmt"
    "log"

    _ "github.com/go-sql-driver/mysql" // MySQL driver
    "github.com/jmoiron/sqlx"
)

var (
    // SQL wrapper
    SQL *sqlx.DB
    // Database info
    databases MySQLInfo
)

// MySQLInfo is the details for the database connection
type MySQLInfo struct {
    Username  string
    Password  string
    Name      string
    Hostname  string
    Port      int
    Parameter string
}

// Connect to the database
func Connect(d MySQLInfo) {
    var err error

    // Store the config
    databases = d

    // Connect to MySQL
    if SQL, err = sqlx.Connect("mysql", DSN(d)); err != nil {
    log.Println("SQL Driver Error", err)
    }

    // Check if is alive
    if err = SQL.Ping(); err != nil {
    log.Println("Database Error", err)
    }
}

// ReadConfig returns the database information
func ReadConfig() MySQLInfo {
    return databases
}

// DSN returns the Data Source Name
func DSN(ci MySQLInfo) string {
    // Example: root:@tcp(localhost:3306)/test
    return ci.Username +
    ":" +
    ci.Password +
    "@tcp(" +
    ci.Hostname +
    ":" +
    fmt.Sprintf("%d", ci.Port) +
    ")/" +
    ci.Name + ci.Parameter
}

// AffectedRows returns the number of rows affected by the query
// Will panic if result does not exist
func AffectedRows(result sql.Result) int {
    // If successful, get the number of affected rows
    count, err := result.RowsAffected()
    if err != nil { // Feature not supported
    // Only show error for admin
    log.Println(err)
    return 1
    }

    return int(count)
}

// UUID generates UUID for use as an ID
func UUID() (string, error) {
    b := make([]byte, 16)
    _, err := rand.Read(b)
    if err != nil {
    return "", err
    }

    return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:]), nil
}
