package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	deliveryChan := make(chan kafka.Event)
	producer := NewKafkaProducer()
	// setting the key will ensure that the message is delivered to the same partition
	Publish("transfer sent", "mytest", producer, []byte("transfer"), deliveryChan)

	go DeliveryReport(deliveryChan)

	fmt.Println("Delivery report channel created")
	producer.Flush(5000)

	// wait for delivery report
	// e := <-deliveryChan
	// msg := e.(*kafka.Message)
	// if msg.TopicPartition.Error != nil {
	// 	fmt.Println("Delivery failed: %v\n", msg.TopicPartition.Error)
	// } else {
	// 	fmt.Println("Message sent: ", msg.TopicPartition)
	// }

	// producer.Flush(1000)
}

func NewKafkaProducer() *kafka.Producer {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers":   "gokafka_kafka:9092",
		"delivery.timeout.ms": "0",    // 0 means infinite timeout
		"acks":                "all",  // ensure that the message is written to all replicas before acknowledging the producer
		"enable.idempotence":  "true", // ensure that the message is written only once
	}

	p, err := kafka.NewProducer(configMap)
	if err != nil {
		log.Println(err.Error())
	}
	return p
}

func Publish(msg string, topic string, producer *kafka.Producer, key []byte, deliveryChan chan kafka.Event) error {
	message := &kafka.Message{
		Value:          []byte(msg),
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            key,
	}
	err := producer.Produce(message, deliveryChan)
	if err != nil {
		return err
	}
	return nil
}

func DeliveryReport(deliveryChan chan kafka.Event) {
	for e := range deliveryChan {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Println("Delivery failed: %v\n", ev.TopicPartition.Error)
			} else {
				fmt.Println("Message sent: ", ev.TopicPartition)
			}
		}
	}
}
