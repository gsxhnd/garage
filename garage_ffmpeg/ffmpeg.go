package garage_ffmpeg

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/gsxhnd/garage/utils"
	"github.com/reactivex/rxgo/v2"
)

type VideoBatchOption struct {
	InputPath      string
	InputFormat    string
	OutputPath     string
	OutputFormat   string
	FontsPath      string
	InputSubSuffix string
	InputSubNo     int
	InputSubTitle  string
	InputSubLang   string
	Advance        string
	Exec           bool
}

type VideoBatcher interface {
	GetVideosList() ([]string, error)        // 获取视频列表s
	GetFontsList() ([]string, error)         // 获取字体列表
	GetFontsParams() (string, error)         // 获取字体列表
	GetConvertBatch() ([]string, error)      // 获取转换视频命令
	GetAddFontsBatch() ([]string, error)     // 获取添加字体命令
	GetAddSubtittleBatch() ([]string, error) // 获取添加字幕命令
	ExecuteBatch(wOut, wError io.Writer, batchCmd []string) error
	GetExecBatch() rxgo.Observable
}

type videoBatch struct {
	option   *VideoBatchOption
	cmdBatch []string
	Ob       *Observable
}

var FONT_EXT = []string{".ttf", ".otf", ".ttc"}

const CONVERT_TEMPLATE = `ffmpeg.exe -i "%v" %v "%v"`
const ADD_SUB_TEMPLATE = `ffmpeg.exe -i "%s" -sub_charenc UTF-8 -i "%s" -map 0 -map 1 -metadata:s:s:%v language=%v -metadata:s:s:%v title="%v" -c copy %s "%v"`
const ADD_FONT_TEMPLATE = `ffmpeg.exe -i "%s" -c copy %s "%v"`
const FONT_TEMPLATE = `-attach "%s" -metadata:s:t:%v mimetype=application/x-truetype-font `

func NewVideoBatch(opt *VideoBatchOption) (VideoBatcher, error) {
	if err := utils.MakeDir(opt.OutputPath); err != nil {
		return nil, err
	}

	return &videoBatch{
		option:   opt,
		cmdBatch: make([]string, 0),
		Ob:       ObWriterNew(),
	}, nil
}

func (vb *videoBatch) GetVideosList() ([]string, error) {
	var videosList []string = make([]string, 0)
	if err := filepath.Walk(vb.option.InputPath, func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if fi.IsDir() {
			return nil
		}

		filename := fi.Name()
		fileExt := filepath.Ext(filename)

		if fileExt == "."+vb.option.InputFormat {
			videosList = append(videosList, path)
			return nil
		}
		return nil
	}); err != nil {
		return nil, err
	} else {
		return videosList, nil
	}
}

func (vb *videoBatch) GetFontsList() ([]string, error) {
	var fontsList []string = make([]string, 0)
	if err := filepath.Walk(vb.option.FontsPath, func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if fi.IsDir() {
			return nil
		}

		filename := fi.Name()
		fileExt := filepath.Ext(filename)

		for _, b := range FONT_EXT {
			if fileExt == b {
				fontsList = append(fontsList, path)
			}
		}
		return nil
	}); err != nil {
		return nil, err
	} else {
		return fontsList, nil
	}
}

func (vb *videoBatch) GetFontsParams() (string, error) {
	var fontsParams = ""
	fontsList, err := vb.GetFontsList()
	if err != nil {
		return "", err
	}

	for i, v := range fontsList {
		fontPath := filepath.Join(vb.option.FontsPath, v)
		fontsParams += fmt.Sprintf(FONT_TEMPLATE, fontPath, i)
	}

	return fontsParams, nil
}

func (vb *videoBatch) GetConvertBatch() ([]string, error) {
	videosList, err := vb.GetVideosList()
	if err != nil {
		return nil, err
	}

	for _, v := range videosList {
		filename, _ := strings.CutSuffix(filepath.Base(v), filepath.Ext(v))
		inputVideo := filepath.Join(vb.option.InputPath, v)
		outputVideo := filepath.Join(vb.option.OutputPath, filename+"."+vb.option.OutputFormat)
		s := fmt.Sprintf(CONVERT_TEMPLATE, inputVideo, vb.option.Advance, outputVideo)
		vb.cmdBatch = append(vb.cmdBatch, s)
	}

	return vb.cmdBatch, nil
}

func (vb *videoBatch) GetAddFontsBatch() ([]string, error) {
	videosList, err := vb.GetVideosList()
	if err != nil {
		return nil, err
	}

	fontsParams, err := vb.GetFontsParams()
	if err != nil {
		return nil, err
	}

	for _, v := range videosList {
		outputVideo := filepath.Join(vb.option.OutputPath, filepath.Base(v))
		s := fmt.Sprintf(ADD_FONT_TEMPLATE, v, fontsParams, outputVideo)
		vb.cmdBatch = append(vb.cmdBatch, s)
	}

	return vb.cmdBatch, nil
}

func (vb *videoBatch) GetAddSubtittleBatch() ([]string, error) {
	videosList, err := vb.GetVideosList()
	if err != nil {
		return nil, err
	}

	fontsParams, err := vb.GetFontsParams()
	if err != nil {
		return nil, err
	}

	for _, v := range videosList {
		sourceVideo := filepath.Join(vb.option.InputPath, v+vb.option.InputFormat)
		sourceSubtitle := filepath.Join(vb.option.InputPath, v+vb.option.InputSubSuffix)
		destVideo := filepath.Join(vb.option.OutputPath, v+vb.option.InputFormat)
		s := fmt.Sprintf(ADD_SUB_TEMPLATE,
			sourceVideo, sourceSubtitle, vb.option.InputSubNo,
			vb.option.InputSubLang, vb.option.InputSubNo, vb.option.InputSubTitle,
			fontsParams, destVideo)
		vb.cmdBatch = append(vb.cmdBatch, s)
	}

	return vb.cmdBatch, nil
}

func (vb *videoBatch) ExecuteBatch(wOut, wError io.Writer, cmdBatch []string) error {
	var name string
	switch runtime.GOOS {
	case "darwin":
		name = "/bin/sh"
	case "windows":
		name = "powershell"
	case "linux":
		name = "/bin/bash"
	default:
		name = ""
	}

	for _, cmd := range cmdBatch {
		if !vb.option.Exec {
			return nil
		}

		cmd := exec.Command(name, cmd)
		cmd.Stdout = wOut
		cmd.Stderr = wError
		err := cmd.Run()
		if err != nil {
			return err
		}
	}
	return nil

}

func (vb *videoBatch) GetExecBatch() rxgo.Observable {
	return vb.Ob.ob
}

type Observable struct {
	ob rxgo.Observable
	ch chan rxgo.Item
}

func ObWriterNew() *Observable {
	ch := make(chan rxgo.Item)
	return &Observable{
		ob: rxgo.FromChannel(ch),
		ch: ch,
	}
}

func (o *Observable) Write(p []byte) (int, error) {
	o.ch <- rxgo.Of(p)
	return len(p), nil
}
