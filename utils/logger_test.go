package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewLoggerNilConfig(t *testing.T) {
	require.NotPanics(t, func() {
		logger := NewLogger(nil)
		require.NotNil(t, logger)
	})
}
