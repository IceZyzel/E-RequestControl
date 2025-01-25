package handlers

import (
	"github.com/gin-gonic/gin"
)

type Handlers struct{}

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
			tickets.GET("/", h.getAllTickets)
			tickets.GET("/:ticketID", h.getTicketByID)
			tickets.POST("/", h.createTicket)
			tickets.PUT("/:ticketID", h.updateTicket)
			tickets.DELETE("/:ticketID", h.deleteTicket)
		}

		statuses := api.Group("/statuses")
		{
			statuses.GET("/", h.getAllStatuses)
			statuses.GET("/:statusID", h.getStatusByID)
			statuses.POST("/", h.createStatus)
			statuses.PUT("/:statusID", h.updateStatus)
			statuses.DELETE("/:statusID", h.deleteStatus)
		}

		users := api.Group("/users")
		{
			users.GET("/", h.getAllUsers)
			users.GET("/:userID", h.getUserByID)
			users.PUT("/:userID", h.updateUser)
			users.DELETE("/:userID", h.deleteUser)
		}

	}

	return router
}
