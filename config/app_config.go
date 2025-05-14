package config

import (
	"github.com/gin-gonic/gin"
)

const (
	serverPort = "8080"

	logLevel = "info"
)

type ApplicationConfig struct {
	Database DatabaseConfig

	ServerPort string

	LogLevel string
}

func InitAppConfig() ApplicationConfig {
	var appConfig ApplicationConfig
	appConfig.Database = InitDatabaseConfig()

	appConfig.ServerPort = serverPort

	appConfig.LogLevel = logLevel

	return appConfig

}

func AppStart() {
	appConfig := InitAppConfig()

	ConnectAndCheckDB(appConfig.Database)

	// Start the server here
	r := gin.Default()

	r.Run(":" + appConfig.ServerPort)

}
