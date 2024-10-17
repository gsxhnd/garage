package model

type MovieActor struct {
	Id        uint    `json:"id"`
	MovieId   uint    `json:"movie_id"`
	ActorId   uint    `json:"actor_id"`
	ActorName *string `json:"actor_name"`
}
