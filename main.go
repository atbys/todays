package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

var engine *gin.Engine

func main() {
	engine = gin.Default()
	gopath := os.Getenv("GOPATH")
	engine.LoadHTMLGlob(gopath + "/src/github.com/atbys/todays/template/*")
	//engine.LoadHTMLGlob("./template/*")
	initializeSession() //セッション管理の初期化の関係で絶対にこっちを先に処理する
	initializeRoutes()

	engine.Run(":8080")
}
