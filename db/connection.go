package db

import (
    "fmt"
    "os"
    "sync"

    _ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
)


var connections = make(map[string]*sqlx.DB, 0)
var mutex sync.Mutex

// CreateDbConnection a db connection
func CreateDbConnection(dbname string) (*sqlx.DB, error) {
    mutex.Lock()
    defer mutex.Unlock()
    if _, ok := connections[dbname]; ok != false {
        fmt.Println("Have a connection will return the same one.....")
        return connections[dbname], nil
    }
    mysqlHost := os.Getenv("MYSQL_HOSTNAME")
    mysqlUser := os.Getenv("MYSQL_USERNAME")
    mysqlPass := os.Getenv("MYSQL_PASSWORD")
    conn, err := sqlx.Connect(
        "mysql",
        fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", mysqlUser, mysqlPass, mysqlHost, 3306, dbname),
    )
    connections[dbname] = conn
    return conn, err
}

// MustCreateDbConnection creates a db connection or panics.....
func MustCreateDbConnection(dbname string) *sqlx.DB {
    conn, err := CreateDbConnection(dbname)
    if err != nil {
        panic(err)
    }
    return conn
}