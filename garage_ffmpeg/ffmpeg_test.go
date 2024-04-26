package garage_ffmpeg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestNewVideoBatch(t *testing.T) {
	tests := []struct {
		name    string
		opt     *VideoBatchOption
		wantErr bool
	}{
		{"test_succ", &VideoBatchOption{
			OutputPath: "../testdata",
		}, false},
		{"test_fail", &VideoBatchOption{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewVideoBatch(tt.opt)
			if tt.wantErr {
				assert.NotNil(t, err)
				return
			}
			assert.Nil(t, err)
			assert.NotNil(t, got)
		})
	}
}

func Test_videoBatch_GetVideosList(t *testing.T) {
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

var correctFonts = []string{
	"../testdata/1/1.ttf",
	"../testdata/1/2.ttf",
	"../testdata/2/1.ttf",
	"../testdata/2/2.ttf",
}

func Test_videoBatch_GetFontsList(t *testing.T) {
	type args struct {
		FontsPath string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{"test_succ", args{FontsPath: "../testdata"}, correctFonts, false},
		{"test_err", args{FontsPath: "../111"}, nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vb := &videoBatch{
				option: &VideoBatchOption{
					FontsPath: tt.args.FontsPath,
				},
			}

			got, err := vb.GetFontsList()
			if tt.wantErr {
				assert.NotNil(t, err)
				return
			}
			assert.Nil(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

var correctFontsCmd = []string{
	"-attach", "../testdata/1/1.ttf", "-metadata:s:t:0", "mimetype=application/x-truetype-font",
	"-attach", "../testdata/1/2.ttf", "-metadata:s:t:1", "mimetype=application/x-truetype-font",
	"-attach", "../testdata/2/1.ttf", "-metadata:s:t:2", "mimetype=application/x-truetype-font",
	"-attach", "../testdata/2/2.ttf", "-metadata:s:t:3", "mimetype=application/x-truetype-font",
}

func Test_videoBatch_GetFontsParams(t *testing.T) {
	type args struct {
		FontsPath string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{"test_succ", args{FontsPath: "../testdata"}, correctFontsCmd, false},
		{"test_err", args{FontsPath: "../111"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vb := &videoBatch{
				option: &VideoBatchOption{
					FontsPath: tt.args.FontsPath,
				},
			}

			got, err := vb.GetFontsParams()
			if tt.wantErr {
				assert.NotNil(t, err)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

var correctConvertBatch = [][]string{
	{"-i", "../testdata/1/1.mp4", "../testdata/output/1.mkv"},
	{"-i", "../testdata/1/2.mp4", "../testdata/output/2.mkv"},
	{"-i", "../testdata/2/1.mp4", "../testdata/output/1-1.mkv"},
	{"-i", "../testdata/2/2.mp4", "../testdata/output/2-1.mkv"},
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
		want    [][]string
		wantErr bool
	}{
		{"test_succ", args{InputPath: "../testdata", InputFormat: "mp4", OutputPath: "../testdata/output", OutputFormat: "mkv"}, correctConvertBatch, false},
		{"test_fail", args{InputPath: "../111", InputFormat: "mp4", OutputPath: "../testdata/output", OutputFormat: "mkv"}, nil, true},
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
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

var correctAddFontsCmdBatch = func() [][]string {
	var a = [][]string{
		{"-i", "../testdata/1/1.mkv", "-c", "copy"},
		{"-i", "../testdata/1/2.mkv", "-c", "copy"},
		{"-i", "../testdata/2/1.mkv", "-c", "copy"},
		{"-i", "../testdata/2/2.mkv", "-c", "copy"},
	}

	var b = []string{
		"../testdata/output/1.mkv",
		"../testdata/output/2.mkv",
		"../testdata/output/1-1.mkv",
		"../testdata/output/2-1.mkv",
	}
	var c = [][]string{}
	for i, v := range a {
		v = append(v, correctFontsCmd...)
		v = append(v, b[i])
		c = append(c, v)
	}
	return c
}

func Test_videoBatch_GetAddFontsBatch(t *testing.T) {
	type args struct {
		InputPath    string
		InputFormat  string
		OutputPath   string
		OutputFormat string
		FontsPath    string
	}
	tests := []struct {
		name    string
		args    args
		want    [][]string
		wantErr bool
	}{
		{"test_succ", args{InputPath: "../testdata", InputFormat: "mkv", FontsPath: "../testdata", OutputPath: "../testdata/output"}, correctAddFontsCmdBatch(), false},
		{"test_fail", args{InputPath: "../111", InputFormat: "mkv", OutputPath: "../testdata/output", OutputFormat: "mkv"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vb := &videoBatch{
				option: &VideoBatchOption{
					InputPath:    tt.args.InputPath,
					InputFormat:  tt.args.InputFormat,
					OutputPath:   tt.args.OutputPath,
					OutputFormat: tt.args.InputFormat,
					FontsPath:    tt.args.FontsPath,
				},
			}

			got, err := vb.GetAddFontsBatch()
			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

var correctMp4FilterOutput = map[string]string{
	"../testdata/1/1.mp4": "../output/1.mkv",
	"../testdata/1/2.mp4": "../output/2.mkv",
	"../testdata/2/1.mp4": "../output/1-1.mkv",
	"../testdata/2/2.mp4": "../output/2-1.mkv",
}

func Test_videoBatch_filterOutput(t *testing.T) {
	type args struct {
		outPath   string
		outFormat string
	}

	tests := []struct {
		name      string
		inputPath []string
		args      args
		want      map[string]string
	}{
		{"test", correctMp4Videos, args{outPath: "../output", outFormat: "mkv"}, correctMp4FilterOutput},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vb := &videoBatch{
				option: &VideoBatchOption{
					OutputPath:   tt.args.outPath,
					OutputFormat: tt.args.outFormat,
				},
			}

			got := vb.filterOutput(tt.inputPath)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_videoBatch_ExecuteBatch(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// var a = []string{"-i", "../testdata/1/1.mp4", "", "../data/1.mkv"}

			// cmd := exec.Command("ffmpeg", a...)
			// cmd.Stdout = os.Stdout
			// cmd.Stderr = os.Stdout
			// cmd.Run()
		})
	}
}
