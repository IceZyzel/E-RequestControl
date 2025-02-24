package request_manager_api

import (
	"errors"
	"regexp"
)

type User struct {
	UserID    int    `json:"user_id" db:"UserID"`
	Username  string `json:"username" db:"Username"`
	Password  string `json:"password" db:"Password"`
	FirstName string `json:"firstname" db:"FirstName"`
	LastName  string `json:"lastname" db:"LastName"`
	Email     string `json:"email" db:"Email"`
	RoleID    int    `json:"role_id" db:"RoleID"`
	CreatedAt string `json:"created_at" db:"CreatedAt"`
	UpdatedAt string `json:"updated_at" db:"UpdatedAt"`
}

type UpdateUserInput struct {
	UserID    *int    `json:"user_id"`
	Username  *string `json:"username"`
	Password  *string `json:"password"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Email     *string `json:"email"`
	RoleID    *int    `json:"role_id"`
	CreatedAt *string `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
}

func (i UpdateUserInput) Validate() error {
	if i.UserID == nil && i.Username == nil && i.Email == nil && i.Password == nil && i.RoleID == nil &&
		i.FirstName == nil && i.LastName == nil && i.CreatedAt == nil && i.UpdatedAt == nil {
		return errors.New("Update structure has no value")
	}
	return nil
}
func (u *User) ValidateEmail() error {
	if u.Email != "" {
		emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
		match, _ := regexp.MatchString(emailRegex, u.Email)
		if !match {
			return errors.New("invalid email format")
		}
	}
	return nil
}
func (u *User) ValidatePassword() error {
	if u.Password != "" {
		if len(u.Password) < 8 {
			return errors.New("password should be at least 8 characters long")
		}

		letterRegex := regexp.MustCompile(`[a-zA-Z]`)
		if !letterRegex.MatchString(u.Password) {
			return errors.New("password should contain at least 1 letter")
		}

		digitRegex := regexp.MustCompile(`\d`)
		if !digitRegex.MatchString(u.Password) {
			return errors.New("password should contain at least 1 digit")
		}

		specialCharRegex := regexp.MustCompile(`[^a-zA-Z0-9\s]`)
		if !specialCharRegex.MatchString(u.Password) {
			return errors.New("password should contain at least 1 special character")
		}
	}
	return nil
}
