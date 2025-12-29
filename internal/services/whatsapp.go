package services

import (
	"fmt"
	"sync"
	"whatsapp-lld/internal/domain"
)

// this is the facade that is exosed to clients

type whatsapp struct {
	MessageService      *MessageService
	StrategyService     *StrategyService
	UserService         *UserService
	NotificationService *NotificationService
}

var (
	WhatsappInst *whatsapp
	WhatsappOnce sync.Once
)

func NewWhatsap() *whatsapp {
	WhatsappOnce.Do(func() {
		messageService := NewMessageService()
		userService := NewUserService()
		notificationService := NewNotificationService()
		WhatsappInst = &whatsapp{
			MessageService:      messageService,
			StrategyService:     NewStrategyService(),
			UserService:         userService,
			NotificationService: notificationService,
		}
	})
	return WhatsappInst
}

func (w *whatsapp) RegisterUser(user *domain.User) {
	w.UserService.RegisterUser(user)
	// Subscribe user to notifications
	observer := NewUserObserver(user.Id, user.Name)
	w.NotificationService.Subscribe(user.Id, observer)
}

func (w *whatsapp) AddUserToGroup(user *domain.User, gId int) {
	w.MessageService.AddUserToGroup(user, gId)
}

func (w *whatsapp) CreateGroup(group *domain.Group) {
	w.MessageService.CreateGroup(group)
}

func (w *whatsapp) Send(message *domain.Message) {
	sender, err := w.StrategyService.GetDeliveryStrategy(message.Type)
	if err != nil {
		fmt.Printf("Failed to get delivery strategy with error : %v\n", err)
		return
	}
	err = w.MessageService.SendMessage(message, sender, w.UserService, w.NotificationService)
	if err != nil {
		fmt.Printf("Failed to send message with error : %v\n", err)
	}
}
