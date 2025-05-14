package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	id          uint `gorm:"primaryKey"`
	username    string
	password    string
	create_time time.Time
	update_time time.Time
}

func GetAllUser(db *gorm.DB) ([]User, error) {
	var users []User
	result := db.Find(&users)
	return users, result.Error
}

func GetUserByID(db *gorm.DB, id uint) (*User, error) {
	var user User
	result := db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
