package jav

// Movie is site-agnostic metadata for one title.
type Movie struct {
	Code           string
	Title          string
	Cover          string
	PublishDate    string
	Length         string
	Director       string
	ProduceCompany string
	PublishCompany string
	Series         string
	Stars          string
	PageURL        string
	Magnets        []Magnet
}

// Magnet is one torrent magnet listing for a movie.
type Magnet struct {
	Name     string
	Link     string
	Size     float64
	Subtitle bool
	HD       bool
}
