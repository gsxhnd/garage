package ffmpeg

import (
	"testing"

	"github.com/gsxhnd/garage/src/utils"
	"github.com/stretchr/testify/assert"
)

var logger = utils.GetLogger()

func Test_videoBatch_getVideosList(t *testing.T) {
	type args struct {
		inputPath   string
		InputFormat string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{"test_mkv", args{inputPath: "../testdata", InputFormat: ".mkv"}, []string{"1", "2", "3", "4", "5"}, false},
		{"test_mp4", args{inputPath: "../testdata", InputFormat: ".mp4"}, []string{"1", "2", "3", "4", "5"}, false},
		// {"test_err", args{sourceRootPath: "../../testdata", sourceVideoType: ".mp4"}, []string{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vb := &videoBatch{
				logger: logger,
				option: &VideoBatchOption{
					InputPath:   tt.args.inputPath,
					InputFormat: tt.args.InputFormat,
				},
			}
			err := vb.getVideosList()
			t.Log(vb.videosList)
			assert.Nil(t, err)
		})
	}
}

func Test_videoBatch_getFontsParams(t *testing.T) {
	type fields struct {
		option *VideoBatchOption
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vb := &videoBatch{
				option: tt.fields.option,

				logger: logger,
			}
			got, err := vb.getFontsParams()
			if (err != nil) != tt.wantErr {
				t.Errorf("videoBatch.getFontsParams() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("videoBatch.getFontsParams() = %v, want %v", got, tt.want)
			}
		})
	}
}
