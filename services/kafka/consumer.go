// services/kafka/consumer.go
package kafka

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

func StartConsumer() {
	ConsumeMessages()
}

func ConsumeMessages() {
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition("test-topic", 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatal(err)
	}
	defer partitionConsumer.Close()

	for msg := range partitionConsumer.Messages() {
		fmt.Printf("Message received: %s\n", string(msg.Value))
	}
}
