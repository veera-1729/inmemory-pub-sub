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
