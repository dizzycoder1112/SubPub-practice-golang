package main

import (
	"fmt"
	"testing"
)

func TestSubscribe(t *testing.T) {
	tests := []struct {
		topic      string
		subscriber *Subscriber
	}{
		{
			topic: "a",
			subscriber: &Subscriber{
				Topics: map[string]bool{},
				ID:     "alan",
			},
		},
		{
			topic: "a",
			subscriber: &Subscriber{
				Topics: map[string]bool{},
				ID:     "paul",
			},
		},
		{
			topic: "b",
			subscriber: &Subscriber{
				Topics: map[string]bool{},
				ID:     "champer",
			},
		},
		{
			topic: "b",
			subscriber: &Subscriber{
				Topics: map[string]bool{},
				ID:     "alex",
			},
		},
	}
	ps := NewPubSubService()
	for _, tt := range tests {
		t.Run(tt.topic, func(t *testing.T) {
			ps.Subscribe(tt.topic, tt.subscriber)
			fmt.Printf("user %s has %+v\n", tt.subscriber.ID, tt.subscriber.Topics)
		})
	}
	fmt.Printf("%+v\n", ps.Subscribers)

}

func TestPublish(t *testing.T) {
	tests := []struct {
		topic      string
		subscriber *Subscriber
	}{
		{
			topic: "a",
			subscriber: &Subscriber{
				Topics: map[string]bool{},
				ID:     "alan",
			},
		},
		{
			topic: "a",
			subscriber: &Subscriber{
				Topics: map[string]bool{},
				ID:     "paul",
			},
		},
		{
			topic: "b",
			subscriber: &Subscriber{
				Topics: map[string]bool{},
				ID:     "champer",
			},
		},
		{
			topic: "b",
			subscriber: &Subscriber{
				Topics: map[string]bool{},
				ID:     "alex",
			},
		},
	}
	ps := NewPubSubService()
	for _, tt := range tests {
		t.Run(tt.topic, func(t *testing.T) {
			ps.Subscribe(tt.topic, tt.subscriber)
			ps.Publish("a", []byte("hello"))
			ps.Publish("b", []byte("world"))

		})
	}

}
