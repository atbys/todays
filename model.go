package main

type Movie struct {
	Id       string
	Titile   string
	Rate     string
	Abstruct string
	Reviews  []string
}

type User struct {
	Name       string
	NumOfClips int
	Clips      []string
	NumOfMarks int
	Marks      []string
}
