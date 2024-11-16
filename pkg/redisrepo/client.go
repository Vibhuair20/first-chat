package redisrepo

//all the tools and the libraries
import (
	"context"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

// declaring a pointer as a global variable-> it is storing the redis connection
var redisClient *redis.Client

// function to connect to redis
func InitializeRedis() *redis.Client {
	//connecting to redis using all the credentials
	conn := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_CONNECTION_STRING"),
		Password: os.Getenv("REDIS_PASSWORDREDIS_PASSWORD"),
		DB:       0,
	})

	//now we test if the connection works
	succ, err := conn.Ping(context.Background()).Result()
	//if connection fails then it shows the error and stores it
	if err != nil {
		log.Fatal("reddis connection failed", err)
	}
	//if connection is successfull then it is stored as successfully
	log.Println("reddis is successfully connected", "Ping", succ)

	//saving the connection
	redisClient = conn
	return redisClient
}
