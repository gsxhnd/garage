package database

type Driver interface {
	Ping() error
	GetMovie()
	Migrate() error
}
