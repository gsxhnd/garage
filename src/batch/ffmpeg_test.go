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

func Test_videoBatch_GetVideosList(t *testing.T) {
	var s = &videoBatch{}
	type args struct {
		sourceRootPath  string
		sourceVideoType string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{"test_mkv", args{sourceRootPath: "../../testdata", sourceVideoType: ".mkv"}, []string{"1", "2", "3", "4", "5"}, false},
		{"test_mp4", args{sourceRootPath: "../../testdata", sourceVideoType: ".mp4"}, []string{"1", "2", "3", "4", "5"}, false},
		// {"test_err", args{sourceRootPath: "../../testdata", sourceVideoType: ".mp4"}, []string{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.GetVideosList(tt.args.sourceRootPath, tt.args.sourceVideoType)
			if tt.wantErr {
				assert.Error(t, err)
				assert.NotEqual(t, tt.want, got)
			} else {
				assert.Equal(t, tt.want, got)
				assert.NoError(t, err)
			}
		})
	}
}
