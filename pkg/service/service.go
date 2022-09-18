package service

import (
	gopherinskillbox "skillbox-test"
	"skillbox-test/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type User interface {
	CreateUser(user gopherinskillbox.User) (gopherinskillbox.User, error)
	GetAllUsers() ([]gopherinskillbox.User, error)
	DeleteUser(deleteId int) error
	UpdateUser(updateId int, data gopherinskillbox.UpdateUserData) error
	AddFriend(id int, data gopherinskillbox.UserFriendData) error
}

type Service struct {
	User
}

// NewService constructor
func NewService(repo *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repo),
	}
}
