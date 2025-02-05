package request_manager_api

type User struct {
	UserID    int    `json:"UserID" db:"UserID"`
	Username  string `json:"Username" db:"Username"`
	Password  string `json:"Password" db:"Password"`
	FirstName string `json:"FirstName" db:"FirstName"`
	LastName  string `json:"LastName" db:"LastName"`
	Email     string `json:"Email" db:"Email"`
	RoleID    int    `json:"RoleID" db:"RoleID"`
	CreatedAt string `json:"CreatedAt" db:"CreatedAt"`
	UpdatedAt string `json:"UpdatedAt" db:"UpdatedAt"`
}
