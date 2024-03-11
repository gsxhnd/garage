package garage_task

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTask(t *testing.T) {
	tests := []struct {
		name string
		want Task
	}{
		{"", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			task := NewTask()
			assert.NotNil(t, task)
			assert.NotEmpty(t, task)
		})
	}
}
