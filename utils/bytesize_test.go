package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func bytesFromGB(n float64) uint64 {
	return uint64(n * float64(1<<30))
}

func TestParseByteSize(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    uint64
		wantErr bool
	}{
		{name: "bytes", input: "1024B", want: 1024},
		{name: "kilobytes short", input: "1KB", want: 1024},
		{name: "megabytes spaced", input: "512 MB", want: 512 * 1024 * 1024},
		{name: "gigabytes decimal", input: "1.5GB", want: bytesFromGB(1.5)},
		{name: "javbus style", input: "4.85GB", want: bytesFromGB(4.85)},
		{name: "long unit", input: "1kilobyte", want: 1024},
		{name: "iec unit", input: "2GiB", want: 2 * 1024 * 1024 * 1024},
		{name: "empty", input: "", wantErr: true},
		{name: "no unit", input: "12", wantErr: true},
		{name: "unknown unit", input: "12XB", wantErr: true},
		{name: "negative", input: "-1MB", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseByteSize(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestByteSizeToMB(t *testing.T) {
	tests := []struct {
		name string
		in   uint64
		want float64
	}{
		{name: "exact megabyte", in: 1024 * 1024, want: 1},
		{name: "from gigabytes", in: bytesFromGB(4.85), want: 4966.4},
		{name: "zero", in: 0, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, ByteSizeToMB(tt.in))
		})
	}
}
