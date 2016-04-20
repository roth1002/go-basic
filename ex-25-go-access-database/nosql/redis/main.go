package main

import (
	"fmt"
	"gopkg.in/redis.v4"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "go-redis", // no password set
		DB:       0,          // use default DB
	})
	//字符串操作
	client.Set("a", []byte("hello"), 0)
	val := client.Get("a")
	fmt.Println(val)
	client.Del("a")

	//list操作
	vals := []string{"a", "b", "c", "d", "e"}
	for _, v := range vals {
		client.RPush("l", v)
	}
	dbvals := client.LRange("l", 0, 5)
	fmt.Println(dbvals.Result())
	// for i, v := range dbvals {
	// 	println(i, ":", string(v))
	// }
	client.Del("l")
}
