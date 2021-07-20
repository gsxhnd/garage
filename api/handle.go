package api

import (
	"garage/dao"
	"github.com/gorilla/mux"
	"net/http"
)

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
	vars := mux.Vars(req)
	movieCode := vars["code"]
	res, err := dao.GetJavMovieInfo(movieCode)
	if err != nil {
		SendRes(w, err, nil)
		return
	}
	SendRes(w, nil, res)
}

func CrawlJavMovie(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	movieCode := vars["code"]
	SendRes(w, nil, movieCode)
}

func CrawlJavStar(w http.ResponseWriter, req *http.Request) {
	SendRes(w, nil, "movie")
}
