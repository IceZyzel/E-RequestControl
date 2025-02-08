package repository

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	Request_Manager "request_manager_api"
	"time"
)

type TicketMysql struct {
	db *sqlx.DB
}

func NewTicketMysql(db *sqlx.DB) *TicketMysql {
	return &TicketMysql{db: db}
}

func (r *TicketMysql) GetUserTickets(userID int) ([]Request_Manager.Ticket, error) {
	var tickets []Request_Manager.Ticket
	query := `
		SELECT t.TicketID, t.Title, t.Description, ts.Status, t.CreatedAt, t.UpdatedAt, 
		       t.AssignedTo, t.UserID, u.Username AS SenderUsername, a.Username AS AssigneeUsername
		FROM Ticket t
		JOIN TicketStatus ts ON t.StatusID = ts.StatusID
		JOIN User u ON t.UserID = u.UserID
		LEFT JOIN User a ON t.AssignedTo = a.UserID
		WHERE t.UserID = ?`

	err := r.db.Select(&tickets, query, userID)
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

func (r *TicketMysql) CreateTicket(ticket Request_Manager.Ticket) (int, error) {
	createdAt := time.Now().UTC().Format("2006-01-02 15:04:05")

	if ticket.AssignedTo == 0 {
		return 0, errors.New("AssignedTo не може бути пустим")
	}

	var assignedTo int
	err := r.db.QueryRow("SELECT UserID FROM User WHERE UserID = ?", ticket.AssignedTo).Scan(&assignedTo)
	if err != nil {
		return 0, errors.New("Обраного користувача не існує")
	}

	var statusID int
	err = r.db.QueryRow("SELECT StatusID FROM TicketStatus WHERE Status = ?", "Новий").Scan(&statusID)

	if err == sql.ErrNoRows {
		_, err := r.db.Exec("INSERT INTO TicketStatus (Status) VALUES (?)", "Новий")
		if err != nil {
			return 0, err
		}

		err = r.db.QueryRow("SELECT StatusID FROM TicketStatus WHERE Status = ?", "Новий").Scan(&statusID)
		if err != nil {
			return 0, err
		}
	} else if err != nil {
		return 0, err
	}

	query := `
		INSERT INTO Ticket (Title, Description, StatusID, CreatedAt, UpdatedAt, AssignedTo, UserID)
		VALUES (?, ?, ?, ?, ?, ?, ?)`
	res, err := r.db.Exec(query, ticket.Title, ticket.Description, statusID, createdAt, createdAt, ticket.AssignedTo, ticket.UserID)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *TicketMysql) UpdateTicket(ticketID int, input Request_Manager.UpdateTicketInput) error {
	updatedAt := time.Now().UTC().Add(2 * time.Hour).Format("2006-01-02 15:04:05")

	existingTicket, err := r.GetByID(ticketID)
	if err != nil {
		return err
	}

	if input.Title != nil {
		existingTicket.Title = *input.Title
	}
	if input.Description != nil {
		existingTicket.Description = *input.Description
	}
	if input.AssignedTo != nil {
		existingTicket.AssignedTo = *input.AssignedTo
	}

	var statusID int
	err = r.db.Get(&statusID, `SELECT StatusID FROM TicketStatus WHERE Status = ?`, "Оновлено")
	if err != nil {
		result, err := r.db.Exec(`INSERT INTO TicketStatus (Status) VALUES (?)`, "Оновлено")
		if err != nil {
			return err
		}
		statusID64, err := result.LastInsertId()
		if err != nil {
			return err
		}
		statusID = int(statusID64)
	}

	query := `UPDATE Ticket SET Title=?, Description=?, AssignedTo=?, StatusID=?, UpdatedAt=? WHERE TicketID=?`
	_, err = r.db.Exec(query, existingTicket.Title, existingTicket.Description, existingTicket.AssignedTo, statusID, updatedAt, ticketID)
	return err
}

func (r *TicketMysql) GetByID(ticketID int) (Request_Manager.Ticket, error) {
	var ticket Request_Manager.Ticket
	query := `SELECT t.TicketID, t.Title, t.Description, ts.Status, t.AssignedTo, t.UserID, t.CreatedAt, t.UpdatedAt
	          FROM Ticket t 
	          JOIN TicketStatus ts ON t.StatusID = ts.StatusID
	          WHERE t.TicketID = ?`
	err := r.db.Get(&ticket, query, ticketID)
	if err != nil {
		return ticket, err
	}
	return ticket, nil
}

func (r *TicketMysql) DeleteTicket(ticketID int) error {
	query := `DELETE FROM Ticket WHERE TicketID=?`
	_, err := r.db.Exec(query, ticketID)
	return err
}
