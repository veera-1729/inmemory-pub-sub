package consumer

import (
	"fmt"
	"github.com/veera-1729/pub-sub-2/broker"
	"time"
)

type Consumer struct {
	Name   string
	broker *broker.Broker
	Topics []string
}

func NewConsumer(name string, broker *broker.Broker) Consumer {
	return Consumer{
		Name:   name,
		broker: broker,
		Topics: make([]string, 0),
	}
}
func (c *Consumer) AddTopic(topic string) {
	c.Topics = append(c.Topics, topic)
}

func (c *Consumer) Consume() {

	for _, t := range c.Topics {

		topic := c.broker.GetData(t)
		offSet := c.broker.GetOffSet(t, c.Name)
		for offSet < len(topic) {

			data := topic[offSet]

			fmt.Printf("Consumer %s consuming data %s\n", c.Name, data)
			time.Sleep(1 * time.Second)

			c.broker.UpdateOffSet(t, c.Name, offSet)

			offSet++
		}
	}
}
