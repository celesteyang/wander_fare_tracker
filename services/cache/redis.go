package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

func main() {
	// Initialize a new Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Default Redis address
	})

	// Test the connection
	ctx := context.Background()
	err := rdb.Ping(ctx).Err()
	if err != nil {
		log.Fatalf("could not connect to Redis: %v", err)
	}

	fmt.Println("Successfully connected to Redis!")

	// Set a key-value pair in Redis
	err = rdb.Set(ctx, "hello", "world", 0).Err()
	if err != nil {
		log.Fatalf("could not set key in Redis: %v", err)
	}

	// Retrieve the value from Redis
	val, err := rdb.Get(ctx, "hello").Result()
	if err != nil {
		log.Fatalf("could not get key from Redis: %v", err)
	}

	fmt.Printf("The value of 'hello' is: %s\n", val)
}
