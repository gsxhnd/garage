package database

type Pagination struct {
	Limit  uint64 `validate:"max=100,min=1,number"`
	Offset uint64 `validate:"min=0,number"`
}
