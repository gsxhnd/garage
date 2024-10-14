package model

type MovieStar struct {
	Id      uint `json:"id"`
	MovieId uint `json:"movie_id"`
	StarId  uint `json:"star_id"`
}
