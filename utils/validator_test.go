package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewValidator(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			type s struct {
				A string `validate:"required"`
				a string `validate:"required"`
			}
			testS := s{
				A: "11",
			}
			v := NewValidator()

			err1 := v.Struct(&testS)
			err2 := v.Var("", "required")
			assert.NotNil(t, err1)
			println(err1)
			assert.NotNil(t, err2)
		})
	}
}
