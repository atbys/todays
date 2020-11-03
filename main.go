package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.LoadHTMLGlob("./template/*")
	engine.GET("/", func(c *gin.Context) {
		render(c, gin.H{
			"description": 1907,
		}, "index.html")
	})
	engine.Run(":8080")
}
