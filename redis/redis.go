package main

import (
	"encoding/json"
	"fmt"
	"runtime"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/gogmod/rejson"
)

func main() {
	//testStream()
	runtime.GOMAXPROCS(1)
	pool := getClient()
	fmt.Println(pool)
	pool2 := getClient()

	for i := 100; i < 110; i++ {
		var str = strconv.Itoa(i)
		var p = getClient()
		go func(str string) {
			fmt.Println(str)
			id, err := p.XAdd(&redis.XAddArgs{
				Stream: "stream",
				ID:     str,
				Values: map[string]interface{}{"name": "gary" + str},
			}).Result()
			fmt.Println(err)
			fmt.Println(id)
		}(str)
	}
	msgs, _ := client.XRangeN("stream", "-", "+", 1).Result()
	fmt.Println(msgs)
	stats := pool.PoolStats()
	stats2 := pool2.PoolStats()
	fmt.Printf("status: %+v\n", stats)
	fmt.Printf("status2: %+v\n", stats2)
	testRejson()
	select {}
}

//Customer ....
type Customer struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedAt int    `json:"created_at"`
	UpdatedAt int    `json:"updated_at"`
}

func testRejson() {
	client := rejson.ExtendClient(getClient())
	defer client.Close()
	customer := &Customer{Username: "xingcuntian", Password: "xingcutian", CreatedAt: 1542357972, UpdatedAt: 1542357972}
	jsons, err := json.Marshal(customer)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("source: ", string(jsons))

	client.JSONSet("customerList", ".", string(jsons))
	jsonStrings, err := client.JSONGet("customerList").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(jsonStrings)

	json := `{"name":"Leonard Cohen","lastSeen":1478476800,"loggedOut": false}`
	client.JSONSet("userList", ".", json)
	jsonString, err := client.JSONGet("userList").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("====>", jsonString)
}
