package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println(err)
		return
	}

	c.Send("blpop", "aaaa", "50")
	c.Flush()
	c.Receive()
	defer c.Close()
}
