package usecase

import (
	"fmt"
	"whatsapp-lld/internal/domain"
)

type PersonalMessageSender struct {
	UserList map[int]*domain.User
}

func (pm PersonalMessageSender) Send(message domain.Message) error {
	for _, user := range pm.UserList {
		if user.Id == message.ReceiverId {
			user.ReceivedMsg[message.Id] = &message
			fmt.Printf("Message : %v\t Sender: %v | Receiever %v\n", message.Content, message.SenderId, message.ReceiverId)
			return nil
		}
	}
	return fmt.Errorf("No person found with id : %v\n", message.ReceiverId)
}
