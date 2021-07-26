package handler

import (
	"fmt"
	"pubsub/model"
)

type TopicHandler struct {
	topic *model.Topic
	subscriberWorkers map[string]*SubscriberWorker
}

func NewTopicHandler(topic *model.Topic) TopicHandler {
	return TopicHandler{topic: topic}
}

func (t *TopicHandler) publish()  {
	for _,topicSubscriber := range t.topic.GetSubscribers() {
		if topicSubscriber.IsSubscriberAvailable(){
			fmt.Println("about to start worker for subscriber ",topicSubscriber.GetSubscriptionId())
			startWorker(t,topicSubscriber)
		}else {
			fmt.Println(topicSubscriber.GetSubscriberFunc())
			fmt.Println("subscriberFunc not available for ",topicSubscriber.GetSubscriptionId())
		}
	}
}

func startWorker(t *TopicHandler,subscriber *model.TopicSubscriber)  {
	fmt.Println("creating new worker for subscriptionId:",subscriber.GetSubscriptionId())
	worker := NewSubscriberWorker(t.topic,subscriber)
	go worker.run()
}