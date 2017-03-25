package dbs

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func InitDB() *sql.DB {
	MysqlDB, err := sql.Open("mysql",
		"h2_user:tdfs.123@tcp(192.168.1.8:3306)/h2_gin?charset=utf8mb4&parseTime=true")
	if err != nil {
		panic(err)
	}
	err = MysqlDB.Ping()
	if err != nil {
		panic(err)
	}
	MysqlDB.SetMaxOpenConns(2000)
	MysqlDB.SetMaxIdleConns(1000)
	return MysqlDB
}
