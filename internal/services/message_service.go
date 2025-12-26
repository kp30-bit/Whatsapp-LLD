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

func (ms *MessageService) GetGroupById(id int) (*domain.Group, error) {
	group, exists := ms.groups[id]
	if !exists {
		return nil, fmt.Errorf("Group not found with id : %v", id)
	}
	return group, nil
}

func (ms *MessageService) AddUserToGroup(user *domain.User, gId int) error {
	group, err := ms.GetGroupById(gId)
	if err != nil {
		return err
	}
	group.Members = append(group.Members, user)
	return nil
}

func (ms *MessageService) SendMessage(message *domain.Message, sender interfaces.IMessageStrategy, userService interfaces.IUserService) error {
	err := sender.Send(*message, userService, ms)
	if err != nil {
		return err
	}
	return nil
}
