package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {

	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.210.228:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	//	fmt.Println(client)

	err = client.Set("key", "ssssssss", 0).Err()
	if err != nil {
		panic(err)
	}

	val, _ := client.Get("key").Result()

	fmt.Println("key", val)

	// Output: key value
	// key2 does not exists
}
