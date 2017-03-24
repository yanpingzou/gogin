package dbs

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

var (
	redis_host string
	db_num     int
)

func InitRedis() *redis.Pool {
	//使用方法
	//rds := Rds.Get()
	//defer rds.Close()
	//rds.Do("SET", "key", "value")
	redis_host = "172.16.18.201:6379"
	db_num = 0
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		MaxActive:   10,
		// Other pool configuration not shown in this example.
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redis_host)
			if err != nil {
				return nil, err
			}
			// 需要认证的redis连接方法
			//if _, err := c.Do("AUTH", password); err != nil {
			//	c.Close()
			//	return nil, err
			//}
			if _, err := c.Do("SELECT", db_num); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
	}
}
