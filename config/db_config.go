package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string

	MaxIdleConns int
	MaxOpenConns int
}

func InitDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		Host:         "127.0.0.1",
		Port:         3306,
		User:         "root",
		Password:     "yyl19981127",
		DBName:       "testdb",
		MaxIdleConns: 10,
		MaxOpenConns: 100,
	}
}

func ConnectDB(config DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User, config.Password, config.Host, config.Port, config.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	return db, nil
}

func ConnectAndCheckDB(config DatabaseConfig) *gorm.DB {
	db, err := ConnectDB(config)
	if err != nil {
		panic("Failed to connect database: " + err.Error())
	}
	// 可以在这里添加连接成功后的其他逻辑
	fmt.Println("Database connection successful")
	return db
}
