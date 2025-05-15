package config

import (
	"go_backend/common"
	"go_backend/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RouteRegistry 路由注册中心
type RouteRegistry struct {
	loaders []common.RouteLoader
}

// NewRouteRegistry 创建新的路由注册中心
func NewRouteRegistry() *RouteRegistry {
	return &RouteRegistry{loaders: make([]common.RouteLoader, 0)}
}

// Register 添加路由加载器
func (r *RouteRegistry) Register(loader common.RouteLoader) {
	r.loaders = append(r.loaders, loader)
}

// LoadAllRoutes 加载所有注册的路由
func (r *RouteRegistry) LoadAllRoutes(engine *gin.Engine, services *common.ServiceContainer) {
	for _, loader := range r.loaders {
		loader.LoadRoutes(engine, services)
	}
}

// 注册所有应用路由
func RegisterRoutes(r *gin.Engine, db *gorm.DB) {

	services := common.InitSeviceContainer(db)

	// 创建路由注册中心
	registry := NewRouteRegistry()

	// 注册各个模块的路由加载器
	registry.Register(routes.NewUserRouteLoader())
	// 未来可以添加更多路由加载器: registry.Register(routes.NewProductRouteLoader())

	// 注册路由
	registry.LoadAllRoutes(r, services)
}
