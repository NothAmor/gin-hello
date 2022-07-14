package router

import (
	"github.com/NothAmor/gin-hello/app/controller"

	"github.com/gin-gonic/gin"
)

func Init(engine *gin.Engine) {
	engine.GET("/", controller.Index)
	engine.GET("/login", controller.Login)
	engine.GET("/register", controller.Register)

	engine.POST("/user/login", controller.LoginRequest)
	engine.POST("/user/register", controller.RegisterRequest)
}
