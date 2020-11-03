package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func showIndex(c *gin.Context) {
	render(c, gin.H{
		"description": 1907,
	}, "index.html")
}

func showPickUpMovie(c *gin.Context) {
	id := "83583" // 本来はコンテクストのパラメータから取得
	m := GetMovieInfo(id)
	fmt.Println("[+] Got movie infomation")
	render(c, gin.H{
		"title":   m.Titile,
		"rate":    m.Rate,
		"reviews": m.Reviews,
	}, "pickup_movie.html")
}
