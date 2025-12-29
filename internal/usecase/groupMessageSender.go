package usecase

import (
	"fmt"
	"whatsapp-lld/internal/domain"
	interfaces "whatsapp-lld/internal/interface"
)

type GroupMessageSender struct{}

func NewGroupMessageSender() *GroupMessageSender {
	return &GroupMessageSender{}
}

func (g *GroupMessageSender) isMember(group *domain.Group, userId int) bool {
	for _, member := range group.Members {
		if member.Id == userId {
			return true
		}
	}
	return false
}

func (g *GroupMessageSender) Send(message domain.Message, userService interfaces.IUserService, messageService interfaces.IMessageService, notificationService interfaces.INotificationService) error {
	group, err := messageService.GetGroupById(message.ReceiverId)
	if err != nil {
		return fmt.Errorf("No group found with id : %v", message.ReceiverId)
	}

	// Validate sender is a member of the group
	if !g.isMember(group, message.SenderId) {
		return fmt.Errorf("Sender %d is not a member of group %d", message.SenderId, group.Id)
	}

	// Deliver message to all group members (except sender)
	for _, member := range group.Members {
		if member.Id != message.SenderId {
			member.ReceivedMsg[message.Id] = &message
		}
	}

	// Store message in group
	group.ReceivedMsg[message.Id] = &message
	fmt.Printf("Message : %v\t Sender: %v | Receiver (Group) %v\n", message.Content, message.SenderId, message.ReceiverId)
	
	// Notify all group members using observer pattern
	if notificationService != nil {
		notificationService.NotifyGroupMembers(&message, group.Id)
	}
	
	return nil
}
