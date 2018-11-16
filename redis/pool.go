package main

import (
	"time"

	"github.com/go-redis/redis"
)

var client *redis.Client
var redisAddr = "127.0.0.1:6379"
var password = "xingcuntian"

func getClient() *redis.Client {
	opt := redisOptions()
	if client == nil {
		client = redis.NewClient(opt)
	}
	pong, err := client.Ping().Result()
	if err != nil || pong != "PONG" {
		client = redis.NewClient(opt)
	}
	if pong == "PONG" {
		client = redis.NewClient(opt)
	}
	return client
}

func redisOptions() *redis.Options {
	return &redis.Options{
		Addr:               redisAddr,
		DB:                 15,
		Password:           password,
		DialTimeout:        10 * time.Second,
		ReadTimeout:        30 * time.Second,
		WriteTimeout:       30 * time.Second,
		PoolSize:           10,
		PoolTimeout:        30 * time.Second,
		IdleTimeout:        500 * time.Millisecond,
		IdleCheckFrequency: 500 * time.Millisecond,
	}
}
