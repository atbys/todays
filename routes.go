package main

func initializeRoutes() {
	engine.GET("/", showIndex)
	engine.GET("/login", GetLogin)
	engine.POST("/login", PostLogin)
	authGroup := engine.Group("/")
	authGroup.Use(sessionCheck()) //ログイン状態で閲覧するページをこのグループに追加する
	{
		authGroup.GET("/loggedin", showLoggedIn) // テストページ
	}
	engine.POST("/logout", PostLogout)
}
