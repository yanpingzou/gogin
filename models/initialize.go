package models

import (
	"database/sql"
	"github.com/garyburd/redigo/redis"
	"gogin/dbs"
)

var (
	db  *sql.DB
	rds *redis.Pool
)

func InitConnect() {
	rds = dbs.InitRedis()
	db = dbs.InitDB()
}
