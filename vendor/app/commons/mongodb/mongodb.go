package mongodb

import (
    "crypto/rand"
    "fmt"
    "log"
    "gopkg.in/mgo.v2"
)

var (
    Session *mgo.Session
    // Database info
    databases DialInfo
)

type DialInfo struct {
    // Addrs holds the addresses for the seed servers.
    Addrs string

    // Database is the default database name used when the Session.DB method
    // is called with an empty name, and is also used during the initial
    // authentication if Source is unset.
    Database string

    // Username and Password inform the credentials for the initial authentication
    // done on the database defined by the Source field. See Session.Login.
    Username string
    Password string
}

// Connect to the database
func Connect(d DialInfo) {
    var err error

    // Store the config
    databases = d

    Session, err := mgo.Dial(DSN(d))

    if err != nil {
        log.Println("Failed to connect mongodb! ", err)
    }

    defer Session.Close()
    Session.SetMode(mgo.Monotonic, true)

}

// ReadConfig returns the database information
func ReadConfig() DialInfo {
    return databases
}

func GetSession() (*mgo.Session, error) {
    //Establish our database connection
    if Session == nil {
        var err error
        Session, err = mgo.Dial(DSN(databases))
        if err != nil {
            return nil, err
        }

        //Optional. Switch the session to a monotonic behavior.
        Session.SetMode(mgo.Monotonic, true)
    }

    return Session.Copy(), nil
}

// DSN returns the Data Source Name
func DSN(ci DialInfo) string {
    // Example: mongodb://myuser:mypass@localhost:40001,otherhost:40001/mydb
    return "mongodb://" +
    ci.Username +
    ":" +
    ci.Password +
    "@" +
    ci.Addrs +
    "/" +
    ci.Database
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