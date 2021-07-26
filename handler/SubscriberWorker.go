package handler

import (
	"fmt"
	"pubsub/model"
	"sync"
	"time"
)

type SubscriberWorker struct {
	topic *model.Topic
	topicSubscriber *model.TopicSubscriber
	wg sync.WaitGroup
}

func NewSubscriberWorker(topic *model.Topic,subscriber *model.TopicSubscriber) *SubscriberWorker {
	return &SubscriberWorker{topic: topic,topicSubscriber: subscriber}
}

func (w *SubscriberWorker) run()  {
	for {
		currOffset :=w.topicSubscriber.GetOffset()
		if currOffset>=w.topic.GetMessageLength() {
			fmt.Printf("no message left to push going to sleep offset %v for subscriptionId %v\n", currOffset, w.topicSubscriber.GetSubscriptionId())
			break
		}
		fmt.Println("pushing message to subscriber:",w.topicSubscriber.GetSubscriptionId())
		message :=w.topic.GetMessages(currOffset)
		consumable := model.NewConsumable(message.Get(),currOffset,w.topicSubscriber.GetSubscriptionId())
		go (w.topicSubscriber.GetSubscriberFunc()).Consume(consumable)

		for i:=0;i<5;i++ {
			if consumable.IsAck() {
				w.topicSubscriber.IncrementOffset(currOffset)
				fmt.Printf("ack success incrementing offset %v for subscriptionId %v\n", w.topicSubscriber.GetOffset(), w.topicSubscriber.GetSubscriptionId())
				break
			}
			fmt.Println("waiting for ack...",w.topicSubscriber.GetSubscriptionId())
			time.Sleep(1 * time.Second)
		}
		if !consumable.IsAck() {
			fmt.Println("no ack retrying in a sec for subscriptionID:",w.topicSubscriber.GetSubscriptionId())
			time.Sleep(1 * time.Second)
		}
	}
}

