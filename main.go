package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"math/rand"
	"net/http"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Nickname string `gorm:"type:varchar(20);not null"`
	Phone    string `gorm:"varchar(10);not null;unique"`
	Pwd      string `gorm:"size:255"`
}

func main() {
	db := InitDb()
	r := gin.Default()
	r.POST("/register", func(ctx *gin.Context) {
		// 获取参数
		name := ctx.PostForm("nickname")
		phone := ctx.PostForm("phone")
		pwd := ctx.PostForm("pwd")

		// 数据验证
		if len(phone) != 11 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"msg": "手机号必须为11位",
			})
			return
		}

		if len(pwd) < 6 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"msg": "密码不能小于6位",
			})
			return
		}

		if len(name) == 0 {
			name = RandomString(10)
		}

		// 创建用户
		if isPhoneExists(db, phone) {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"msg": "手机已存在",
			})
			return
		}

		user := User{
			Nickname: name,
			Phone:    phone,
			Pwd:      pwd,
		}

		_ = createUser(db, user)

		// 返回结果
		ctx.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
		})

	})
	panic(r.Run(":8988"))
}

func createUser(db *gorm.DB, user User) *gorm.DB {
	return db.Create(&user)
}

func isPhoneExists(db *gorm.DB, phone string) bool {
	var user User
	db.Where("phone = ?", phone).First(&user)
	if user.ID != 0 {
		return true
	}

	return false
}

func InitDb() *gorm.DB {
	host := "192.168.88.116"
	port := "3306"
	database := "study"
	username := "root"
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

	err = db.AutoMigrate(&User{})
	if err != nil {
		panic("failed migrate" + err.Error())
	}

	return db
}

func RandomString(n int) string {
	var letters = "qwertyuiopasdfghjklzxcvbnm"
	result := make([]byte, n)

	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}

	return string(result)
}
