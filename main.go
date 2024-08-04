package main

import (
	"math/rand"
	"time"
)

// Data structure for simulated data
type Data struct {
	Timestamp int64 `json:"timestamp"`
	Value     int   `json:"value"`
}

// Simulate data integration
func generateData(dataChan chan Data) {
	for {
		data := Data{
			Timestamp: time.Now().UnixNano() / int64(time.Millisecond),
			Value:     rand.Intn(100),
		}
		dataChan <- data
		time.Sleep(1000 * time.Millisecond)
	}
}

func main() {
	dataChan := make(chan Data)
	topic := "my-topic"
	brokers := []string{"localhost:9092"}
	// Generate Data
	go generateData(dataChan)
	// Start Producer
	go produce(dataChan, topic, brokers)
	// Start Consumer
	consume(topic, brokers)
}
