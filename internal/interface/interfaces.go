package interfaces

import "whatsapp-lld/internal/domain"

// IUserService defines the interface for user-related operations
type IUserService interface {
	GetUserById(id int) (*domain.User, error)
}

// IMessageService defines the interface for message-related operations
type IMessageService interface {
	GetGroupById(id int) (*domain.Group, error)
}

type IMessageStrategy interface {
	Send(message domain.Message, userService IUserService, messageService IMessageService, notificationService INotificationService) error
}

// IObserver defines the interface for observers in the observer pattern
type IObserver interface {
	Notify(message *domain.Message, groupId int)
}

// INotificationService defines the interface for notification operations
type INotificationService interface {
	Subscribe(userId int, observer IObserver)
	Unsubscribe(userId int)
	NotifyGroupMembers(message *domain.Message, groupId int)
}
