package model

type JavMovieSatr struct {
	JavMovieCode string `json:"jav_movie_code" gorm:"type:varchar(100)"` // 番号
	JavStarId    uint64 `json:"jav_star_id"`                             // 演员ID
	JavStarName  string `json:"jav_star_name" gorm:"type:varchar(100)"`  // 演员名
}

func (jms *JavMovieSatr) TableName() string {
	return "jav_movie_star"
}
