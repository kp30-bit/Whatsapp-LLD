package services

import (
	"fmt"
	"sync"
	"whatsapp-lld/internal/domain"
	interfaces "whatsapp-lld/internal/interface"
)

// NotificationService implements the observer pattern for group message notifications
type NotificationService struct {
	observers map[int]interfaces.IObserver // userId -> observer mapping
	mu        sync.RWMutex
}

var (
	NotificationServiceInstance *NotificationService
	NotificationServiceOnce     sync.Once
)

func NewNotificationService() *NotificationService {
	NotificationServiceOnce.Do(func() {
		NotificationServiceInstance = &NotificationService{
			observers: make(map[int]interfaces.IObserver),
		}
	})
	return NotificationServiceInstance
}

// Subscribe adds an observer for a user
func (ns *NotificationService) Subscribe(userId int, observer interfaces.IObserver) {
	ns.mu.Lock()
	defer ns.mu.Unlock()
	ns.observers[userId] = observer
	fmt.Printf("User %d subscribed to notifications\n", userId)
}

// Unsubscribe removes an observer for a user
func (ns *NotificationService) Unsubscribe(userId int) {
	ns.mu.Lock()
	defer ns.mu.Unlock()
	delete(ns.observers, userId)
	fmt.Printf("User %d unsubscribed from notifications\n", userId)
}

// NotifyGroupMembers notifies all members of a group about a new message
func (ns *NotificationService) NotifyGroupMembers(message *domain.Message, groupId int) {
	ns.mu.RLock()
	defer ns.mu.RUnlock()

	// Get the group to find all members
	group, err := MessageServiceInstance.GetGroupById(groupId)
	if err != nil {
		fmt.Printf("Failed to get group %d for notification: %v\n", groupId, err)
		return
	}

	// Notify all group members (except the sender)
	for _, member := range group.Members {
		if member.Id != message.SenderId {
			if observer, exists := ns.observers[member.Id]; exists {
				observer.Notify(message, groupId)
			}
		}
	}
}

