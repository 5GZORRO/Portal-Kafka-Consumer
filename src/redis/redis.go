package redis

import (
	"context"
	"log"
	"os"

	redis "github.com/go-redis/redis/v9"
)

var ctx = context.TODO()

func connect() (*redis.Client, error) {
	// Fetching env vars
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	password := "" // No password      -> Add this if needed
	db := 0        // Default Database -> Add this if needed

	// Creates instance of new Redis client
	client := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       db,
	})

	// Checks if there was an error
	if err := client.Ping(ctx).Err(); err != nil {
		log.Printf("Error connecting to Redis: %s", err)
		return nil, err
	}
	//log.Println("Connected to Redis")
	return client, nil
}

// Method that stores a key, value pair from Redis
func Set(key, value string) error {

	// Connect to Redis
	db, err := connect()
	if err != nil {
		os.Exit(1)
	}

	log.Printf("Storing k,v: %s, %s", key, value)
	err = db.Set(ctx, key, value, 0).Err()
	if err != nil {
		log.Printf("Error setting k,v in redis: %s, %s ", key, value)
		return err
	}

	// Close client
	db.Close()
	return nil
}

// Method that fetches value based on a key from Redis
func Get(key string) (string, error) {
	// Connect to Redis
	db, err := connect()
	if err != nil {
		os.Exit(1)
	}
	//log.Printf("Fetching value from key: %s", key)
	value, err := db.Get(ctx, key).Result()
	if err != nil {
		log.Printf("Error getting k in redis: %s", key)
		return "", err
	}

	// Close client
	db.Close()
	return value, nil
}

// Method that fetches all keys
func GetAllKeys() ([]string, error) {

	// Connect to Redis
	db, err := connect()
	if err != nil {
		os.Exit(1)
	}
	//log.Printf("Fetching all key, values")

	// Get all keys
	keys, _, err := db.Scan(ctx, 0, "*", 0).Result()
	if err != nil {
		log.Printf("Error getting all keys from Redis")
		return nil, err
	}

	// Close client
	db.Close()
	return keys, nil
}

// Method that fetches value based on a key from Redis
func Delete(key string) error {
	// Connect to Redis
	db, err := connect()
	if err != nil {
		os.Exit(1)
	}

	//log.Printf("Deleting key from Redis: %s", key)
	err = db.Del(ctx, key).Err()
	if err != nil {
		log.Printf("Error deleting key in redis: %s", key)
		return err
	}

	// Close client
	db.Close()
	return nil
}
