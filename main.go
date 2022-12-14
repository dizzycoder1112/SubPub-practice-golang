package main

import (
	"fmt"
	"sync"
)

type Subscriber struct {
	ID     string
	Topics map[string]bool
}

func NewSubscriber() *Subscriber {
	return &Subscriber{
		Topics: make(map[string]bool),
	}
}

type PubSubService struct {
	Subscribers map[string]*Subscriber
	Topics      map[string]map[string]*Subscriber
	Mutex       sync.RWMutex // mutex lock
	Message     chan []byte
}

func NewPubSubService() *PubSubService {
	return &PubSubService{
		Subscribers: make(map[string]*Subscriber),
		Topics:      make(map[string]map[string]*Subscriber),
	}

}

func (ps *PubSubService) Subscribe(topic string, subscriber *Subscriber) {
	ps.Mutex.Lock()
	ps.Subscribers[subscriber.ID] = subscriber
	defer ps.Mutex.Unlock()

	newSub := map[string]*Subscriber{}
	newSub[subscriber.ID] = subscriber

	ps.Subscribers[subscriber.ID] = subscriber
	if _, ok := ps.Topics[topic]; !ok {
		ps.Topics[topic] = newSub

	}
	ps.Topics[topic][subscriber.ID] = subscriber
	subscriber.Topics[topic] = true

}

func (ps *PubSubService) Publish(topic string, message []byte) {

	ps.Mutex.RLock()
	topics := ps.Topics[topic]

	defer ps.Mutex.RUnlock()
	for _, sub := range topics {
		go (func(subscriber *Subscriber) {
			subscriber.OnMessage(topic, message)
		})(sub)
	}

}

func (s *Subscriber) OnMessage(topic string, message []byte) {
	fmt.Printf("Got message: %+s\n", string(message))
}

func main() {
	ps := NewPubSubService()
	sub1 := NewSubscriber()
	sub1.ID = "foo"
	ps.Subscribe("haha", sub1)
	fmt.Printf("%+v\n", sub1)
}
