package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", gin.H{
		"appName": "TicketAware",
	})
}

func Login(context *gin.Context) {
	context.HTML(http.StatusOK, "login.html", gin.H{
		"appName":  "TicketAware",
		"pageName": "登陆",
	})
}

func Register(context *gin.Context) {
	context.HTML(http.StatusOK, "register.html", gin.H{
		"appName":  "TicketAware",
		"pageName": "注册",
	})
}
