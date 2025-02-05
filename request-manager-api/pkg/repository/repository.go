package repository

import (
	"github.com/jmoiron/sqlx"
	Request_Manager "request_manager_api"
)

type Ticket interface {
}
type Notification interface {
}
type Authorization interface {
	CreateAdmin(user Request_Manager.User) (int, error)
	CreateUser(user Request_Manager.User) (int, error)
	GetUser(username, password string) (Request_Manager.User, error)
}
type Admin interface {
	GetUserByID(userID int) (Request_Manager.User, error)
}
type Repository struct {
	Ticket
	Notification
	Authorization
	Admin
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthMysql(db),
		Admin:         NewAdminMysql(db),
	}
}
