package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func ConnectDB() error {
    var err error

    db, err = sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/databasename")
    if err != nil {
        return err
    }

    err = db.Ping()
    if err != nil {
        return err
    }

    fmt.Println("Database connected successfully")
    return nil
}

func GetDB() *sql.DB {
    return db
}