package batch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
			err := s.getVideosList()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
