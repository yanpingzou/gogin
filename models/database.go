package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func InitDB() {
	db, err = sql.Open("mysql",
		"h2_user:tdfs.123@tcp(172.16.18.201:3306)/h2_gin")
	if err != nil {
		panic(err.Error())
		return
	}
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1000)
}
