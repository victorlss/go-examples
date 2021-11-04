package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
)

func main() {
	produce("my-topic", "Hello World")
}

func produce(topic string, message string) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"client.id":         "example",
		"acks":              "all"})

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}

	deliveryChan := make(chan kafka.Event, 10000)
	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message)},
		deliveryChan,
	)

	kafkaEvent := <-deliveryChan
	kafkaMessage := kafkaEvent.(*kafka.Message)

	if kafkaMessage.TopicPartition.Error != nil {
		fmt.Printf("Delivery failed: %v\n", kafkaMessage.TopicPartition.Error)
	} else {
		fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
			*kafkaMessage.TopicPartition.Topic, kafkaMessage.TopicPartition.Partition, kafkaMessage.TopicPartition.Offset)
	}

	close(deliveryChan)
}
