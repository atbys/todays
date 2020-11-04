package main

func initializeRoutes() {
	engine.GET("/", showIndex)
	engine.GET("/login", LoginFromGet)
	engine.GET("/pickup_movie", showPickUpMovie)
	engine.POST("/login", LoginFromPost)
	engine.GET("/signup", SignupFromGet)
	authGroup := engine.Group("/")
	authGroup.Use(sessionCheck()) //ログイン状態で閲覧するページをこのグループに追加する
	{
		authGroup.GET("/loggedin", showLoggedIn) // テストページ
	}
	engine.POST("/logout", LogoutFromPost)
}
