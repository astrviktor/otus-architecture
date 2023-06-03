package memorymq

import (
	"otus-architecture/hw08/broker"
)

type MemoryMQ struct {
	messages chan *broker.Message
}

func New() *MemoryMQ {
	return &MemoryMQ{}
}

func (m *MemoryMQ) Connect() error {
	messages := make(chan *broker.Message, 100)
	m.messages = messages

	return nil
}

func (m *MemoryMQ) Send(message *broker.Message) error {
	m.messages <- message

	return nil
}

func (m *MemoryMQ) Receive() (*broker.Message, error) {
	message := <-m.messages

	return message, nil
}

func (m *MemoryMQ) Close() error {
	return nil
}
