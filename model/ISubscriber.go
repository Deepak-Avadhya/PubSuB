package model

type ISubscriber interface {
	Consume(message *Consumable)
}
