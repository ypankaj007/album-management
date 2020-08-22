package service

import (
	"fmt"
	"log"

	"github.com/micro/go-micro/broker"
)

const (
	logTopic = "go.micro.topic.log"
)

// PubSub godoc
// publisher pubish the message in the specified topic
type PubSub struct {
}

// NewPublisher godoc
// Initialize publisher and br
func NewPublisher() *PubSub {
	if err := broker.Init(); err != nil {
		log.Fatalf("broker Init error: %v", err)
	}

	if err := broker.Connect(); err != nil {
		log.Fatalf("broker Connect error: %v", err)
	}
	return &PubSub{}
}

// Log godoc
// log publish message to logger service to put
// log in file
func (ps *PubSub) Log(msg string) {
	content := []byte(fmt.Sprintf("\n%s", msg))
	topicMsg := &broker.Message{
		Body: content,
	}
	if err := broker.Publish(logTopic, topicMsg); err != nil {
		log.Printf("[pub] failed: %v", err)
	} else {
		fmt.Println("[pub] pubbed message:", string(topicMsg.Body))
	}

}
