package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	Request_Manager "request_manager_api"
	"time"
)

type AdminMysql struct {
	db *sqlx.DB
}

func NewAdminMysql(db *sqlx.DB) *AdminMysql {
	return &AdminMysql{db: db}
}

func (r *AdminMysql) GetUserByID(userID int) (Request_Manager.User, error) {
	var user Request_Manager.User
	query := `SELECT UserID, Username, Password, FirstName, LastName, Email, RoleID, CreatedAt, UpdatedAt 
	          FROM User WHERE UserID = ?`
	err := r.db.Get(&user, query, userID)
	if err != nil {
		log.Printf("Error fetching user by ID %d: %s", userID, err)
	}
	return user, err
}
func (r *AdminMysql) GetAllUsers() ([]Request_Manager.User, error) {
	var users []Request_Manager.User
	query := fmt.Sprintf("SELECT UserID, Username, Password, FirstName, LastName, Email, RoleID, CreatedAt, UpdatedAt FROM User")
	err := r.db.Select(&users, query)
	if err != nil {
		log.Printf("Error fetching all users: %s", err)
	}
	return users, err
}
func (r *AdminMysql) Delete(UserID int) error {
	query := fmt.Sprintf("DELETE FROM User WHERE UserID = ? ")
	_, err := r.db.Exec(query, UserID)
	return err
}

func (r *AdminMysql) CreateUser(user Request_Manager.User) (int, error) {
	checkQuery := fmt.Sprintf("SELECT COUNT(*) FROM User WHERE email=? OR username=?")
	var count int
	err := r.db.Get(&count, checkQuery, user.Email, user.Username)
	if err != nil {
		return 0, err
	}

	if count > 0 {
		return 0, fmt.Errorf("user with this email or username already exists")
	}

	createdAt := time.Now().UTC().Add(2 * time.Hour).Format("2006-01-02 15:04:05")

	query := fmt.Sprintf(`INSERT INTO User (Username, Email, Password, RoleID, FirstName, LastName, CreatedAt, UpdatedAt) 
	          VALUES (?, ?, ?, ?, ?, ?, ?, ?)`)

	result, err := r.db.Exec(query, user.Username, user.Email, user.Password, user.RoleID, user.FirstName, user.LastName, createdAt, createdAt)

	if err != nil {
		mysqlErr, ok := err.(*mysql.MySQLError)
		if ok && mysqlErr.Number == 1062 {
			return 0, fmt.Errorf("user with this email or username already exists")
		}
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *AdminMysql) UpdateUser(UserID int, input Request_Manager.UpdateUserInput, user Request_Manager.User) error {
	existingUser, err := r.GetUserByID(UserID)
	if err != nil {
		return err
	}

	updatedAt := time.Now().UTC().Add(2 * time.Hour).Format("2006-01-02 15:04:05")

	if input.Email != nil {
		checkEmailQuery := "SELECT UserID FROM User WHERE Email=? AND UserID != ?"
		var otherUserID int
		err := r.db.Get(&otherUserID, checkEmailQuery, *input.Email, UserID)
		if err != nil && err != sql.ErrNoRows {
			return err
		}
		if otherUserID != 0 {
			return errors.New("user with this email already exists")
		}
		existingUser.Email = *input.Email
	}

	if input.Password != nil {
		existingUser.Password = *input.Password
	}

	if input.Username != nil {
		checkUsernameQuery := "SELECT UserID FROM User WHERE Username=? AND UserID != ?"
		var otherUserID int
		err := r.db.Get(&otherUserID, checkUsernameQuery, *input.Username, UserID)
		if err != nil && err != sql.ErrNoRows {
			return err
		}
		if otherUserID != 0 {
			return errors.New("user with this username already exists")
		}
		existingUser.Username = *input.Username
	}

	if input.RoleID != nil {
		existingUser.RoleID = *input.RoleID
	}
	if input.FirstName != nil {
		existingUser.FirstName = *input.FirstName
	}
	if input.LastName != nil {
		existingUser.LastName = *input.LastName
	}

	query := `UPDATE User 
	          SET Username=?, Email=?, Password=?, RoleID=?, FirstName=?, LastName=?, UpdatedAt=? 
	          WHERE UserID=?`
	_, err = r.db.Exec(query, existingUser.Username, existingUser.Email, *input.Password, existingUser.RoleID, existingUser.FirstName, existingUser.LastName, updatedAt, UserID)
	if err != nil {
		return err
	}

	return nil
}
