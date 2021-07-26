package model

type Message struct {
	Message string
}

func (m *Message) Get() string {
	return m.Message
}