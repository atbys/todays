package main

func initializeRoutes() {
	engine.GET("/", showIndex)
	engine.GET("/login", LoginFromGet)
	engine.GET("/pickup_movie/:id", showPickUpMovie)
	engine.GET("/random", randMovie)
	engine.GET("/from_clip", randMovieFromClip)
	engine.GET("rand_clip", randClipMovie)
	engine.POST("/login", LoginFromPost)
	engine.GET("/signup", SignupFromGet)
	engine.POST("/signup", SignupFromPost)
	authGroup := engine.Group("/")
	authGroup.Use(sessionCheck()) //ログイン状態で閲覧するページをこのグループに追加する
	{
		authGroup.GET("/loggedin", showLoggedIn) // テストページ
	}
	engine.POST("/logout", LogoutFromPost)
}
