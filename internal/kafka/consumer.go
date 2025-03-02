package kafka

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)


func NewConsumer(broker, groupID, topic string) (*kafka.Consumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, err
	}

	err = c.Subscribe(topic, nil)
	if err != nil {
		return nil, err
	}

	return c, nil
}


func ConsumeMessages(consumer *kafka.Consumer) {
	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			log.Printf("Received message: %s from topic %s\n", string(msg.Value), *msg.TopicPartition.Topic)
		} else {
			log.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}
