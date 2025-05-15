// routes/user_routes.go
package routes

import (
	"go_backend/common"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserRouteLoader 用户模块路由加载器
type UserRouteLoader struct{}

// NewUserRouteLoader 创建用户路由加载器
func NewUserRouteLoader() *UserRouteLoader {
	return &UserRouteLoader{}
}

// LoadRoutes 实现RouteLoader接口
func (u *UserRouteLoader) LoadRoutes(r *gin.Engine, container *common.ServiceContainer) {
	userGroup := r.Group("/user")
	{
		userGroup.GET("", func(c *gin.Context) {
			users, err := container.UserService.GetAllUsers()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
				return
			}
			c.JSON(http.StatusOK, users)
		})

		userGroup.GET("/:id", func(c *gin.Context) {
			id := c.Param("id")
			userID, err := strconv.ParseUint(id, 10, 64) // 将字符串转换为uint64，然后根据需要转换为uint
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
				return
			}
			user, err := container.UserService.GetUserByID(uint(userID)) // 转换为uint类型
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
				return
			}
			c.JSON(http.StatusOK, user)
		})

		/*
		   userGroup.POST("", func(c *gin.Context) {
		       var user services.CreateUserDTO
		       if err := c.ShouldBindJSON(&user); err != nil {
		           c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		           return
		       }

		       createdUser, err := container.UserService.CreateUser(user)
		       if err != nil {
		           c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		           return
		       }

		       c.JSON(http.StatusCreated, createdUser)
		   })

		   // 其他用户相关路由...

		*/
	}
}
