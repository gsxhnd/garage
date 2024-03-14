package service

import (
	"github.com/gsxhnd/garage/garage_server/dao"
	"github.com/gsxhnd/garage/utils"
)

type TestService interface{}

type testService struct {
	logger utils.Logger
	td     dao.TestDao
}

func NewTestService(l utils.Logger, td dao.TestDao) TestService {
	return &testService{
		logger: l,
		td:     td,
	}
}
