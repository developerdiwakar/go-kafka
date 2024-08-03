package main

import (
	"log"

	"github.com/IBM/sarama"
)

func consume(topic string, brokers []string) error {
	config := sarama.NewConfig()
	// If enabled, any errors that occurred while consuming are returned on the Errors channel (default disabled).
	config.Consumer.Return.Errors = true
	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		return err
	}

	defer func() {
		if err := consumer.Close(); err != nil {
			log.Println("Failed to close consumer", err)
		}
	}()

	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)

	if err != nil {
		return err
	}

	defer partitionConsumer.Close()

	for msg := range partitionConsumer.Messages() {
		log.Printf("Received message: key=%s value=%s\n ", string(msg.Key), string(msg.Value))
		// Process the data here (e.g., store in database, send to another system)
	}

	return nil

}
