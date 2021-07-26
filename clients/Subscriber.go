package clients

import (
	"fmt"
	"pubsub/model"
)

type Subscriber struct {
}

func (s *Subscriber) Consume(message *model.Consumable) {
	fmt.Printf("new msg received offset: %v and payload: %s\n", message.GetOffset(), message.GetMessage())
	message.Ack()
}

