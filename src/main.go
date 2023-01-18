package main

import (
	"log"
	"net/http"
	"os"
	"portal-kafka-consumer/api"
	"portal-kafka-consumer/kafka"
	"portal-kafka-consumer/redis"
	"time"
)

func main() {
	// Fetches Kafka env variables
	host, hostPresent := os.LookupEnv("KAFKA_HOST")
	if !hostPresent {
		log.Println("Error occurred while fetching kafka host")
		return
	}
	topic, topicPresent := os.LookupEnv("KAFKA_TOPIC_IN")
	if !topicPresent {
		log.Println("Error occurred while fetching kafka topic")
		return
	}

	// Create a Kafka struct
	kafka := kafka.CreateKafkaInstance(host, topic)

	// Opens consumer
	go kafka.Consume()

	go automaticCleanup()

	// Fetches Server Port
	port, portPresent := os.LookupEnv("PORT")
	if !portPresent {
		log.Println("Error occurred while fetching Port")
		return
	}

	// Starts the server
	http.HandleFunc("/", api.GetViolations)
	log.Println("Server is Running on localhost:" + port) // Should be below but oh well
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Println("Error occured while creating Server" + err.Error())
		return
	}
}

// Method that executes every minute and clears stack of notifications
func automaticCleanup() {

	time.Sleep(1 * time.Minute)

	// Get All stored violations
	keys, _ := redis.GetAllKeys()

	for _, key := range keys {
		redis.Delete(key) // Delete that violation entry from stack
	}
}
