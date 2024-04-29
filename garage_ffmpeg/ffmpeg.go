package garage_ffmpeg

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/gsxhnd/garage/utils"
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
	GetVideosList() ([]string, error)          // 获取视频列表s
	GetFontsList() ([]string, error)           // 获取字体列表
	GetFontsParams() ([]string, error)         // 获取字体列表
	GetConvertBatch() ([][]string, error)      // 获取转换视频命令
	GetAddFontsBatch() ([][]string, error)     // 获取添加字体命令
	GetAddSubtittleBatch() ([][]string, error) // 获取添加字幕命令
	ExecuteBatch(wOut, wError io.Writer, batchCmd [][]string) error
	// GetExecBatch() rxgo.Observable
}

type videoBatch struct {
	option    *VideoBatchOption
	cmdBatch  []string
	cmdBatchs [][]string
}

var FONT_EXT = []string{".ttf", ".otf", ".ttc"}

func NewVideoBatch(opt *VideoBatchOption) (VideoBatcher, error) {
	if err := utils.MakeDir(opt.OutputPath); err != nil {
		return nil, err
	}

	return &videoBatch{
		option:    opt,
		cmdBatch:  make([]string, 0),
		cmdBatchs: make([][]string, 0),
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

func (vb *videoBatch) GetFontsParams() ([]string, error) {
	var fontsCmdList = []string{}
	fontsList, err := vb.GetFontsList()
	if err != nil {
		return nil, err
	}

	for i, v := range fontsList {
		fontsCmdList = append(fontsCmdList,
			"-attach",
			filepath.Join(vb.option.FontsPath, v),
			fmt.Sprintf("-metadata:s:t:%v", i),
			"mimetype=application/x-truetype-font",
		)
	}

	return fontsCmdList, nil
}

func (vb *videoBatch) GetConvertBatch() ([][]string, error) {
	videosList, err := vb.GetVideosList()
	if err != nil {
		return nil, err
	}

	outputVideosMap := vb.filterOutput(videosList)

	for _, v := range videosList {
		cmd := []string{"-i"}
		cmd = append(cmd, v)
		if vb.option.Advance != "" {
			cmd = append(cmd, vb.option.Advance)
		}
		cmd = append(cmd, outputVideosMap[v])
		vb.cmdBatchs = append(vb.cmdBatchs, cmd)
	}

	return vb.cmdBatchs, nil
}

func (vb *videoBatch) GetAddFontsBatch() ([][]string, error) {
	videosList, err := vb.GetVideosList()
	if err != nil {
		return nil, err
	}

	fontCmd, err := vb.GetFontsParams()
	if err != nil {
		return nil, err
	}

	outputVideosMap := vb.filterOutput(videosList)
	for _, v := range videosList {
		var batchCmd = []string{
			"-i", v,
			"-c", "copy",
		}
		batchCmd = append(batchCmd, fontCmd...)
		batchCmd = append(batchCmd, outputVideosMap[v])
		vb.cmdBatchs = append(vb.cmdBatchs, batchCmd)
	}

	return vb.cmdBatchs, nil
}

func (vb *videoBatch) GetAddSubtittleBatch() ([][]string, error) {
	videosList, err := vb.GetVideosList()
	if err != nil {
		return nil, err
	}

	fontsParams, err := vb.GetFontsParams()
	if err != nil {
		return nil, err
	}

	outputVideosMap := vb.filterOutput(videosList)

	for _, v := range videosList {
		var cmd = []string{}
		// TODO: sub title error
		sourceSubtitle := filepath.Join(vb.option.InputPath, v+vb.option.InputSubSuffix)
		cmd = append(cmd, "-i", fmt.Sprintf(`"%v"`, v))
		cmd = append(cmd, "-sub_charenc", "UTF-8")
		cmd = append(cmd, "-i", fmt.Sprintf(`"%v"`, sourceSubtitle), "-map", "0", "-map", "1")
		cmd = append(cmd, fmt.Sprintf("-metadata:s:s:%v", vb.option.InputSubNo))
		cmd = append(cmd, fmt.Sprintf("language=%v", vb.option.InputSubLang))
		cmd = append(cmd, fmt.Sprintf("-metadata:s:s:%v", vb.option.InputSubNo))
		cmd = append(cmd, fmt.Sprintf(`title="%v"`, vb.option.InputSubTitle))
		cmd = append(cmd, "-c", "copy")
		cmd = append(cmd, fontsParams...)
		cmd = append(cmd, fmt.Sprintf(`"%v"`, outputVideosMap[v]))
		vb.cmdBatchs = append(vb.cmdBatchs, cmd)
	}

	return vb.cmdBatchs, nil
}

func (vb *videoBatch) ExecuteBatch(wOut, wError io.Writer, cmdBatch [][]string) error {
	if !vb.option.Exec {
		return nil
	}

	for _, c := range cmdBatch {
		var cmd *exec.Cmd
		switch runtime.GOOS {
		case "linux":
			cmd = exec.Command("ffmpeg", c...)
		}

		cmd.Stdout = wOut
		cmd.Stderr = wError
		err := cmd.Run()
		if err != nil {
			return err
		}
	}
	return nil
}

func (vb *videoBatch) filterOutput(input []string) map[string]string {
	var output map[string]string = make(map[string]string, 0)
	var existOutput = arraylist.New()
	for _, v := range input {
		filename, _ := strings.CutSuffix(filepath.Base(v), filepath.Ext(v))
		for existOutput.Contains(filename) {
			filename += "-1"
		}
		existOutput.Add(filename)
		output[v] = filepath.Join(vb.option.OutputPath, filename+"."+vb.option.OutputFormat)
	}
	return output
}
