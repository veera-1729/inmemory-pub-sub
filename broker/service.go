package broker

import (
	"fmt"
	"sync"
)

type Subscriber struct {
	active bool
	offset int
	name   string
}

type Subscribers map[string]Subscriber

type Broker struct {
	Topics map[string]Subscribers
	Data   map[string][]string

	mutex *sync.Mutex
}

func NewBroker() *Broker {
	return &Broker{
		Topics: make(map[string]Subscribers),
		Data:   make(map[string][]string),
		mutex:  &sync.Mutex{},
	}
}

func (b *Broker) AddTopic(name string) bool {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	if _, ok := b.Topics[name]; ok {
		panic("topic already exists")
	}
	b.Topics[name] = Subscribers{}
	b.Data[name] = make([]string, 0)
	return true
}

func (b *Broker) Subscribe(topic string, sub string) bool {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	if _, ok := b.Topics[topic]; !ok {
		panic("topic not found")
	}
	t := b.Topics[topic]

	if _, ok := t[sub]; ok {
		fmt.Println("already subscribed")
		return false
	}
	s := Subscriber{
		name:   sub,
		offset: 0,
		active: true,
	}
	t[sub] = s
	return true
}

func (b *Broker) GetData(topic string) []string {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	return b.Data[topic]
}
func (b *Broker) GetOffSet(topic string, consumer string) int {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	s := b.Topics[topic]
	return s[consumer].offset
}

func (b *Broker) UpdateOffSet(topic, consumer string, offset int) {

	b.mutex.Lock()
	defer b.mutex.Unlock()
	t := b.Topics[topic]
	s := t[consumer]
	s.offset = offset

	b.Topics[topic][consumer] = s
}

func (b *Broker) AddDataToTopic(topic, data string) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	d := b.Data[topic]

	d = append(d, data)

	b.Data[topic] = d
}
