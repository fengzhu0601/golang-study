package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main()  {
	client := createClient()
	defer client.Close()
}

// 创建 redis 客户端
func createClient() *redis.Client {
	client := redis.NewClient( &redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		PoolSize: 5,
	} )

	pong, err := client.Ping().Result()
	fmt.Println( pong, err )

	return client
}
