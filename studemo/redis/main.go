package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

func connection() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})

	// 通过 cient.Ping() 来检查是否成功连接到了 redis 服务器
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	return client
}

func main() {
	client := connection()
	cmd := client.Get("TEST111")
	v, e := cmd.Result()

	fmt.Println(cmd)
	fmt.Println(cmd.Val())
	fmt.Println(v, e)
	client.Set("test1", "222", 200*time.Second)
	cmd2 := client.Get("test1")
	fmt.Println(cmd2)
	fmt.Println(cmd2.Result())
	fmt.Println(cmd2.Val())
}
