package main

import (
	"flag"

	"github.com/NothAmor/gin-hello/app"
	"github.com/NothAmor/gin-hello/app/database"
)

var (
	root    string
	migrate bool
)

func main() {
	// 指定程序运行目录，必须参数
	flag.StringVar(&root, "appRoot", "", "Usage: --appRoot=/www/gin-hello       Define Application main directory")

	// 是否migrate数据库，非必需，但首次运行时必须携带此参数
	flag.BoolVar(&migrate, "migrate", false, "Usage: --migrate                  Migrate the database.")
	flag.Parse()

	// 如果有migrate参数，则migrate数据库
	database.DatabaseMigrate(database.InitDatabase(), migrate)

	// 启动系统
	app.Init(root)
}
