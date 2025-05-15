package config

import (
	"fmt"

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

	db := ConnectAndCheckDB(appConfig.Database)

	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			fmt.Println("Error getting sql.DB object for closing:", err)
			return
		}
		err = sqlDB.Close()
		if err != nil {
			fmt.Println("Error closing database connection:", err)
		}
	}()

	// Start the server here
	r := gin.Default()

	RegisterRoutes(r, db)

	r.Run(":" + appConfig.ServerPort)

}
