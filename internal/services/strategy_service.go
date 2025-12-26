package services

import (
	"fmt"
	"whatsapp-lld/internal/domain"
	interfaces "whatsapp-lld/internal/interface"
	"whatsapp-lld/internal/usecase"
)

type StrategyService struct {
	personalStrategy *usecase.PersonalMessageSender
	groupStrategy    *usecase.GroupMessageSender
}

func NewStrategyService() *StrategyService {
	return &StrategyService{
		personalStrategy: usecase.NewPersonalMessageSender(),
		groupStrategy:    usecase.NewGroupMessageSender(),
	}
}

func (ss *StrategyService) GetDeliveryStrategy(messageType domain.MessageType) (interfaces.IMessageStrategy, error) {
	switch messageType {
	case domain.GroupMessage:
		return ss.groupStrategy, nil
	case domain.PersonalMessage:
		return ss.personalStrategy, nil
	default:
		return nil, fmt.Errorf("unknown message type: %v", messageType)
	}
}
