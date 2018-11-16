package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func testStream() {

	client := getClient()

	id, err := client.XAdd(&redis.XAddArgs{
		Stream: "stream",
		ID:     "1-0",
		Values: map[string]interface{}{"name": "gary"},
	}).Result()

	if err != nil {
		panic(err)
	}

	fmt.Println(id)

	id, err = client.XAdd(&redis.XAddArgs{
		Stream: "stream",
		ID:     "2-0",
		Values: map[string]interface{}{"name": "hary"},
	}).Result()

	if err != nil {
		panic(err)
	}
	fmt.Println(id)

	id, err = client.XAdd(&redis.XAddArgs{
		Stream: "stream",
		ID:     "3-0",
		Values: map[string]interface{}{"name": "pary"},
	}).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(id)
}
