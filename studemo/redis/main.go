package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func connection() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:"127.0.0.1:6379",
	})

	// 通过 cient.Ping() 来检查是否成功连接到了 redis 服务器
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	return client
}

func main()  {
	client := connection()
	cmd := client.Get("TEST")
	fmt.Println(cmd)


}