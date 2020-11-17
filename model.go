package main

type MovieInfo struct {
	Id       string
	Titile   string
	Rate     string
	Abstruct string
	FLink    string
	Reviews  []string
}

//DB用
type User struct {
	Name       string
	Password   string
	FilmarksID string
}

//DB用
type Movie struct {
	MovieID      int
	ClipedUserID int
	Title        string
	Rate         float32
}
