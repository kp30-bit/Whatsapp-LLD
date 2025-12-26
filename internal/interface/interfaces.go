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
	Send(message domain.Message, userService IUserService, messageService IMessageService) error
}
