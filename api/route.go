package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

func routes(imgDir string) *mux.Router {
	r := mux.NewRouter()
	r.PathPrefix("/img/").Handler(http.StripPrefix("/img/", http.FileServer(http.Dir(imgDir))))

	api := r.PathPrefix("/v1/api").Subrouter()
	api.HandleFunc("/jav/movie", GetJavMovie).Methods("GET") // get jav movie list
	api.HandleFunc("/jav/movie", UpdateJavMovie).Methods("PUT")
	api.HandleFunc("/jav/star", GetJavStar).Methods("GET") // get jav star list

	api.HandleFunc("/crawl/jav/movie", GetJavMovie).Methods("POST")

	return r
}
