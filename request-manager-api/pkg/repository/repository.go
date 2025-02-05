package repository

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

func NewRepository() *Repository {
	return &Repository{}
}
