package main

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

// Kafka-producer
func produce(dataChan chan Data, topic string, brokers []string) error {
	config := sarama.NewConfig()
	// If enabled, successfully delivered messages will be returned on the Successes channel (default disabled).
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return err
	}

	defer func() {
		if err := producer.Close(); err != nil {
			log.Println("Failed to close producer", err)
		}
	}()

	for data := range dataChan {
		message := &sarama.ProducerMessage{
			Topic: topic,
			Key:   sarama.StringEncoder(fmt.Sprintf("%d", data.Timestamp)),
			Value: sarama.StringEncoder(fmt.Sprintf("%d", data.Value)),
		}
		partition, offset, err := producer.SendMessage(message)

		if err != nil {
			log.Println("Failed to send message", err)
		} else {
			log.Printf("Message sent to partition %d at offset %d\n", partition, offset)
		}
	}

	return nil
}
