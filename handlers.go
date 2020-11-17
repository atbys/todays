package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const allMovieNum = 85000

func showIndex(c *gin.Context) {
	render(c, gin.H{
		"description": 1907,
	}, "index.html")
}

func showPickUpMovie(c *gin.Context) {
	//id := "83583" // 本来はコンテクストのパラメータから取得
	id := c.Param("id")
	m := GetMovieInfo(id)
	fmt.Println("[+] Got movie information")
	render(c, gin.H{
		"title":   m.Titile,
		"link":    m.FLink,
		"rate":    m.Rate,
		"reviews": m.Reviews,
	}, "pickup_movie.html")
}

func showLoggedIn(c *gin.Context) {

	userID, _ := c.Get("UserID")
	log.Println("[+] login is seccess: user_id is ", userID)
	render(c, gin.H{
		"user_id": userID,
	}, "loggedin")
}

func randMovie(c *gin.Context) {
	movieID := randInt(allMovieNum)
	c.Redirect(http.StatusFound, "/pickup_movie/"+strconv.Itoa(movieID))
}

func randMovieFromClip(c *gin.Context) {
	render(c, gin.H{}, "input_id.html")
}
func randClipMovie(c *gin.Context) {
	filmarks_id := c.Query("filmarks_id")
	//filmarks_id := "nekoneon"
	movieIds := GetIdOfClipMovies(filmarks_id)
	num := len(movieIds)
	log.Println("[+] movie num : ", num)
	movieID := movieIds[randInt(num)]
	c.Redirect(http.StatusFound, "/pickup_movie/"+movieID)
}
