package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeDir(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		wantErr bool
	}{
		{"", "../javbus/cover", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := MakeDir(tt.path)
			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
