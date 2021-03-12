package main

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

func main() {

	s := fmt.Sprintf("%x", sha1.Sum([]byte(`return redis.call('get','foo')`)))
	fmt.Println(s)

	return
	pp := redis.Pool{
		MaxIdle:     1,
		MaxActive:   1,
		IdleTimeout: time.Second * 1,
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp", "localhost:6379", redis.DialPassword(""), redis.DialDatabase(0))
		},
	}

	conn := pp.Get()

	//res, err := redis.String(conn.Do("eval", "return redis.call('set',KEYS[1],'value')", 1, "eval"))

	res, err := redis.Strings(conn.Do("eval", "return {KEYS[1],KEYS[2],ARGV[1],ARGV[2]}", 2, "key1", "key2", "first", "second"))
	fmt.Println(res, err)

}
