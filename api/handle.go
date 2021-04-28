package api

import "net/http"

func GetJavMovie(w http.ResponseWriter, req *http.Request) {
	SendRes(w, nil, "movie")
}

func UpdateJavMovie(w http.ResponseWriter, req *http.Request) {

}

func GetJavStar(w http.ResponseWriter, req *http.Request) {

}

func DeleteJavMovie(w http.ResponseWriter, req *http.Request) {

}

func GetJavMovieInfo(w http.ResponseWriter, req *http.Request) {

}

func CrawlJavMovie(w http.ResponseWriter, req *http.Request) {
	SendRes(w, nil, "movie")
}

func CrawlJavStar(w http.ResponseWriter, req *http.Request) {
	SendRes(w, nil, "movie")
}
