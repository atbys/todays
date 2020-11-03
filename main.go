package main

import (
	"github.com/gin-gonic/gin"
)

var engine *gin.Engine

func main() {
	engine = gin.Default()
	engine.LoadHTMLGlob("./template/*")
	initializeRoutes()
	engine.Run(":8080")
}
