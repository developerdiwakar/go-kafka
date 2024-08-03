package main

import (
	"math/rand"
	"time"
)

// Data structure for simulated data
type Data struct {
	Timestamp int64
	Value     int
}

// Simulate data integration
func generateData() chan Data {
	dataChan := make(chan Data)

	go func() {
		for {
			data := Data{
				Timestamp: time.Now().UnixNano() / int64(time.Millisecond),
				Value:     rand.Intn(100),
			}
			dataChan <- data
			time.Sleep(1000 * time.Millisecond)
		}
	}()

	return dataChan
}

func main() {
	topic := "my-topic"
	brokers := []string{"localhost:9092"}
	dataChan := generateData()

	go produce(dataChan, topic, brokers)
	consume(topic, brokers)
}
