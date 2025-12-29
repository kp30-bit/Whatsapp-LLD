package services

import (
	"fmt"
	"whatsapp-lld/internal/domain"
)

// UserObserver implements IObserver for user notifications
type UserObserver struct {
	UserId int
	Name   string
}

func NewUserObserver(userId int, name string) *UserObserver {
	return &UserObserver{
		UserId: userId,
		Name:   name,
	}
}

// Notify is called when a group message is sent
func (uo *UserObserver) Notify(message *domain.Message, groupId int) {
	fmt.Printf("🔔 Notification to User %d (%s): New message in Group %d from User %d: \"%s\"\n",
		uo.UserId, uo.Name, groupId, message.SenderId, message.Content)
}

