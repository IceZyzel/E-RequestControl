package request_manager_api

type User struct {
	UserID    uint   `json:"UserID" db:"UserID"`
	FirstName string `json:"FirstName" db:"FirstName"`
	LastName  string `json:"LastName" db:"LastName"`
	Email     string `json:"Email" db:"Email"`
	Role      string `json:"Role" db:"Role"`
	CreatedAt string `json:"CreatedAt" db:"CreatedAt"`
	UpdatedAt string `json:"UpdatedAt" db:"UpdatedAt"`
}
