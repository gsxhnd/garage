package storage

type localStorage struct{}

func NewLocalStorage() Storage {
	return &localStorage{}
}

func (s *localStorage) Ping() error {
	return nil
}
