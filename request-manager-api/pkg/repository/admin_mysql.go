package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	Request_Manager "request_manager_api"
)

type AdminMysql struct {
	db *sqlx.DB
}

func NewAdminMysql(db *sqlx.DB) *AdminMysql {
	return &AdminMysql{db: db}
}
func (r *AdminMysql) GetUserByID(userID int) (Request_Manager.User, error) {
	var user Request_Manager.User
	query := fmt.Sprintf("SELECT UserID, Username, Email, Role FROM User WHERE UserID=?")
	err := r.db.Get(&user, query, userID)
	if err != nil {
		log.Printf("Error fetching user by ID %d: %s", userID, err)
	}
	return user, err
}
