package interfaces

import "whatsapp-lld/internal/domain"

type IMessageStrategy interface {
	Send(message domain.Message) error
}
