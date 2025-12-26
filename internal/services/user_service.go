package services

import (
	"fmt"
	"sync"
	"whatsapp-lld/internal/domain"
)

type UserService struct {
	users map[int]*domain.User
}

var (
	UserInstance *UserService
	UserOnce     sync.Once
)

func NewUserService() *UserService {
	UserOnce.Do(func() {
		UserInstance = &UserService{
			users: make(map[int]*domain.User),
		}
	})
	return UserInstance

}

func (u *UserService) RegisterUser(user *domain.User) {
	u.users[user.Id] = user
}

func (u *UserService) GetUserById(id int) (*domain.User, error) {
	for _, user := range u.users {
		if user.Id == id {
			return user, nil
		}
	}
	return nil, fmt.Errorf("User with id : %v not found\n", id)
}
