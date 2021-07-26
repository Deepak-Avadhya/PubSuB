package handler

import (
	"fmt"
	"pubsub/model"
)

type PubSub struct {
	topicMap map[string]*TopicHandler
	subscriptionMap map[string]*model.Topic
}

func NewPubSub() PubSub {
	return PubSub{topicMap: map[string]*TopicHandler{},subscriptionMap: map[string]*model.Topic{}}
}

func (p * PubSub) CreateTopic(topicId string) *model.Topic {
	if v,ok :=p.topicMap[topicId]; ok {
		fmt.Println("topic Already exist")
		return v.topic
	}
	topic := model.NewTopic(topicId)
	topicHandler :=NewTopicHandler(&topic)
	p.topicMap[topic.GetTopicId()]=&topicHandler
	return &topic
}

func (p * PubSub) DeleteTopic(topicId string) {
	delete(p.topicMap,topicId)
	for k,v := range p.subscriptionMap {
		if v.GetTopicId()==topicId {
			delete(p.subscriptionMap,k)
		}
	}
	fmt.Println("topic deleted successfully:",topicId)
}

func (p *PubSub) AddSubscription(topic *model.Topic,subscriptionID string) {
	topic.AddSubscriber(model.NewTopicSubscriber(subscriptionID))
	p.subscriptionMap[subscriptionID]=topic
	fmt.Printf("subscription %v added to topic %v\n", subscriptionID, topic.GetTopicId())
}

func (p *PubSub) DeleteSubscription(subscriptionID string) {
	if topic,ok := p.subscriptionMap[subscriptionID] ;ok {
		topic.RemoveSubscriber(subscriptionID)
	}else {
		fmt.Println("subscriptionId not present:",subscriptionID)
	}
	delete(p.subscriptionMap,subscriptionID)
	fmt.Println("subscription deleted",subscriptionID)
}

func (p *PubSub) Publish(topic *model.Topic,message string)  {
	if topicHandler,ok :=p.topicMap[topic.GetTopicId()]; ok {
		topic.AddMessage(model.Message{Message:message})
		topicHandler.publish()
	}else {
		fmt.Println("topic not found")
	}

}
func (p *PubSub) Subscribe(subscriptionID string,subscriberFunc model.ISubscriber)  {
	if topic,ok := p.subscriptionMap[subscriptionID] ;ok {
		topic.AddSubscriptionFunc(subscriptionID,subscriberFunc)
	}else{
		//return error
		fmt.Println("subscriptionId not present:",subscriptionID)
	}
}

func (p *PubSub) UnSubscribe(subscriptionID string)  {
	if topic,ok := p.subscriptionMap[subscriptionID] ;ok {
		topic.RemoveSubscriptionFunc(subscriptionID)
	}else {
		fmt.Println("subscriptionId not present:",subscriptionID)
	}
}
