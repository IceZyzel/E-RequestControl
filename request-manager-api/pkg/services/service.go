package services

import (
	Request_Manager "request_manager_api"
	"request_manager_api/pkg/repository"
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
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}
type Admin interface {
	GetUserByID(userID int) (Request_Manager.User, error)
	CreateUser(user Request_Manager.User) (int, error)
	GetAllUsers() ([]Request_Manager.User, error)
	Delete(UserID int) error
	UpdateUser(UserID int, input Request_Manager.UpdateUserInput, user Request_Manager.User) error
}
type Service struct {
	Authorization
	Ticket
	Notification
	Admin
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Admin:         NewAdminService(repos.Admin),
		Ticket:        NewTicketService(repos.Ticket),
	}
}
