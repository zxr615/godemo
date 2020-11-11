package main

import (
	"g/study/common"
	"g/study/router"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode("debug")

	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery())

	router.SetApiRoute(engine)

	common.InitDb()

	panic(engine.Run(":8980"))
}
