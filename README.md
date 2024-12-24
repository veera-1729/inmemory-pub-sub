# In-Memory Pub-Sub System

This repository implements an in-memory Publish-Subscribe (Pub-Sub) system using Golang. It includes a broker, producers (publishers), and consumers, enabling data communication between components via topics.

## Features
- Add and manage topics dynamically.
- Subscribe consumers to topics.
- Publish messages to topics.
- Consume messages from topics.

---

## How It Works

### Components
1. **Broker**
   - Acts as a mediator between publishers and consumers.
   - Stores topics and routes messages to subscribed consumers.

2. **Producer**
   - Publishes messages to topics managed by the broker.

3. **Consumer**
   - Subscribes to topics and consumes messages.

---

## Example Workflow

Below is a sample implementation showcasing the interaction between producers, consumers, and the broker:

```go
package main

import (
	broker2 "github.com/veera-1729/pub-sub-2/broker"
	"github.com/veera-1729/pub-sub-2/consumer"
	"github.com/veera-1729/pub-sub-2/producer"
	"sync"
)

func main() {
	broker := broker2.NewBroker()

	c1 := consumer.NewConsumer("consumer_1", broker)
	c2 := consumer.NewConsumer("consumer_2", broker)

	p1 := producer.NewPublisher("prod_1", broker)
	p2 := producer.NewPublisher("prod_2", broker)

	broker.AddTopic("topic_1")
	broker.AddTopic("topic_2")

	broker.Subscribe("topic_1", c1.Name)
	broker.Subscribe("topic_2", c2.Name)
	c1.AddTopic("topic_1")
	c2.AddTopic("topic_2")

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		p1.Publish("topic_1")
	}()
	go func() {
		defer wg.Done()
		p2.Publish("topic_2")
	}()
	wg.Wait()

	wg.Add(2)
	go func() {
		defer wg.Done()
		c1.Consume()
	}()
	go func() {
		defer wg.Done()
		c2.Consume()
	}()

	wg.Wait()
}
```

---

## Sample Output

```plaintext
producer prod_2 -> produced data_1
producer prod_1 -> produced data_1
producer prod_1 -> produced data_12
producer prod_2 -> produced data_12
producer prod_2 -> produced data_123
producer prod_1 -> produced data_123
producer prod_1 -> produced data_1234
producer prod_2 -> produced data_1234
producer prod_2 -> produced data_12345
producer prod_1 -> produced data_12345
Consumer consumer_2 consuming data data_1
Consumer consumer_1 consuming data data_1
Consumer consumer_1 consumed data data_1
Consumer consumer_1 consuming data data_12
Consumer consumer_2 consumed data data_1
Consumer consumer_2 consuming data data_12
Consumer consumer_2 consumed data data_12
Consumer consumer_2 consuming data data_123
Consumer consumer_1 consumed data data_12
Consumer consumer_1 consuming data data_123
Consumer consumer_1 consumed data data_123
Consumer consumer_2 consumed data data_123
Consumer consumer_2 consuming data data_1234
Consumer consumer_1 consuming data data_1234
Consumer consumer_1 consumed data data_1234
Consumer consumer_1 consuming data data_12345
Consumer consumer_2 consumed data data_1234
Consumer consumer_2 consuming data data_12345
Consumer consumer_1 consumed data data_12345
Consumer consumer_2 consumed data data_12345
```

---

## Key Functions

### Broker
- `AddTopic(topic string)` - Adds a new topic.
- `Subscribe(topic string, consumerName string)` - Subscribes a consumer to a topic.

### Producer
- `Publish(topic string)` - Publishes messages to a topic.

### Consumer
- `Consume()` - Consumes messages from subscribed topics.
- `AddTopic(topic string)` - Associates the consumer with a topic.

---

## Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/veera-1729/pub-sub-2.git
   ```
2. Navigate to the project directory:
   ```bash
   cd pub-sub-2
   ```
3. Run the program:
   ```bash
   go run main.go
   ```

---

## License
This project is licensed under the [MIT License](LICENSE).

---

## Contributing
Contributions are welcome! Feel free to open issues or submit pull requests.

