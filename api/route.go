package api

import (
	"github.com/gorilla/mux"
)

func routes() *mux.Router {
	r := mux.NewRouter()
	api := r.PathPrefix("/v1/api").Subrouter()
	api.HandleFunc("/jav/movie", GetJavMovie).Methods("GET") // get jav movie list
	api.HandleFunc("/jav/movie", UpdateJavMovie).Methods("PUT")
	api.HandleFunc("/jav/star", GetJavStar).Methods("GET") // get jav star list

	return r
}
