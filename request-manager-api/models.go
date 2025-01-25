package request_manager_api

type Ticket struct {
	TicketID    int    `json:"TicketID" db:"TicketID"`
	Title       string `json:"Title" binding:"required" db:"Title"`
	Description string `json:"Description" db:"Description"`
	Status      string `json:"Status" db:"Status"`
	CreatedAt   string `json:"CreatedAt" db:"CreatedAt"`
	UpdatedAt   string `json:"UpdatedAt" db:"UpdatedAt"`
	AssignedTo  int    `json:"AssignedTo" db:"AssignedTo"`
	UserID      int    `json:"UserID" db:"UserID"`
}

type TicketStatus struct {
	StatusID  int    `json:"StatusID" db:"StatusID"`
	Status    string `json:"Status" db:"Status"`
	CreatedAt string `json:"CreatedAt" db:"CreatedAt"`
	UpdatedAt string `json:"UpdatedAt" db:"UpdatedAt"`
}

type Role struct {
	RoleID    int    `json:"RoleID" db:"RoleID"`
	RoleName  string `json:"RoleName" db:"RoleName"`
	CreatedAt string `json:"CreatedAt" db:"CreatedAt"`
	UpdatedAt string `json:"UpdatedAt" db:"UpdatedAt"`
}

type Notification struct {
	NotificationID int    `json:"NotificationID" db:"NotificationID"`
	Message        string `json:"Message" db:"Message"`
	UserID         int    `json:"UserID" db:"UserID"`
	CreatedAt      string `json:"CreatedAt" db:"CreatedAt"`
	UpdatedAt      string `json:"UpdatedAt" db:"UpdatedAt"`
}
