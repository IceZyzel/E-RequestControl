package services

import (
	Request_Manager "request_manager_api"
	"request_manager_api/pkg/repository"
)

type TicketService struct {
	repo repository.Ticket
}

func NewTicketService(repo repository.Ticket) *TicketService {
	return &TicketService{repo: repo}
}

func (s *TicketService) CreateTicket(ticket Request_Manager.Ticket) (int, error) {
	return s.repo.CreateTicket(ticket)
}

func (s *TicketService) GetUserTickets(ticketID int) ([]Request_Manager.Ticket, error) {
	return s.repo.GetUserTickets(ticketID)
}

func (s *TicketService) DeleteTicket(ticketID int) error {
	return s.repo.DeleteTicket(ticketID)
}
func (s *TicketService) UpdateTicket(ticketID int, input Request_Manager.UpdateTicketInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.UpdateTicket(ticketID, input)
}
