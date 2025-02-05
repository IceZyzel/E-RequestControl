package services

import (
	Request_Manager "request_manager_api"
	"request_manager_api/pkg/repository"
)

type Ticket interface {
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
	}
}
