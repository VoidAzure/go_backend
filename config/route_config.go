package config

import (
	"go_backend/services"

	"github.com/gin-gonic/gin"
)

type ServiceContainer struct {
	UserService *services.UserService
	// 可以添加更多服务...
}

type RouteLoader func(r *gin.Engine, services *ServiceContainer)
