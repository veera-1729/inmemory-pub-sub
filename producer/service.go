package producer

import (
	"fmt"
	"github.com/veera-1729/pub-sub-2/broker"
	"time"
)

type Publisher struct {
	Name   string
	broker *broker.Broker
	Topics []string
}

func NewPublisher(name string, broker *broker.Broker) Publisher {
	return Publisher{
		Name:   name,
		broker: broker,
	}

}

func (p *Publisher) Publish(topic string) {
	d := "data_"
	for i := 1; i <= 5; i++ {
		d = fmt.Sprintf("%s%v", d, i)

		fmt.Println(fmt.Sprintf("producer %s -> produced %s", p.Name, d))

		p.broker.AddDataToTopic(topic, d)

		time.Sleep(1 * time.Second)
	}

}
