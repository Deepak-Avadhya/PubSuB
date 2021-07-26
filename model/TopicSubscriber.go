package model

import "fmt"

type TopicSubscriber struct {
	offset     int
	subscriber ISubscriber
	subscriptionId string
}

func NewTopicSubscriber(subscriptionId string) TopicSubscriber  {
	fmt.Println("creating new topic subscriber",subscriptionId)
	return TopicSubscriber{offset: 0,subscriptionId:subscriptionId}
}
func (t *TopicSubscriber)GetOffset() int {
	return t.offset
}
func (t *TopicSubscriber)GetSubscriptionId() string {
	return t.subscriptionId
}
func (t *TopicSubscriber)GetSubscriberFunc() ISubscriber{
	return t.subscriber
}
func (t *TopicSubscriber) IncrementOffset(currOffset int) bool {
	// implement sync
	t.offset++
	return true
}
func (t *TopicSubscriber) SetSubscriberFunc(subscriberFunc ISubscriber){
	t.subscriber=subscriberFunc
}
func (t *TopicSubscriber) IsSubscriberAvailable() bool {
	return t.subscriber!=nil
}