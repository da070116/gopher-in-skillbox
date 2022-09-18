package service

import (
	gopherinskillbox "skillbox-test"
	"skillbox-test/pkg/repository"
)

type UserService struct {
	repo repository.User
}

func (s *UserService) DeleteUser(deleteId int) error {
	return s.repo.DeleteUser(deleteId)
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user gopherinskillbox.User) (gopherinskillbox.User, error) {
	return s.repo.CreateUser(user)
}

func (s *UserService) GetAllUsers() ([]gopherinskillbox.User, error) {
	return s.repo.GetAllUsers()
}

func (s *UserService) UpdateUser(userId int, userData gopherinskillbox.UpdateUserData) error {
	return s.repo.UpdateUser(userId, userData)
}

func (s *UserService) AddFriend(id int, data gopherinskillbox.UserFriendData) error {
	return s.repo.AddFriend(id, data)
}
