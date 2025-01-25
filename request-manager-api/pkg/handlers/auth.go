package handlers

import "github.com/gin-gonic/gin"

func (h *Handlers) register(c *gin.Context) {
	c.JSON(200, gin.H{"message": "register endpoint"})
}
func (h *Handlers) registerAdmin(c *gin.Context) {
	c.JSON(200, gin.H{"message": "registerAdmin endpoint"})
}
func (h *Handlers) login(c *gin.Context) {
	c.JSON(200, gin.H{"message": "login endpoint"})
}
