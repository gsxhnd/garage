package model

import "time"

type JavCrawlConfig struct {
	Proxy          string        `json:"proxy"`
	DestPath       string        `json:"dest_path"`
	DownloadMagent bool          `json:"download_magent"`
	DownloadCover  bool          `json:"download_cover"`
	Code           string        `json:"code"`
	StarCode       string        `json:"star_code"`
	PrefixCode     string        `json:"prefix_code"`
	PrefixMinNo    uint          `json:"prefix_min_no"`
	PrefixMaxNo    uint          `json:"prefix_max_no"`
	PageStartNo    uint          `json:"page_start_no"`
	RandomDelay    time.Duration `json:"random_delay"`
	Parallelism    int           `json:"parallelism"`
}

type JavMovie struct {
	Code           string `json:"code"`
	Title          string `json:"title"`
	Cover          string `json:"cover"`
	PublishDate    string `json:"publish_date"`
	Length         string `json:"length"`
	Director       string `json:"director"`
	ProduceCompany string `json:"produce_company"`
	PublishCompany string `json:"publish_company"`
	Series         string `json:"series"`
	Stars          string `json:"stars"`
}

type JavMovieMagnet struct {
	Name     string  `json:"name"`
	Link     string  `json:"link"`
	Size     float64 `json:"size"`
	Subtitle bool    `json:"subtitle"`
	HD       bool    `json:"hd"`
}
