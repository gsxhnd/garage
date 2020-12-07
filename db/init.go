package db

type Databaser interface {
	Connect() error
	Close()
}
