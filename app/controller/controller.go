package controller

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/NothAmor/gin-hello/app/database"
	"github.com/NothAmor/gin-hello/app/models"
	"github.com/NothAmor/gin-hello/app/security"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

type userInfo struct {
	username string
	email    string
}

func cookieValidator(context *gin.Context) (user *userInfo, status bool) {
	// 如果cookie中有token，那就解析这个token，token有效并且还是请求的登陆后不能请求的页面就跳转首页，其余的正常
	cookie, err := context.Cookie("token")
	if err != nil {
		fmt.Println(err)
	}
	username, email, ok := security.Validator(cookie)

	user = new(userInfo)
	if ok {
		status = true
		user.email = email
		user.username = username
		return user, true
	} else {
		return user, false
	}
}

func Index(context *gin.Context) {
	user, status := cookieValidator(context)
	if status {
		context.HTML(http.StatusOK, "index.html", gin.H{
			"appName":  "TicketAware",
			"username": user.username,
			"email":    user.email,
			"pageName": "首页",
		})
	} else {
		context.HTML(http.StatusOK, "index.html", gin.H{
			"appName":  "TicketAware",
			"pageName": "首页",
		})
	}
}

func Login(context *gin.Context) {
	user, status := cookieValidator(context)
	if status {
		context.HTML(http.StatusOK, "login.html", gin.H{
			"appName":  "TicketAware",
			"username": user.username,
			"email":    user.email,
			"pageName": "登陆",
		})
	} else {
		context.HTML(http.StatusOK, "login.html", gin.H{
			"appName":  "TicketAware",
			"pageName": "登陆",
		})
	}
}

func Register(context *gin.Context) {
	user, status := cookieValidator(context)
	if status {
		context.HTML(http.StatusOK, "register.html", gin.H{
			"appName":  "TicketAware",
			"username": user.username,
			"email":    user.email,
			"pageName": "注册",
		})
	} else {
		context.HTML(http.StatusOK, "register.html", gin.H{
			"appName":  "TicketAware",
			"pageName": "注册",
		})
	}
}

func LoginRequest(context *gin.Context) {
	var (
		email    string
		password string
		msg      string
		token    string
		code     int
		userInfo *models.Users
		ok       bool
	)

	email = context.PostForm("email")
	password = context.PostForm("password")

	userInfo, token, ok = security.SignIn(security.LoginInfo{Email: email, Password: password})

	if ok {
		msg = "登陆成功！"
		code = 200
	} else {
		msg = "登陆失败，请检查登陆信息！"
		code = 500
	}

	context.JSON(http.StatusOK, gin.H{
		"code":  code,
		"msg":   msg,
		"token": token,
		"userInfo": gin.H{
			"email":    userInfo.Email,
			"username": userInfo.Username,
			"balance":  userInfo.Balance,
		},
	})
}

func RegisterRequest(context *gin.Context) {
	// 配置数据库连接
	db := database.InitDatabase()

	// 获取post表单数据
	email := context.PostForm("email")
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

	var msg string
	var code int

	// 插入数据到mysql中
	userRegister := models.Users{
		Email:        email,
		Username:     username,
		Password:     password,
		Ip:           reqIP,
		Balance:      0,
		RegisterTime: time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := db.Create(&userRegister).Error; err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			msg = "邮箱已注册，请使用其它邮箱，注册失败！"
			code = 500
		}
	} else {
		msg = "注册成功，即将跳转到登陆页面！"
		code = 200
	}

	// 返回json串
	context.JSON(http.StatusOK, gin.H{
		"code":   code,
		"userId": userRegister.ID,
		"msg":    msg,
	})
}
