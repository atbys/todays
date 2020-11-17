package main

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const baseURL = "https://filmarks.com"

func GetMovieTitle(doc *goquery.Document) string {
	movie_titile := doc.Find("div.p-content-detail__main > h2 > span").Text()
	return movie_titile
}

func GetMovieRate(doc *goquery.Document) string {
	movie_rate := doc.Find("div.p-content-detail-state > div > div > div.c-rating__score").Text()

	return movie_rate
}

func GetMovieAbstruct(doc *goquery.Document) string { //TODO
	movie_abstruct, _ := doc.Find("body > div.l-main > div.p-content-detail > div.p-content-detail__head > div > div.p-content-detail__body > div.p-content-detail__main > div.p-content-detail__synopsis").Html()

	return movie_abstruct
}

func GetMovieReviews(doc *goquery.Document) []string {
	movie_reviews_raw := doc.Find("body > div.l-main > div.p-content-detail > div.p-content-detail__foot > div.p-main-area.p-timeline > div.p-mark > div.p-mark__review")
	var movie_reviews []string
	movie_reviews_raw.Each(func(i int, s *goquery.Selection) {
		movie_reviews = append(movie_reviews, s.Text())
	})
	return movie_reviews
}

func GetMovieInfo(id string) *MovieInfo {
	targetURL := baseURL + "/movies/" + id + "/no_spoiler"

	doc, err := goquery.NewDocument(targetURL)
	if err != nil {
		panic("failed to get html")
	}

	movie := &MovieInfo{
		Titile:   GetMovieTitle(doc),
		Rate:     GetMovieRate(doc),
		Abstruct: "TODO", //GetMovieAbstruct(doc)
		FLink:    targetURL,
		Reviews:  GetMovieReviews(doc),
	}

	return movie
}

func GetNumOfClips(doc *goquery.Document) int {
	numOfClips_str := doc.Find("body > div.l-main > div.p-users-navi > div > ul > li.p-users-navi__item.p-users-navi__item--clips.is-active > div > span.p-users-navi__count").Text()
	numOfClips, _ := strconv.Atoi(numOfClips_str)

	return numOfClips
}

func GetIdOfClipMovies(username string) []string {
	page := 1
	targetURL := baseURL + "/users/" + username + "/clips" + "?page=" + strconv.Itoa(page)
	var doc *goquery.Document
	var err error
	doc, err = goquery.NewDocument(targetURL)
	if err != nil {
		panic(err)
	}
	numOfClips := GetNumOfClips(doc)
	var ids []string

	for numOfClips > 0 {
		clipCountInPage := 0
		res := doc.Find("body > div.l-main > div.p-content > div.p-contents-grid > div.c-content-item > a")
		res.Each(func(i int, s *goquery.Selection) {
			l, _ := s.Attr("href")
			tmp := strings.Split(l, "/")
			id := tmp[len(tmp)-1]
			ids = append(ids, id)
			clipCountInPage += 1
		})
		numOfClips -= clipCountInPage
		page += 1
		targetURL := baseURL + "/users/" + username + "/clips" + "?page=" + strconv.Itoa(page)
		doc, err = goquery.NewDocument(targetURL)
		if err != nil {
			panic(err)
		}
	}

	return ids
}

func GetNumOfMarks(doc *goquery.Document) int {
	res := doc.Find("body > div.l-main > div.p-users-navi > div > ul.p-users-navi__list > li.p-users-navi__item.p-users-navi__item--marks.is-active > div > span.p-users-navi__count").Text()
	numOfMarks, _ := strconv.Atoi(res)

	return numOfMarks
}

func GetIdOfMarkMovies(username string) []string {
	page := 1
	targetURL := baseURL + "/users/" + username + "?page=" + strconv.Itoa(page)
	var doc *goquery.Document
	var err error
	doc, err = goquery.NewDocument(targetURL)
	if err != nil {
		panic(err)
	}
	numOfMarks := GetNumOfMarks(doc)
	var ids []string
	for numOfMarks > 0 { //future work: 一回の表示件数で割った数のごルーチンを建てる
		markCountInPage := 0
		res := doc.Find("body > div.l-main > div.p-content > div.p-contents-list > div.c-content-card > div.c-content-card__right > div > a")
		res.Each(func(i int, s *goquery.Selection) {
			l, _ := s.Attr("href")
			tmp := strings.Split(l, "/")                 //
			id := strings.Split(tmp[len(tmp)-1], "?")[0] //リンクのIDの部分だけ取り出す
			ids = append(ids, id)
			markCountInPage += 1
		})
		numOfMarks -= markCountInPage
		page += 1
		targetURL := baseURL + "/users/" + username + "?page=" + strconv.Itoa(page)
		doc, err = goquery.NewDocument(targetURL)
		if err != nil {
			panic(err)
		}
	}

	return ids
}
