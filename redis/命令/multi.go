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

	conn.Send("multi")
	conn.Send("incr", "foo")
	conn.Send("incr", "bar")

	fmt.Println(conn.Do("DISCARD"))
	conn.Send("incr", "foo1")
	res, err := redis.Ints(conn.Do("exec"))
	if err != nil {
		panic(err)
	}

	fmt.Println(res[0], res[1], res[2])

}
