package usecase

import (
	"fmt"
	"whatsapp-lld/internal/domain"
	interfaces "whatsapp-lld/internal/interface"
)

type PersonalMessageSender struct{}

func NewPersonalMessageSender() *PersonalMessageSender {
	return &PersonalMessageSender{}
}

func (pm *PersonalMessageSender) Send(message domain.Message, userService interfaces.IUserService, messageService interfaces.IMessageService) error {
	receiver, err := userService.GetUserById(message.ReceiverId)
	if err != nil {
		return fmt.Errorf("No person found with id : %v", message.ReceiverId)
	}
	receiver.ReceivedMsg[message.Id] = &message
	fmt.Printf("Message : %v\t Sender: %v | Receiver %v\n", message.Content, message.SenderId, message.ReceiverId)
	return nil
}
