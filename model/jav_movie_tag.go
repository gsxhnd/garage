package model

type JavMovieTag struct {
	TagId        uint64 `json:"id"`
	JavMovieCode string `json:"jav_movie_code" gorm:"type:varchar(100)"`
}

func (jmt *JavMovieTag) TableName() string {
	return "jav_movie_tag"
}
