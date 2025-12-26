package services

import (
	"fmt"
	"sync"
	"whatsapp-lld/internal/domain"
)

// this is the facade that is exosed to clients

type whatsapp struct {
	MessageService  *MessageService
	StrategyService *StrategyService
	UserService     *UserService
}

var (
	WhatsappInst *whatsapp
	WhatsappOnce sync.Once
)

func NewWhatsap() *whatsapp {
	WhatsappOnce.Do(func() {
		WhatsappInst = &whatsapp{
			MessageService:  NewMessageService(),
			StrategyService: NewStrategyService(),
			UserService:     NewUserService(),
		}
	})
	return WhatsappInst
}

func (w *whatsapp) RegisterUser(user *domain.User) {
	w.UserService.RegisterUser(user)
}

func (w *whatsapp) AddUserToGroup(user *domain.User, gId int) {
	w.MessageService.AddUserToGroup(user, gId)
}

func (w *whatsapp) CreateGroup(group *domain.Group) {
	w.MessageService.CreateGroup(group)
}

func (w *whatsapp) Send(message *domain.Message) {
	sender := w.StrategyService.GetDeliveryStrategy(message.Type, w.UserService.users, w.MessageService.groups)
	err := w.MessageService.SendMessage(message, sender)
	if err != nil {
		fmt.Printf("Failed to send message with error : %v\n", err)
	}
}
