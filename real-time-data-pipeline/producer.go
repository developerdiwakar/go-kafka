package main

import (
	"encoding/json"
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
		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Println("Error JSON Marshalling data", err)
			continue
		}
		// Prepare Prodcuer Message
		message := &sarama.ProducerMessage{
			Topic: topic,
			Key:   sarama.StringEncoder(fmt.Sprintf("%d", data.Timestamp)),
			Value: sarama.ByteEncoder(jsonData),
		}
		// Send Message to Broker
		partition, offset, err := producer.SendMessage(message)

		if err != nil {
			log.Println("Failed to send message", err)
		} else {
			log.Printf("Message sent to partition %d at offset %d\n", partition, offset)
		}
	}

	return nil
}
