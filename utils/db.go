package utils

import (
    "fmt"
    "os"
    "time"
    "github.com/jmoiron/sqlx"
)

var conn *sqlx.DB

func DB() *sqlx.DB {
    if conn == nil {
        var db *sqlx.DB
        var err error

        // We set up a loop which will attempt to connect to the database.
        // This way we can bring up the authorisation service, even if there
        // are connection/server issues with the database.
        for connected := false; connected != true; db, err = createDbConnection() {
            if err == nil {
                connected = true
            } else {
                time.Sleep(time.Minute / 10)
            }
        }

        //    defer db.Close()

        conn = db
    }

    return conn
}

// Creates a database connection and check connectivity in one.
func createDbConnection() (*sqlx.DB, error) {

    host := os.Getenv("MYSQL_HOSTNAME")
    user := os.Getenv("MYSQL_USERNAME")
    pass := os.Getenv("MYSQL_PASSWORD")
    port := os.Getenv("MYSQL_PORT")

    return sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, pass, host, port , "authorisation"))
}
