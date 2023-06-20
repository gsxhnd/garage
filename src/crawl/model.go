package crawl

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
	// Stars          []string `json:"stars"`
	Stars string `json:"stars"`
}

type JavMovieMagnet struct {
	Name     string  `json:"name"`
	Link     string  `json:"link"`
	Size     float64 `json:"size"`
	Subtitle bool    `json:"subtitle"`
	HD       bool    `json:"hd"`
}
