package common

import (
	"fmt"
	"g/study/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDb() *gorm.DB {
	//host := "192.168.232.129"
	host := "db.tuju.cn"
	port := "3306"
	database := "study"
	username := "root"
	//password := "fanwei"
	password := "AwytEJ1puYBlcy5K"
	charset := "utf8mb4"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True",
		username,
		password,
		host,
		port,
		database,
		charset)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		panic("failed migrate" + err.Error())
	}

	DB = db
	return db
}
