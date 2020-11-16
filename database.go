package main

import (
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func initializeDB() {
	db := DbConnect()
	defer db.Close()
	db.AutoMigrate(&Movie{})
	db.AutoMigrate(&User{})
}

func DbConnect() *gorm.DB {
	db, err := gorm.Open("sqlite3", "todays.db")
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func GetCommonClipedMovie(db *gorm.DB, participants []string) []Movie {
	var movies []Movie

	query := "cliped_user_id IN (" + strings.Join(participants, ", ") + ")"
	tmp := db.Model(&Movie{}).Where(query)
	tmp.Group("movie_id").Find(&movies)
	// for rows.Next() {
	// 	var title string
	// 	var rank int
	// 	rows.Scan(&title, &rank)
	// 	fmt.Println(title, rank)
	// }

	return movies
}

func InsertMovie(db *gorm.DB, m *Movie) {
	//ex: Movie{MovieID: 1, ClipedUserID: 1, Title: "TENET", Rate: 4.1}
	db.Model(&Movie{}).Create(m)
}

func RegistUser(user *User) {
	db := DbConnect()
	defer db.Close()
	db.Model(&User{}).Create(user)
}

func GetUser(username string) User {
	db := DbConnect()
	var user User
	db.First(&user, "filmarks_id=?", username)
	return user
}
