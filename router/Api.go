package router

import (
	"g/study/controller"
	"github.com/gin-gonic/gin"
)

func SetApiRoute(engine *gin.Engine) {

	userRoute := engine.Group("/user")
	{
		userRoute.POST("/register", controller.Register)
		userRoute.POST("login", controller.Login)
	}
}
