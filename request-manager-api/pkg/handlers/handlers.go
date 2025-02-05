package handlers

import (
	"github.com/gin-gonic/gin"
	"request_manager_api/pkg/services"
)

type Handlers struct {
	service *services.Service
}

func NewHandler(services *services.Service) *Handlers {
	return &Handlers{service: services}
}

func (h *Handlers) InitRoutes() *gin.Engine {
	router := gin.Default()

	auth := router.Group("/auth")
	{
		auth.POST("/register", h.register)
		auth.POST("/registerAdmin", h.registerAdmin)
		auth.POST("/login", h.login)
	}

	api := router.Group("/api")
	{
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
			notifications.PUT("/:notificationID", h.markNotificationAsRead)
		}
	}

	admin := router.Group("/admin", h.userIdentity, h.adminRequired)
	{
		adminTickets := admin.Group("/tickets")
		{
			adminTickets.GET("/", h.getAllTickets)
			adminTickets.GET("/:ticketID", h.getTicketByID)
			adminTickets.PUT("/:ticketID", h.adminUpdateTicket)
			adminTickets.DELETE("/:ticketID", h.adminDeleteTicket)
		}
		adminNotifications := admin.Group("/notifications")
		{
			adminNotifications.GET("/", h.getAllNotifications)
			adminNotifications.POST("/", h.createNotification)
			adminNotifications.DELETE("/:notificationID", h.deleteNotification)
		}
		statuses := admin.Group("/statuses")
		{
			statuses.GET("/", h.getAllStatuses)
			statuses.GET("/:statusID", h.getStatusByID)
			statuses.POST("/", h.createStatus)
			statuses.PUT("/:statusID", h.updateStatus)
			statuses.DELETE("/:statusID", h.deleteStatus)
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
