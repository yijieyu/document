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

	//fmt.Println(redis.String(conn.Do("set", "mykey", "value_10", "ex", "100")))
	//fmt.Println(redis.String(conn.Do("get", "key1")))
	//res, err := redis.StringMap(conn.Do("blpop", "lkey1", 0))
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(res["lkey1"])

	// 事件提醒

	//go p(pp, 1)
	//go p(pp, 2)
	go p(pp, 3)
	go c(pp)
	for {
		select {}
	}

}

func p(p redis.Pool, index int) {

	conn := p.Get()

	defer func() {
		conn.Close()
		if err := recover(); err != nil {
			fmt.Println("p panic", err)
			return
		}
	}()

	for {
		err := conn.Send("multi")
		if err != nil {
			panic(err)
		}

		err = conn.Send("sadd", "skey", fmt.Sprintf("%d_sadd_key_%s", index, time.Now().String()))
		if err != nil {
			panic(err)
		}

		err = conn.Send("rpush", "help_skey_list", "wait")
		if err != nil {
			panic(err)
		}

		_, err = conn.Do("exec")
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Second * time.Duration(index))
	}

}
func c(p redis.Pool) {
	conn := p.Get()

	defer func() {
		conn.Close()
		if err := recover(); err != nil {
			fmt.Println("c panic:", err)
			return
		}
	}()

	//fn := func() string {
	//	conn := p.Get()
	//
	//	defer func() {
	//		conn.Close()
	//	}()
	//
	//	fmt.Println(1111)
	//	s, err := redis.StringMap(conn.Do("blpop", "help_skey_list", 0))
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Println(s)
	//
	//	res, err := redis.String(conn.Do("spop", "skey"))
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	return res
	//}

	for {

		res, err := redis.String(conn.Do("spop", "skey"))
		if err != redis.ErrNil && err != nil {
			panic(err)
		}
		if res != "" {
			fmt.Println("res", res)
			continue
		}

		_, err = redis.StringMap(conn.Do("blpop", "help_skey_list", 0))

		if err != redis.ErrNil && err != nil {
			panic(err)
		}

	}
}
