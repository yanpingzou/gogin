package models

import (
	"gogin/dbs"
)

var (
	db  = dbs.InitDB()
	rds = dbs.InitRedis()
)
