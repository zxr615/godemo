package models

import (
	"errors"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Nickname string `gorm:"type:varchar(20);not null"`
	Phone    string `gorm:"varchar(10);not null;unique"`
	Pwd      string `gorm:"size:255"`
}

func IsPhoneExists(db *gorm.DB, phone string) bool {
	var user User
	result := db.Where("phone = ?", phone).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false
	}

	return true
}

func CreateUser(db *gorm.DB, user User) *gorm.DB {
	return db.Create(&user)
}
