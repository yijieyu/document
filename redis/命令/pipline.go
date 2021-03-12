package main

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

func main() {
	pp := redis.Pool{
		MaxIdle:     1,
		MaxActive:   1,
		IdleTimeout: time.Second * 1,
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp", "localhost:6379", redis.DialPassword(""), redis.DialDatabase(0))
		},
	}

	conn := pp.Get()

	conn.Send("set", "foo", "bar")
	conn.Send("get", "foo")
	conn.Send("get", "foo")

	conn.Flush()
	fmt.Println(conn.Receive())
	fmt.Println(redis.String(conn.Receive()))
	fmt.Println(redis.String(conn.Receive()))
}
