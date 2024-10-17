package model

import (
	"time"
)

type Movie struct {
	Id             uint       `json:"id"`
	Code           string     `json:"code" validate:"required"`
	Title          string     `json:"title" validate:"required"`
	Cover          *string    `json:"cover"`
	PublishDate    *time.Time `json:"publish_date"`
	Director       *string    `json:"director"`
	ProduceCompany *string    `json:"produce_company"`
	PublishCompany *string    `json:"publish_company"`
	Series         *string    `json:"series"`
	CreatedAt      *time.Time `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
}

type MovieInfo struct {
	Movie  Movie        `json:"movie"`
	Actors []MovieActor `json:"actors"`
	Tags   []MovieTag   `json:"tags"`
}
