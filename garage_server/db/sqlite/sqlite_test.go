package sqlite

import (
	"github.com/gsxhnd/garage/garage_server/db/database"
	"github.com/gsxhnd/garage/utils"
)

func getMockDB() (database.Driver, error) {
	var logger = utils.NewLogger(&utils.Config{
		Mode: "dev",
		Log: utils.LogConfig{
			Level: "debug",
		},
	})

	var mockSqliteDB, err = NewSqliteDB("../../../data/garage.db", logger)
	if err != nil {
		return nil, err
	}

	return mockSqliteDB, nil
}
