package kafka

import (
	"context"
	"fmt"
	"log"

	"portal-kafka-consumer/model"
	"portal-kafka-consumer/redis"
	"portal-kafka-consumer/utils"

	"github.com/segmentio/kafka-go"
)

type Kafka struct {
	Host  string
	Topic string
}

func CreateKafkaInstance(host, topic string) *Kafka {
	return &Kafka{
		Host:  host,
		Topic: topic,
	}
}

// Kafka Consume method that receives violations, validates them and stores them on the database
func (kfa *Kafka) Consume() {

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{kfa.Host},
		Topic:   kfa.Topic,
		GroupID: "none",
	})

	for {
		msg, err := r.ReadMessage(context.Background()) // the `ReadMessage` method blocks until we receive the next event
		if err != nil {
			log.Println("could not read message " + err.Error())
		}

		message := string(msg.Value)
		fmt.Println("Received Message: ", message)

		// Validate Message by converting to struct
		violation := &model.Violation{}
		err = utils.JsonToStruct(violation, message)
		if err != nil {
			log.Println("Error converting message to struct: ", err)
			continue
		}

		// Store in Database
		redis.Set(violation.ID, message)
	}
}
