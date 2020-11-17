package chat

import "all/internal/config"

type Message struct {
	ID   int64
	Text string
}

type chatService interface {
	AddMessage(Message) error
	FindByID(int) *Message
	FindAll() []*Message
}

type service struct {
	conf *config.Config
}

func New(c *config.Config) (chatService, error) {
	return service{c}, nil
}

func (s service) AddMessage(Message) error {
	return nil
}
func (s service) FindByID(int) *Message {
	return nil
}

func (s service) FindAll() []*Message {
	var list []*Message
	list = append(list, &Message{0, "hello wordl"})
	return list
}
