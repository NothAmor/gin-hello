package app

import (
	"fmt"
	"net/http"

	"github.com/NothAmor/gin-hello/app/router"

	"github.com/gin-gonic/gin"
)

func Init(appRoot string) *gin.Engine {
	// 初始化Gin引擎
	engine := gin.Default()

	// 初始化模版
	loadTemplates(appRoot, engine)

	// 配置路由
	router.Init(engine)

	// 启动Gin，并监听8081端口
	engine.Run(":8080")

	return engine
}

func loadTemplates(appRoot string, engine *gin.Engine) {
	engine.StaticFS("/assets", http.Dir(appRoot+"/assets"))
	engine.LoadHTMLGlob(appRoot + "/templates/*")
	fmt.Println(appRoot)
}
