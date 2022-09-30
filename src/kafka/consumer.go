package kafka

import (
	"context"
	"fmt"
	"log"
	"time"

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

// To be removed -> Function that just pushes a boilerplate violation to Kafka to test application
func (kfa *Kafka) Test() {

	time.Sleep(10 * time.Second)

	conn, err := kafka.DialLeader(context.Background(), "tcp", kfa.Host, kfa.Topic, 0)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte("{ \"id\": \"9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d\", \"productID\": \"9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d\", \"transactionID\": \"9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d\",  \"sla\": { \"id\": \"9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d\",  \"href\": \"http://www.acme.com/slaManagement/sla/123444\" }, \"rule\": { \"id\": \"availability\", \"metric\": \"http://www.provider.com/metrics/availability\", \"unit\": \"%\", \"referenceValue\": \"99.95\", \"operator\": \".ge\", \"tolerance\": \"0.05\", \"consequence\": \"http://www.provider.com/contract/claus/30\" }, \"violation\": { \"actualValue\": \"90.0\" } }")},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
