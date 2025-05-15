package common

import (
	"go_backend/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ServiceContainer 持有所有应用服务
type ServiceContainer struct {
	UserService *services.UserService
	// 可以添加更多服务...
}

func InitSeviceContainer(db *gorm.DB) *ServiceContainer {
	return &ServiceContainer{
		UserService: services.NewUserService(db),
		// 可以添加更多服务...
	}
}

// RouteLoader 定义路由加载接口
type RouteLoader interface {
	LoadRoutes(r *gin.Engine, services *ServiceContainer)
}
