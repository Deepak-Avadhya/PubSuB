package main

import (
	"pubsub/clients"
	"pubsub/handler"
	"time"
)

func main() {
	pubsub := handler.NewPubSub()
	topicA := pubsub.CreateTopic("deepak")
	pubsub.Publish(topicA,"hello world")

	pubsub.AddSubscription(topicA,"sd1")

	pubsub.Subscribe("sd1", &clients.Subscriber{})
	pubsub.AddSubscription(topicA,"sd2")

	pubsub.Subscribe("sd2", &clients.Subscriber{})

	pubsub.Publish(topicA,"hello world 2")
	time.Sleep(10*time.Second)
	pubsub.DeleteSubscription("sd1")

	pubsub.Publish(topicA,"hello world 3")


	time.Sleep(20*time.Second)

}
