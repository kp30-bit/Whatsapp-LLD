package services

import (
	"whatsapp-lld/internal/domain"
	interfaces "whatsapp-lld/internal/interface"
	"whatsapp-lld/internal/usecase"
)

type StrategyService struct {
}

func NewStrategyService() *StrategyService {
	return &StrategyService{}
}

func (ss *StrategyService) GetDeliveryStrategy(messageType domain.MessageType, userList map[int]*domain.User, groupList map[int]*domain.Group) interfaces.IMessageStrategy {
	switch messageType {
	case domain.GroupMessage:
		return usecase.GroupMessageSender{
			GroupList: groupList,
		}
	case domain.PersonalMessage:
		return usecase.PersonalMessageSender{
			UserList: userList,
		}
	}
	return nil
}
