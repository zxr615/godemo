package controller

import (
	"g/study/common"
	"g/study/models"
	"g/study/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Register(ctx *gin.Context) {
	// 获取参数
	name := ctx.PostForm("nickname")
	phone := ctx.PostForm("phone")
	pwd := ctx.PostForm("pwd")

	// 数据验证
	if len(phone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": "手机号必须为11位.",
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
		name = utils.RandomString(10)
	}

	existsPhone := models.IsPhoneExists(common.DB, phone)
	if existsPhone {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": "手机已存在",
		})
		return
	}

	bcryptPwd, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)

	user := models.User{
		Nickname: name,
		Phone:    phone,
		Pwd:      string(bcryptPwd),
	}

	models.CreateUser(common.DB, user)

	// 返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
	})
}

func Login(ctx *gin.Context) {
	phone := ctx.PostForm("phone")
	pwd := ctx.PostForm("pwd")

	userInfo := models.GetUserInfo(common.DB, phone)

	err := bcrypt.CompareHashAndPassword([]byte(userInfo.Pwd), []byte(pwd))
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": "密码错误" + err.Error(),
		})
		return
	}

	token := "test"

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": token,
	})
}
