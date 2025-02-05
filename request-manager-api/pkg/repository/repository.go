package repository

import "github.com/jmoiron/sqlx"

type Ticket interface {
}
type Notification interface {
}
type Authorization interface {
}
type Admin interface {
}
type Repository struct {
	Ticket
	Notification
	Authorization
	Admin
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
