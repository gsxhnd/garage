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
		{"test_mkv", args{InputPath: "../testdata", InputFormat: "mkv"}, []string{}, false},
		{"test_mp4", args{InputPath: "../testdata", InputFormat: "mp4"}, []string{}, false},
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
				var cList = createCorrectVideos(tt.args.InputPath, tt.args.InputFormat)
				t.Log(cList)

				assert.Nil(t, err)
				assert.Equal(t, videosList, cList)
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
