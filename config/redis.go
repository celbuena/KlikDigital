package db

import (
	"github.com/gomodule/redigo/redis"
)

var GlobalCache redis.Conn

func initCache() {
	// Initialize the redis connection to a redis instance running on your local machine
	conn, err := redis.DialURL("redis://localhost")
	if err != nil {
		panic(err)
	}
	// Assign the connection to the package level `cache` variable
	GlobalCache = conn
}

func init()  {
	initCache()
}