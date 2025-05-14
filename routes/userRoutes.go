package routes

import (
	"go_backend/services"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine, services *services.UserService) {
	userGroup := r.Group("/user")
	{
		userGroup.GET("", func(c *gin.Context) {
			users, err := services.GetAllUsers()
			if err != nil {
				c.JSON(500, gin.H{"error": "Failed to query users"})
			}
			c.JSON(200, users)
		})
	}
}
