package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

// go get -u github.com/go-redis/redis

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	fmt.Println(client.Get("aa").Int64())

	client.GetSet("b", 1)

	fmt.Println(client.Get("b").Result())

	//time.ParseDuration("10m")

	client.Set("ta",1,100 * time.Second)

	fmt.Println(client.TTL("ta"))
	fmt.Println(client.Get("ta"))

	client.RPush("l", "a")
	client.RPush("l", "b")

	fmt.Println(client.LRange("l", 0, 1).Result())
	fmt.Println(client.LLen("l").Result())
}
