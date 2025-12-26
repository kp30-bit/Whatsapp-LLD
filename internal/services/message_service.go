package services

import (
	"fmt"
	"sync"
	"whatsapp-lld/internal/domain"
	interfaces "whatsapp-lld/internal/interface"
)

type MessageService struct {
	groups map[int]*domain.Group
	// users  map[int]*domain.User
}

var (
	MessageServiceInstance *MessageService
	MessageServiceOnce     sync.Once
)

func NewMessageService() *MessageService {
	MessageServiceOnce.Do(func() {
		MessageServiceInstance = &MessageService{
			groups: make(map[int]*domain.Group),
			// users:  make(map[int]*domain.User),
		}
	})
	return MessageServiceInstance
}

func (ms *MessageService) CreateGroup(group *domain.Group) {
	ms.groups[group.Id] = group
	fmt.Printf("Group with name : %v is created\n", group.Name)
}

func (ms *MessageService) AddUserToGroup(user *domain.User, gId int) error {
	for _, group := range ms.groups {
		if group.Id == gId {
			group.Members = append(group.Members, user)
			return nil
		}
	}
	return fmt.Errorf("Group not found with id : %v\n", gId)
}

func (ms *MessageService) SendMessage(message *domain.Message, sender interfaces.IMessageStrategy) error {
	err := sender.Send(*message)
	if err != nil {
		return err
	}
	return nil
}
