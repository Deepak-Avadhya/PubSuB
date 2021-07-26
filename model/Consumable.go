package model

type Consumable struct {
	message string
	offset  int
	ack   bool
	subscriptionId string
}

func NewConsumable(message string,offset int,subscriptionID string) *Consumable {
	return &Consumable{message: message,offset: offset,ack: false,subscriptionId: subscriptionID}
}

func (c *Consumable) Ack()  {
	c.ack =true
}
func (c *Consumable) GetMessage() string {
	return c.message
}
func (c *Consumable) GetOffset() int {
	return c.offset
}
func (c *Consumable) IsAck() bool {
	return c.ack
}
