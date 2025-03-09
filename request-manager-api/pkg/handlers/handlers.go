package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"request_manager_api/pkg/services"
)

type Handlers struct {
	service *services.Service
}

func NewHandler(services *services.Service) *Handlers {
	return &Handlers{service: services}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		allowedOrigins := []string{
			"http://localhost:5173",
			"http://localhost:3000",
		}

		origin := c.GetHeader("Origin")
		if contains(allowedOrigins, origin) {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		} else {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173") // Default fallback
		}

		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")

		c.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")

		c.Writer.Header().Set("Access-Control-Expose-Headers", "Authorization")

		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func (h *Handlers) InitRoutes() *gin.Engine {
	router := gin.Default()
	router.Use(CORSMiddleware())

	auth := router.Group("/auth")
	{
		auth.POST("/register", h.register)
		auth.POST("/registerAdmin", h.registerAdmin)
		auth.POST("/login", h.login)
	}

	api := router.Group("/api", h.userIdentity)
	{
		api.GET("/users", h.getAllUsers)
		tickets := api.Group("/tickets")
		{
			tickets.GET("/", h.getUserTickets)
			tickets.POST("/", h.createTicket)
			tickets.PUT("/:ticketID", h.updateTicket)
			tickets.DELETE("/:ticketID", h.deleteTicket)
		}
		notifications := api.Group("/notifications")
		{
			notifications.GET("/", h.getUserNotifications)
			notifications.DELETE("/:notificationID", h.markNotificationAsRead)
		}
	}

	admin := router.Group("/admin", h.userIdentity, h.adminRequired)
	{
		adminTickets := admin.Group("/tickets")
		{
			adminTickets.GET("/", h.getAllTickets)
			adminTickets.GET("/:ticketID", h.getTicketByID)
			adminTickets.DELETE("/:ticketID", h.adminDeleteTicket)
		}
		adminNotifications := admin.Group("/notifications")
		{
			adminNotifications.GET("/", h.getAllNotifications)
			adminNotifications.POST("/", h.createNotification)
			adminNotifications.DELETE("/:notificationID", h.deleteNotification)
		}

		users := admin.Group("/users")
		{
			users.GET("/", h.getAllUsers)
			users.POST("/", h.createUser)
			users.GET("/:userID", h.getUserByID)
			users.PUT("/:userID", h.updateUser)
			users.DELETE("/:userID", h.deleteUser)
		}

		data := admin.Group("/data")
		{
			data.POST("/backup", h.backupData)
			data.POST("/restore", h.restoreData)
			data.GET("/export", h.exportData)
			data.POST("/import", h.importData)
		}
	}

	return router
}
