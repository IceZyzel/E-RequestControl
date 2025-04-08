package handlers

import (
	"net/http"
	"request_manager_api/pkg/services"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	service *services.Service
}

func NewHandler(services *services.Service) *Handlers {
	return &Handlers{service: services}
}

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Разрешаем запросы от фронтенда через ingress
        c.Writer.Header().Set("Access-Control-Allow-Origin", c.GetHeader("Origin"))
        
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization, Accept, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length, Content-Range, Authorization")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Max-Age", "86400") // 24 часа

        if c.Request.Method == http.MethodOptions {
            c.AbortWithStatus(http.StatusNoContent)
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
    
    // Группа для /api - все API endpoints теперь будут начинаться с /api
    api := router.Group("/api")
    {
        // Аутентификация
        auth := api.Group("/auth")
        {
            auth.POST("/register", h.register)
            auth.POST("/registerAdmin", h.registerAdmin)
            auth.POST("/login", h.login)
            auth.POST("/logout", h.logout)
        }

        // Пользовательские endpoints (требуют аутентификации)
        user := api.Group("", h.userIdentity)
        {
            user.GET("/users", h.getAllUsers)
            
            tickets := user.Group("/tickets")
            {
                tickets.GET("/", h.getUserTickets)
                tickets.POST("/", h.createTicket)
                tickets.PUT("/:ticketID", h.updateTicket)
                tickets.DELETE("/:ticketID", h.deleteTicket)
            }
            
            notifications := user.Group("/notifications")
            {
                notifications.GET("/", h.getUserNotifications)
                notifications.DELETE("/:notificationID", h.markNotificationAsRead)
            }
        }

        // Админские endpoints
        admin := api.Group("/admin", h.userIdentity, h.adminRequired)
        {
            adminTickets := admin.Group("/tickets")
            {
                adminTickets.GET("/", h.getTickets)
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
    }

    // Health check endpoint
    router.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"status": "ok"})
    })

    return router
}
