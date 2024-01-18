package ffmpeg

import (
	"testing"
)

func Test_videoBatch_getVideosList(t *testing.T) {
	type args struct {
		inputPath string
		inputType string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{"test_mkv", args{inputPath: "../../testdata", inputType: ".mkv"}, []string{"1", "2", "3", "4", "5"}, false},
		{"test_mp4", args{inputPath: "../../testdata", inputType: ".mp4"}, []string{"1", "2", "3", "4", "5"}, false},
		// {"test_err", args{sourceRootPath: "../../testdata", sourceVideoType: ".mp4"}, []string{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {})
	}
}
