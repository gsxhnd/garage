package batch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVideoBatch_GetVideos(t *testing.T) {
	tests := []struct {
		name            string
		sourceRootPath  string
		sourceVideoType string
		want            []string
		wantErr         bool
	}{
		{"test_mkv", "../testdata", ".mkv", []string{"1"}, false},
		{"test_mp4", "../testdata", ".mp4", []string{"1"}, false},
		{"test_err", "./test", ".mp4", []string{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vb := new(VideoBatch)
			vb.SourceRootPath = tt.sourceRootPath
			vb.SourceVideoType = tt.sourceVideoType

			got, err := vb.GetVideos()
			assert.NotEqual(t, tt.want, got)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
