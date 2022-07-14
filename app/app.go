package app

import (
	"net/http"

	"github.com/NothAmor/gin-hello/app/router"

	"github.com/gin-gonic/gin"
)

func Init(appRoot string) *gin.Engine {
	// 初始化Gin引擎
	engine := gin.New()

	// 初始化模版
	loadTemplates(appRoot, engine)

	// 配置路由
	router.Init(engine)

	engine.Run()

	return engine
}

func loadTemplates(appRoot string, engine *gin.Engine) {
	engine.StaticFS("/assets", http.Dir(appRoot+"/assets"))
	engine.LoadHTMLGlob(appRoot + "/templates/*")
}
