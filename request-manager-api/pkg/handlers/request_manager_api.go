package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	Request_Manager "request_manager_api"
	"strconv"
)

func (h *Handlers) getAllTickets(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "getAllTickets endpoint"})
}

func (h *Handlers) getTicketByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "getTicketByID endpoint"})
}
func (h *Handlers) adminDeleteTicket(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "getTicketByID endpoint"})
}
func (h *Handlers) adminUpdateTicket(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "getTicketByID endpoint"})
}
func (h *Handlers) getUserTickets(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user ID"})
		return
	}
	testResults, err := h.service.Ticket.GetUserTickets(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, testResults)
}
func (h *Handlers) createTicket(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user ID"})
		return
	}

	var ticket Request_Manager.Ticket
	if err := c.ShouldBindJSON(&ticket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ticket.UserID = userID

	id, err := h.service.CreateTicket(ticket)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *Handlers) updateTicket(c *gin.Context) {
	ticketID := c.Param("ticketID")
	id, err := strconv.Atoi(ticketID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id param")
		return
	}

	var input Request_Manager.UpdateTicketInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Ticket.UpdateTicket(id, input)

	c.JSON(http.StatusOK, statusResponse{
		Status: "Ok",
	})
}

func (h *Handlers) deleteTicket(c *gin.Context) {
	ticketID, err := strconv.Atoi(c.Param("ticketID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ticket ID"})
		return
	}
	if err := h.service.Ticket.DeleteTicket(ticketID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Ticket deleted"})
}
func (h *Handlers) getUserNotifications(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "deleteTicket endpoint"})
}
func (h *Handlers) markNotificationAsRead(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "deleteTicket endpoint"})
}
func (h *Handlers) getTicketStatuses(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "getTicketStatuses endpoint"})
}
func (h *Handlers) getAllStatuses(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "updateStatus endpoint"})
}

func (h *Handlers) getStatusByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "updateStatus endpoint"})
}

func (h *Handlers) createStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "updateStatus endpoint"})
}

func (h *Handlers) updateStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "updateStatus endpoint"})
}

func (h *Handlers) deleteStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "updateStatus endpoint"})
}

func (h *Handlers) getAllUsers(c *gin.Context) {
	users, err := h.service.Admin.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}
func (h *Handlers) createUser(c *gin.Context) {
	var user Request_Manager.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
		return
	}
	id, err := h.service.Admin.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"UserID": id,
	})
}
func (h *Handlers) getUserByID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user ID"})
		return
	}
	user, err := h.service.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *Handlers) updateUser(c *gin.Context) {
	userIDStr := c.Param("userID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid UserID")
		return
	}
	var input Request_Manager.UpdateUserInput
	var userInput Request_Manager.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	err = h.service.Admin.UpdateUser(userID, input, userInput)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "Ok",
	})
}

func (h *Handlers) deleteUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user ID"})
		return
	}
	if err := h.service.Admin.Delete(userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

func (h *Handlers) getUserRoles(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "getUserRoles endpoint"})
}

func (h *Handlers) updateUserRole(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "updateUserRole endpoint"})
}

func (h *Handlers) createNotification(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "createNotification endpoint"})
}
func (h *Handlers) getAllNotifications(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "createNotification endpoint"})
}
func (h *Handlers) deleteNotification(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "createNotification endpoint"})
}
func (h *Handlers) backupData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "createNotification endpoint"})
}
func (h *Handlers) restoreData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "createNotification endpoint"})
}
func (h *Handlers) exportData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "createNotification endpoint"})
}
func (h *Handlers) importData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "createNotification endpoint"})
}
