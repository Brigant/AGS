package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	var (
		host  = "localhost:9092"
		id    = "sales-manager"
		topic = "MoqTopic"
	)

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": host,
		"group.id":          id,
		"auto.offset.reset": "smallest",
	})
	if err != nil {
		log.Fatal(err)
	}

	err = consumer.Subscribe(topic, nil)
	if err != nil {
		log.Fatal(err)
	}

	for {
		ev := consumer.Poll(100)
		switch e := ev.(type) {
		case *kafka.Message:
			fmt.Printf("Sales manager get order: %s for handeling\n", string(e.Value))
		case *kafka.Error:
			fmt.Printf("Message queue errorr: %s\n", e.Error())
		}
	}
}
