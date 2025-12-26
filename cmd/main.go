package main

import (
	"whatsapp-lld/internal/domain"
	"whatsapp-lld/internal/services"
)

func main() {
	user1 := domain.User{
		Id:          1,
		Name:        "Alice",
		ReceivedMsg: make(map[int]*domain.Message),
	}
	user2 := domain.User{
		Id:          2,
		Name:        "Bob",
		ReceivedMsg: make(map[int]*domain.Message),
	}
	user3 := domain.User{
		Id:          3,
		Name:        "Cunha",
		ReceivedMsg: make(map[int]*domain.Message),
	}
	whatsapp := services.NewWhatsap()

	whatsapp.RegisterUser(&user1)
	whatsapp.RegisterUser(&user2)
	whatsapp.RegisterUser(&user3)

	whatsapp.Send(&domain.Message{
		Id:         1,
		SenderId:   1,
		ReceiverId: 2,
		Type:       domain.PersonalMessage,
		Content:    "Hi Bob, this is Alice",
	})

	whatsapp.CreateGroup(&domain.Group{
		Id:          1,
		Name:        "Friends",
		Members:     []*domain.User{&user1, &user2, &user3},
		ReceivedMsg: make(map[int]*domain.Message),
	})
	whatsapp.Send(&domain.Message{
		Id:         2,
		SenderId:   1,
		ReceiverId: 1,
		Type:       domain.GroupMessage,
		Content:    "Hey Guys!!",
	})
}
