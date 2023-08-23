package main

import (
	"fmt"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type OrderProducer struct {
	producer        *kafka.Producer
	topic           string
	deliveryChannel chan kafka.Event
}

func NewOrderProducer(p *kafka.Producer, topic string) *OrderProducer {
	return &OrderProducer{
		producer:        p,
		topic:           topic,
		deliveryChannel: make(chan kafka.Event, 10000),
	}
}

func (op *OrderProducer) placeOrder(orderType string, size int) error {
	var (
		format  = fmt.Sprintf("%s - %d", orderType, size)
		payload = []byte(format)
	)

	fmt.Println("Make order: ", format)

	err := op.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &op.topic, Partition: kafka.PartitionAny},
		Value:          payload,
	},
		op.deliveryChannel,
	)
	if err != nil {
		log.Fatal(err)
	}

	<-op.deliveryChannel

	return nil
}

func main() {
	var (
		host        = "localhost:9092"
		id          = "go-producer"
		topic       = "MoqTopic"
		orderNumber = 1
	)

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": host,
		"client.id":         id,
		"acks":              "all",
	})
	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
	}

	defer p.Close()

	op := NewOrderProducer(p, topic)
	for {
		if err := op.placeOrder("BUY", orderNumber); err != nil {
			log.Fatal(err)
		}

		orderNumber++

		time.Sleep(time.Second * 3)
	}
}
