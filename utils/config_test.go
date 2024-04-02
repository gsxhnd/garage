package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	tests := []struct {
		name     string
		filaPath string
		wantErr  bool
	}{
		{"", "../config.yaml", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewConfig(tt.filaPath)
			assert.Nil(t, err)
			t.Log("get config:", got)
		})
	}
}
