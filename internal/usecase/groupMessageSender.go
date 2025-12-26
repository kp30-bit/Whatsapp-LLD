package usecase

import (
	"fmt"
	"whatsapp-lld/internal/domain"
)

type GroupMessageSender struct {
	GroupList map[int]*domain.Group
}

func (g GroupMessageSender) Send(message domain.Message) error {
	for _, group := range g.GroupList {
		if group.Id == message.ReceiverId {
			group.ReceivedMsg[message.Id] = &message
			fmt.Printf("Message : %v\t Sender: %v | Receiever %v\n", message.Content, message.SenderId, message.ReceiverId)
			return nil
		}
	}
	return fmt.Errorf("No group found with id : %v\n", message.ReceiverId)
}
