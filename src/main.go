package main

import (
	"log"
	"net/http"
	"os"
	"portal-kafka-consumer/api"
	"portal-kafka-consumer/kafka"
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

	// Starts the server
	http.HandleFunc("/", api.GetViolations)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
