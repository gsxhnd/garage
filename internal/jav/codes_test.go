package jav

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExpandPrefix(t *testing.T) {
	tests := []struct {
		name           string
		prefix         string
		min, max, zero uint64
		want           []string
	}{
		{
			name:   "zero padded inclusive range",
			prefix: "EKDV",
			min:    1, max: 3, zero: 3,
			want: []string{"EKDV001", "EKDV002", "EKDV003"},
		},
		{
			name:   "no padding when zero is 0",
			prefix: "SSIS-",
			min:    1, max: 2, zero: 0,
			want: []string{"SSIS-1", "SSIS-2"},
		},
		{
			name:   "empty prefix",
			prefix: "  ",
			min:    1, max: 2, zero: 3,
			want: nil,
		},
		{
			name:   "max less than min",
			prefix: "EKDV",
			min:    5, max: 1, zero: 3,
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, ExpandPrefix(tt.prefix, tt.min, tt.max, tt.zero))
		})
	}
}

func TestCodesFromVideosDir(t *testing.T) {
	dir := t.TempDir()
	require.NoError(t, os.MkdirAll(filepath.Join(dir, "nested"), 0o755))
	require.NoError(t, os.WriteFile(filepath.Join(dir, "SSIS-001.mp4"), []byte("x"), 0o644))
	require.NoError(t, os.WriteFile(filepath.Join(dir, "nested", "SSIS-002.mkv"), []byte("x"), 0o644))
	require.NoError(t, os.WriteFile(filepath.Join(dir, "readme.txt"), []byte("x"), 0o644))

	codes, err := CodesFromVideosDir(dir)
	require.NoError(t, err)
	assert.ElementsMatch(t, []string{"SSIS-001", "SSIS-002"}, codes)
}

func TestQueryResolvedCodes(t *testing.T) {
	dir := t.TempDir()
	require.NoError(t, os.WriteFile(filepath.Join(dir, "AAA-001.avi"), []byte("x"), 0o644))

	q := Query{
		Codes:      []string{"BBB-002"},
		Prefixes:   []string{"EKDV"},
		PrefixMin:  1,
		PrefixMax:  2,
		PrefixZero: 3,
		VideosDir:  dir,
	}
	codes, err := q.resolvedCodes()
	require.NoError(t, err)
	assert.Equal(t, []string{"BBB-002", "EKDV001", "EKDV002", "AAA-001"}, codes)
}

func TestQueryPrefixMaxLessThanMin(t *testing.T) {
	_, err := (Query{Prefixes: []string{"EKDV"}, PrefixMin: 5, PrefixMax: 1}).resolvedCodes()
	require.EqualError(t, err, "jav: prefix_max < prefix_min")
}
