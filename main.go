package main

import (
	"flag"

	"github.com/NothAmor/gin-hello/app"
)

func main() {
	var root string
	flag.StringVar(&root, "appRoot", "", "Usage: -appRoot=/www/gin-hello       Define Application main directory")
	flag.Parse()
	app.Init(root)
}
