package redisrepo

import (
	"context"
	"encoding/json"
	"log"

	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

func main() {
	// Initialize Redis client
	redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
	})

	// Prepare your data to store (as a map)
	chatData := map[string]interface{}{
		"user":      "John",
		"message":   "Hello, World!",
		"timestamp": "2024-11-21T12:00:00Z",
	}

	// Convert chat data to JSON (byte slice)
	by, err := json.Marshal(chatData)
	if err != nil {
		log.Fatalf("Error marshaling data: %v", err)
	}

	// Set the JSON data in Redis
	chatKey := "chat:1234"
	res, err := redisClient.Do(
		context.Background(),
		"JSON.SET",
		chatKey,
		"$",
		string(by),
	).Result()

	if err != nil {
		log.Fatalf("Error setting JSON data in Redis: %v", err)
	}

	// Log result
	log.Printf("Response from Redis: %v", res)
}
