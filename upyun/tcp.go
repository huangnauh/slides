package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	// START OMIT
	client := redis.NewClient(&redis.Options{
		Addr: "10.0.0.193:31000",
	})

	err := client.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	// END OMIT
	fmt.Println("key", val)
}
