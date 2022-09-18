package repository

import (
	"github.com/jmoiron/sqlx"
	gopherinskillbox "skillbox-test"
)

type User interface {
	CreateUser(user gopherinskillbox.User) (gopherinskillbox.User, error)
	GetAllUsers() ([]gopherinskillbox.User, error)
	DeleteUser(deleteId int) error
	UpdateUser(updateId int, data gopherinskillbox.UpdateUserData) error
	AddFriend(id int, data gopherinskillbox.UserFriendData) error
}

type Repository struct {
	User
}

// NewRepository - constructor for database maintainer
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User: NewUserSqlite(db),
	}
}
