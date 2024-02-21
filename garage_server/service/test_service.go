package service

import "github.com/gsxhnd/garage/utils"

type TestService interface{}

type testService struct {
	logger utils.Logger
	// db     dao.Database
}

func NewTestService(l utils.Logger) TestService {
	return &testService{
		logger: l,
		// db:     db,
	}
}
