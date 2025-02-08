package repository

import (
	"github.com/jmoiron/sqlx"
	Request_Manager "request_manager_api"
)

type Ticket interface {
	GetUserTickets(userID int) ([]Request_Manager.Ticket, error)
	CreateTicket(ticket Request_Manager.Ticket) (int, error)
	UpdateTicket(ticketID int, input Request_Manager.UpdateTicketInput) error
	DeleteTicket(ticketID int) error
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
	CreateUser(user Request_Manager.User) (int, error)
	GetAllUsers() ([]Request_Manager.User, error)
	Delete(UserID int) error
	UpdateUser(UserID int, input Request_Manager.UpdateUserInput, user Request_Manager.User) error
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
		Ticket:        NewTicketMysql(db),
	}
}
