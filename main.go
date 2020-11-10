package main

import (
	"g/study/common"
	"g/study/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	common.InitDb()
	r.POST("/register", controller.Register)
	panic(r.Run(":8988"))
}
