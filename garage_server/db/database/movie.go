package database

type MovieFilter struct {
	Pagination
	Code  string
	Title string
}
