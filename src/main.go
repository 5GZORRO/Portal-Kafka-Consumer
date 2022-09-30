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
