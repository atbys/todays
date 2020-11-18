package main

import (
	"flag"
	"os"

	"github.com/gin-gonic/gin"
)

var engine *gin.Engine
var (
	daemon = flag.Bool("d", false, "daemon")
)

func main() {
	engine = gin.Default()
	gopath := os.Getenv("GOPATH")
	engine.LoadHTMLGlob(gopath + "/src/github.com/atbys/todays/template/*")

	initializeSession() //セッション管理の初期化の関係で絶対にこっちを先に処理する
	initializeRoutes()
	flag.Parse()
	if *daemon {
		engine.Run(":80")
	} else {
		engine.Run(":8080")
	}
}
