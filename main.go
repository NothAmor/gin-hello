package main

import (
	"flag"

	"petehouston.com/gin-hello/app"
)

func main() {
	var root string
	flag.StringVar(&root, "appRoot", "", "Usage: -appRoot=/www/gin-hello       Define Application main directoryapplicationRoot")
	flag.Parse()
	app.Init(root)
}
