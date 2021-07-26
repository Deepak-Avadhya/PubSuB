package model

import "fmt"

type Topic struct {
	topicId   string
	messages  []Message
	subscribers []*TopicSubscriber
}

func (t *Topic) AddMessage(message Message)  {
	t.messages =append(t.messages,message)
	fmt.Printf("new message received in topic %v and currSize %v\n", t.topicId, t.GetMessageLength())
}
func (t *Topic) AddSubscriber(subscriber TopicSubscriber)  {
	t.subscribers =append(t.subscribers,&subscriber)
	fmt.Println("new subscriber added:",subscriber.subscriptionId)
}
func NewTopic(topicId string) Topic {
	fmt.Println("creating new topic",topicId)
	return Topic{topicId: topicId}
}
func (t *Topic)GetMessages(idx int) Message {
	if idx<len(t.messages){
		return t.messages[idx]
	}else {
		//throw error
		fmt.Printf("index out of bound currSize %v idx %v\n", len(t.messages), idx)
		return Message{}
	}
}
func (t *Topic) GetTopicId() string {
	return t.topicId
}
func (t *Topic) GetSubscribers() []*TopicSubscriber {
	return t.subscribers
}
func (t *Topic) GetMessageList() *[]Message {
	return &t.messages
}

func (t *Topic) GetMessageLength() int {
	return len(t.messages)
}

func (t *Topic) AddSubscriptionFunc(subscriptionID string,subscriberFunc ISubscriber)  {
	for _, topicSubscriber :=range t.subscribers {
		if topicSubscriber.GetSubscriptionId()==subscriptionID {
			topicSubscriber.SetSubscriberFunc(subscriberFunc)
			fmt.Println("subscriberFunc added to subscriptionID ",subscriptionID)
		}
	}
}
func (t *Topic) RemoveSubscriptionFunc(subscriptionID string)  {
	for _, topicSubscriber :=range t.subscribers {
		if topicSubscriber.GetSubscriptionId()==subscriptionID {
			topicSubscriber.SetSubscriberFunc(nil)
			fmt.Println("subscriberFunc added to subscriptionID ",subscriptionID)
		}
	}
}
func (t *Topic) RemoveSubscriber(subscriptionID string)  {

	for i:=0;i<len(t.subscribers);i++ {
		if t.subscribers[i].subscriptionId==subscriptionID {
			t.subscribers[i]=t.subscribers[len(t.subscribers)-1]
			t.subscribers=t.subscribers[:len(t.subscribers)-1]
			fmt.Println("subscription removed",subscriptionID)
			return
		}
	}
	fmt.Println("subscription not present",subscriptionID)
}