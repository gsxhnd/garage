package sqlite

import (
	"testing"

	"github.com/gsxhnd/garage/garage_server/model"
	"github.com/stretchr/testify/assert"
)

func Test_sqliteDB_CreateMovies(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, err := getMockDB()
			assert.Nil(t, err)
			var list = make([]model.Movie, 0)
			var data = model.Movie{
				Code:        "123",
				PublishDate: nil,
			}
			list = append(list, data)

			db.CreateMovies(list)
		})
	}
}

func Test_sqliteDB_GetMovies(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, err := getMockDB()
			assert.Nil(t, err)

			db.GetMovies()
		})
	}
}
