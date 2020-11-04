package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var CookieStore cookie.Store

func initializeSession() {
	CookieStore = cookie.NewStore([]byte("secret"))
	engine.Use(sessions.Sessions("session", CookieStore))
}

func sessionCheck() gin.HandlerFunc { //middleware
	return func(c *gin.Context) {
		log.Println("YEH")
		session := sessions.Default(c)
		UserID := session.Get("UserID")
		log.Println("[+] session check process start: user_id is ", UserID)
		if UserID == nil {
			//TODO
			//今はクッキーに存在しているかどうかの確認だけなので危ない
			//ユーザIDのハッシュをクッキーに保存しておいて，
			//DBに保存してあるユーザIDと比較する
			log.Println("not Loged in")
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
		} else {
			c.Set("UserID", UserID)
			c.Next()
		}
		log.Println("login check is ended")
	}
}

func Login(c *gin.Context, UserID string) { //maybe bad
	log.Println("[+] login process start : user_id is ", UserID)
	session := sessions.Default(c) // here is bag
	log.Println("[+] hey")
	session.Set("UserID", UserID)
	session.Save()
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}

func LoginFromGet(c *gin.Context) { // ok
	render(c, gin.H{
		"UserID":       "",
		"ErrorMessage": "",
	}, "login")
}

func LoginFromPost(c *gin.Context) { //maybe bad
	userID := c.PostForm("userId")
	Login(c, userID)
	c.Redirect(http.StatusFound, "/loggedin") // test
}

func LogoutFromPost(c *gin.Context) {
	Logout(c)
	render(c, gin.H{
		"UserID":       "",
		"ErrorMessage": "",
	}, "login")
}

func SignupFromGet(c *gin.Context) {
	render(c, gin.H{}, "signup.html")
}

func SignupFromPost(c *gin.Context) {
	var form User

	if err := c.Bind(&form); err != nil {
		c.HTML(http.StatusBadRequest, "signup.html", gin.H{"err": err})
		c.Abort()
	} else {
		filmarksID := c.PostForm("filmarks_id")
		password := c.PostForm("password")
		name := c.PostForm("name")
		err := createUser(name, filmarksID, password)
		if err != nil {
			c.HTML(http.StatusBadRequest, "signup.html", gin.H{"err": err})
		}
		c.Redirect(http.StatusFound, "/login")
	}
}

func createUser(name, fid, password string) error {
	endcryptedPassword, _ := PasswordEncrypt(password)
	user := User{
		FilmarksID: fid,
		Password:   endcryptedPassword,
		Name:       name,
	}
	RegistUser(&user)
	return nil
}
