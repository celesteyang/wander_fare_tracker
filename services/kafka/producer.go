package kafka

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

func main() {
	Producer()
}

func Producer() {
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: "test-topic",
		Value: sarama.StringEncoder("Hello, World from Kafka Producer!"),
	}

	_, _, err = producer.SendMessage(msg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Message sent to Kafka")
}
