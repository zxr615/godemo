package main

import (
	"g/study/common"
	"g/study/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode("debug")
	engine := gin.Default()

	common.InitDb()
	engine.POST("/register", controller.Register)
	panic(engine.Run(":8988"))
}
