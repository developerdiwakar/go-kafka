package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/IBM/sarama"
)

// Data structure for incoming data
type IncomingData struct {
	TimeStamp int64  `json:"timestamp"`
	Value     string `json:"value"`
}

// Data structure for processed data (optional)
type ProcessedData struct {
	DeliveredAt       time.Time   `json:"deliveredAt"`
	DeliveredAtString string      `json:"deliveredAtString"`
	Response          interface{} `json:"response"`
}

func consume(topic string, brokers []string) <-chan IncomingData {
	dataChan := make(chan IncomingData)
	go func() {
		// Create new Kafka configuration instance
		config := sarama.NewConfig()
		config.Consumer.Return.Errors = true
		// Create new Consumer
		consumer, err := sarama.NewConsumer(brokers, config)
		if err != nil {
			log.Fatal(err)
		}
		// Close consumer at the end
		defer func() {
			if err := consumer.Close(); err != nil {
				log.Println("Failed to close consumer", err)
			}
		}()
		// Consume from topic partitions
		consume_0, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
		if err != nil {
			log.Fatalln("Consuming error:", err)
		}
		defer consume_0.Close()

		// Receiving Messages
		for msg := range consume_0.Messages() {
			var incomingData IncomingData
			if err := json.Unmarshal(msg.Value, &incomingData); err != nil {
				log.Println("Error Unmarshalling data", err)
				continue
			}
			dataChan <- incomingData
		}
		close(dataChan)
	}()
	return dataChan
}

func processData(dataChan <-chan IncomingData) {
	var pd ProcessedData
	for data := range dataChan {
		// Perform computations on the data
		// fmt.Printf("Processing data: %+v\n", data)
		pd.DeliveredAt = time.Unix(0, data.TimeStamp)
		pd.DeliveredAtString = pd.DeliveredAt.Format("2006-01-02 15:04:05") // 24 hour fomat
		pd.Response = data

		fmt.Printf("Processed Data: %+v\n\n", pd)
		// Optionally, produce new data to another topic
		// ...
	}
}

func main() {
	// Receiving from the producer from Real Time Data Pipeline
	topic := "my-topic"
	brokers := []string{"localhost:9092"}

	dataChan := consume(topic, brokers)
	processData(dataChan)
}
