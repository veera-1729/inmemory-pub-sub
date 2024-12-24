# In-Memory Pub-Sub System

An in-memory publish-subscribe (Pub-Sub) system implemented in Go. This project allows producers to publish messages to topics and consumers to subscribe to those topics and receive messages in real-time.

## Features

- **In-Memory Broker**: All data is managed in memory for fast message delivery.
- **Topic Management**: Dynamically add and manage topics.
- **Subscription Management**: Consumers can subscribe to multiple topics.
- **Concurrency**: Supports concurrent publishers and consumers using Go routines.

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) (version 1.18 or above)

### Installation

Clone the repository:

```bash
git clone https://github.com/<your-username>/<your-repo-name>.git
cd <your-repo-name>
```

### Running the Example

1. Build the project:

   ```bash
   go build -o pubsub
   ```

2. Run the main program:

   ```bash
   ./pubsub
   ```

### Code Structure

- `broker/` - Contains the broker implementation responsible for managing topics and subscriptions.
- `consumer/` - Contains the consumer implementation.
- `producer/` - Contains the producer implementation.
- `main.go` - Entry point for running the Pub-Sub example.

### Example Usage

#### Adding Topics and Subscribing Consumers

```go
broker := broker2.NewBroker()

// Add topics
broker.AddTopic("topic_1")
broker.AddTopic("topic_2")

// Create consumers
c1 := consumer.NewConsumer("consumer_1", broker)
c2 := consumer.NewConsumer("consumer_2", broker)

// Subscribe consumers to topics
broker.Subscribe("topic_1", c1.Name)
broker.Subscribe("topic_2", c2.Name)
```

#### Publishing Messages

```go
p1 := producer.NewPublisher("producer_1", broker)
p1.Publish("topic_1", "Message for topic 1")
```

#### Consuming Messages

```go
go c1.Consume()
go c2.Consume()
```

### License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

Feel free to contribute to this project by submitting issues or pull requests!
