package api

import (
	"github.com/gorilla/mux"
	"github.com/gsxhnd/owl"
	"net/http"
)

func routes(imgDir string) *mux.Router {
	r := mux.NewRouter()
	r.Use(AuthMiddleware(owl.GetString("dashboard.key"), owl.GetString("dashboard.secret")).Middleware)
	r.PathPrefix("/img/").Handler(http.StripPrefix("/img/", http.FileServer(http.Dir(imgDir))))
	api := r.PathPrefix("/v1/api").Subrouter()
	api.HandleFunc("/jav/movie", GetJavMovie).Methods("GET")            // get jav movie list
	api.HandleFunc("/jav/movie/{code}", GetJavMovieInfo).Methods("GET") // get jav movie info include star and tag's list
	api.HandleFunc("/jav/movie", UpdateJavMovie).Methods("PUT")         // update jav movie info
	api.HandleFunc("/jav/movie", DeleteJavMovie).Methods("DELETE")      // update jav movie info
	api.HandleFunc("/jav/star", GetJavStar).Methods("GET")              // get jav star list
	api.HandleFunc("/crawl/jav/movie", CrawlJavMovie).Methods("POST")   // crawl jav movie info
	api.HandleFunc("/crawl/jav/star", CrawlJavStar).Methods("POST")     // crawl jav star info
	api.HandleFunc("/version", GetVersion).Methods("GET")               // server version

	return r
}
