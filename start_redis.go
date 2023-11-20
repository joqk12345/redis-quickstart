package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	// Create a new Redis client
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // Redis password (if any)
		DB:       0,                // Redis database number
	})

	// Ping the Redis server to check the connection
	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println("Failed to connect to Redis:", err)
		return
	}
	fmt.Println("Connected to Redis:", pong)

	// Set a key-value pair in Redis
	err = client.Set("mykey", "myvalue", 0).Err()
	if err != nil {
		fmt.Println("Failed to set key:", err)
		return
	}
	fmt.Println("Key set successfully")

	// Get the value for a key from Redis
	value, err := client.Get("mykey").Result()
	if err != nil {
		fmt.Println("Failed to get value for key:", err)
		return
	}
	fmt.Println("Value:", value)

	// Increment a value in Redis
	err = client.Incr("counter").Err()
	if err != nil {
		fmt.Println("Failed to increment counter:", err)
		return
	}
	fmt.Println("Counter incremented")

	// Retrieve the incremented value
	counter, err := client.Get("counter").Result()
	if err != nil {
		fmt.Println("Failed to get counter value:", err)
		return
	}
	fmt.Println("Counter value:", counter)

	// Close the Redis client connection
	err = client.Close()
	if err != nil {
		fmt.Println("Failed to close Redis client:", err)
		return
	}
	fmt.Println("Redis client closed")
}
