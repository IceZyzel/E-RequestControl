package services

import "request_manager_api/pkg/repository"

type Ticket interface {
}
type Notification interface {
}
type Authorization interface {
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
	return &Service{}
}
