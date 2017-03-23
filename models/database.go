package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB
var err error

func InitDB() {
	db, err = sql.Open("mysql",
		"h2_user:tdfs.123@tcp(192.168.1.8:3306)/h2_gin?charset=utf8mb4&parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1000)
}
