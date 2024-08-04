# GoKafka - Building a Real-time Data Pipeline with Go and Kafka

**Disclaimer**: This example provides a basic outline of a real-time data pipeline using Go and Kafka. Real-world implementations would require additional considerations like error handling, performance optimization, and security.

**Prerequisites**
A running Kafka cluster
Go environment set up
Necessary libraries installed: ```github.com/IBM/sarama``` for Kafka integration

**Code Structure**
We'll break down the pipeline into three main components:

- Data Source: Simulating data generation.
- Kafka Producer: Sending data to Kafka.
- Kafka Consumer: Consuming data and processing it.


## Running the Code: A Step-by-Step Guide
### Prerequisites
- **Java:** Ensure Java 8 or later is installed.
- **Kafka:** Download and extract the Apache Kafka distribution.
- **Go:** Set up Go environment and install the ```github.com/IBM/sarama``` library.

### Steps
**1. Start Zookeeper**
- Navigate to the Kafka installation directory.
- Open a terminal and run:
```bash
bin/zookeeper-server-start.sh config/zookeeper.properties
```

**2. Start Kafka**
- In a new terminal, still in the Kafka installation directory, run:
```bash
bin/kafka-server-start.sh config/server.properties
```

**3. Create a Kafka Topic**
In a new terminal, run:
```bash
bin/kafka-topics.sh --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --topic my-topic
```

**4. Run the Go Application**
- Save the provided Go code as a .go file (e.g., main.go).
- Compile and run the Go application:
```bash
go run main.go consumer.go producer.go
or 
go run ./*.go
or 
go build -o myapp && ./myapp
```

**5. Verify**
- To check if data is being produced and consumed, you can use the Kafka console consumer
```bash
bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic my-topic --from-beginning
```