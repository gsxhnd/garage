package database

type Pagination struct {
	Limit  uint `validate:"max=100,min=1,number"`
	Offset uint `validate:"min=0,number"`
}
