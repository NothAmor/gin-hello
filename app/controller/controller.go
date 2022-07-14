package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(context *gin.Context) {
	params := mergeH(context, gin.H{
		"appName":  "TicketAware",
		"pageName": "登陆",
	})
	context.HTML(http.StatusOK, "login.html", params)
}

func Register(context *gin.Context) {
	params := mergeH(context, gin.H{
		"appName":  "TicketAware",
		"pageName": "注册",
	})
	context.HTML(http.StatusOK, "login.html", params)
}

// MergeH 合并默认参数
func mergeH(context *gin.Context, h gin.H) gin.H {
	if context.Keys == nil {
		return h
	}
	if h == nil || len(h) == 0 {
		return context.Keys
	}
	mh := make(gin.H)
	for key, val := range context.Keys {
		mh[key] = val
	}
	for key, val := range h {
		mh[key] = val
	}
	return mh
}
