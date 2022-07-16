package router

import (
	"github.com/NothAmor/gin-hello/app/controller"

	"github.com/gin-gonic/gin"
)

func Init(engine *gin.Engine) {
	engine.GET("/", controller.Index)

	engine.GET("/user/login", controller.Login)
	engine.GET("/user/register", controller.Register)

	engine.POST("/user/loginEntry", controller.LoginRequest)
	engine.POST("/user/registerEntry", controller.RegisterRequest)
	engine.POST("/user/logout", controller.Logout)
}
