package controller

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"

	"github.com/NothAmor/gin-hello/app/database"

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

func LoginRequest(context *gin.Context) {
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

func RegisterRequest(context *gin.Context) {
	// 配置数据库连接
	db := database.InitDatabase()
	defer db.Close()

	// 获取post表单数据
	username := context.PostForm("username")
	password := context.PostForm("password")

	// 将密码加密为md5
	h := md5.New()
	h.Write([]byte(password))
	password = hex.EncodeToString(h.Sum(nil))

	// 获取访问IP
	reqIP := context.ClientIP()
	if reqIP == "::1" {
		reqIP = "127.0.0.1"
	}

	// 插入数据到mysql中
	rs, err := db.Exec("insert into ta_users(username, password, ip, balance) values(?, ?, ?, 0);", username, password, reqIP)
	if err != nil {
		panic(err)
	}

	// 获取用户的ID值
	id, err := rs.LastInsertId()
	if err != nil {
		panic(err)
	}

	// 返回json串
	context.JSON(http.StatusOK, gin.H{
		"code":   200,
		"userId": id,
		"msg":    "Register request successful!",
	})
}
