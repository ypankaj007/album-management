package main

import (
	"fmt"

	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/util/log"
)

const (
	topic = "go.micro.topic.log"
)

func main() {
	if err := broker.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}
	if err := broker.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}

	sub()
	select {}
}

func sub() {
	_, err := broker.Subscribe(topic, func(p broker.Event) error {
		log.Log(string(p.Message().Body))
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}
