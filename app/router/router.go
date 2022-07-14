package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init(engine *gin.Engine) {
	engine.GET("/login", func(context *gin.Context) {
		context.HTML(http.StatusOK, "login.html", gin.H{
			"appName":  "TicketAware",
			"pageName": "登陆",
		})
	})
	engine.GET("/register", func(context *gin.Context) {
		context.HTML(http.StatusOK, "register.html", "")
	})
}
