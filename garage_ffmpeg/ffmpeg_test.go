package garage_ffmpeg

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createCorrectVideos(inputPath, formart string) []string {
	var a = make([]string, 0)
	for i := 1; i < 5; i++ {
		for j := 1; j < 6; j++ {
			a = append(a, inputPath+"/"+strconv.Itoa(i)+"/"+strconv.Itoa(j)+"."+formart)
		}
	}
	return a
}

var correctMp4Videos = []string{
	"../testdata/1/1.mp4",
	"../testdata/1/2.mp4",
	"../testdata/2/1.mp4",
	"../testdata/2/2.mp4",
}
var correctMkvVideos = []string{
	"../testdata/1/1.mkv",
	"../testdata/1/2.mkv",
	"../testdata/2/1.mkv",
	"../testdata/2/2.mkv",
}

var correctConvertBatch = []string{
	`ffmpeg.exe -i "../testdata/1/1.mp4"  "../output/1.mkv"`,
	`ffmpeg.exe -i "../testdata/1/2.mp4"  "../output/2.mkv"`,
	`ffmpeg.exe -i "../testdata/2/1.mp4"  "../output/1.mkv"`,
	`ffmpeg.exe -i "../testdata/2/2.mp4"  "../output/2.mkv"`,
}

func Test_videoBatch_getVideosList(t *testing.T) {
	type args struct {
		InputPath   string
		InputFormat string
	}

	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{"test_mkv", args{InputPath: "../testdata", InputFormat: "mkv"}, correctMkvVideos, false},
		{"test_mp4", args{InputPath: "../testdata", InputFormat: "mp4"}, correctMp4Videos, false},
		{"test_err", args{InputPath: "../111", InputFormat: "mp4"}, []string{}, true},
		{"test_err", args{InputPath: "../111", InputFormat: "mp4"}, []string{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vb := &videoBatch{
				option: &VideoBatchOption{
					InputPath:   tt.args.InputPath,
					InputFormat: tt.args.InputFormat,
				},
			}

			videosList, err := vb.GetVideosList()
			t.Log(videosList)

			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, videosList, tt.want)
			}
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
			}

			got, err := vb.GetFontsParams()
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

func Test_videoBatch_GetExecBatch(t *testing.T) {
	type fields struct {
		option   *VideoBatchOption
		cmdBatch []string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"", fields{nil, nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vb, err := NewVideoBatch(tt.fields.option)
			assert.Nil(t, err)
			vb.GetExecBatch()
		})
	}
}

func Test_videoBatch_GetConvertBatch(t *testing.T) {
	type args struct {
		InputPath    string
		InputFormat  string
		OutputPath   string
		OutputFormat string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{"", args{InputPath: "../testdata", InputFormat: "mp4", OutputPath: "../output", OutputFormat: "mkv"}, correctConvertBatch, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vb := &videoBatch{
				option: &VideoBatchOption{
					InputPath:    tt.args.InputPath,
					InputFormat:  tt.args.InputFormat,
					OutputPath:   tt.args.OutputPath,
					OutputFormat: tt.args.OutputFormat,
				},
			}

			got, err := vb.GetConvertBatch()
			t.Log(got)

			if tt.wantErr {
				assert.NotNil(t, err)
			}

			assert.Equal(t, correctConvertBatch, got)
		})
	}
}
