package router

import (
	"github.com/NothAmor/gin-hello/app/controller"

	"github.com/gin-gonic/gin"
)

func Init(engine *gin.Engine) {
	engine.GET("/login", controller.Login)
	engine.GET("/register", controller.Register)
}
