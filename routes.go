package main

func initializeRoutes() {
	engine.GET("/", showIndex)
	engine.GET("/movie", showPickUpMovie)
}
