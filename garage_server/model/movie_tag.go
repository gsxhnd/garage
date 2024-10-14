package model

type MovieTag struct {
	Id      uint `json:"id"`
	MovieId uint `json:"movie_id"`
	TagId   uint `json:"tag_id"`
}
