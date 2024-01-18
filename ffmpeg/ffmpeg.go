package ffmpeg

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"go.uber.org/zap"
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
	createDestDir() error                  // 创建输出后的文件夹
	getVideosList() error                  // 获取视频列表
	getFontsList() error                   // 获取字体列表
	getFontsParams(fontsPath string) error // 获取字体
	StartAddSubtittleBatch() error         // 添加字幕
	StartAddFontsBatch() error             // 添加字体
	StartConvertBatch() error              // 转换视频
}

type videoBatch struct {
	option      VideoBatchOption
	videosList  []string
	fontsList   []string
	fontsParams string
	cmdBatch    []string
	logger      *zap.Logger
}

func NewVideoBatch(l *zap.Logger, opt VideoBatchOption) (VideoBatcher, error) {
	client := &videoBatch{
		logger:      l,
		option:      opt,
		videosList:  make([]string, 0),
		fontsList:   make([]string, 0),
		fontsParams: "",
		cmdBatch:    make([]string, 0),
	}

	if err := client.createDestDir(); err != nil {
		return nil, err
	}
	return client, nil
}

func (vb *videoBatch) StartConvertBatch() error {
	template := `ffmpeg.exe -i "%v" %v "%v"`
	if err := vb.getVideosList(); err != nil {
		return err
	}

	for _, v := range vb.videosList {
		inputVideo := filepath.Join(vb.option.InputPath, v+vb.option.InputFormat)
		outputVideo := filepath.Join(vb.option.OutputPath, v+vb.option.OutputFormat)
		s := fmt.Sprintf(template, inputVideo, vb.option.Advance, outputVideo)
		vb.cmdBatch = append(vb.cmdBatch, s)
	}

	if vb.option.Exec {
		return nil
	}
	return nil
}

func (vb *videoBatch) StartAddSubtittleBatch() error {
	if err := vb.getVideosList(); err != nil {
		return err
	}

	vb.logger.Info("Source videos directory: " + vb.option.InputPath)
	vb.logger.Info("Get matching video count: " + strconv.Itoa(len(vb.videosList)))
	vb.logger.Info("Target video's subtitle stream number: " + strconv.Itoa(vb.option.InputSubNo))
	vb.logger.Info("Target video's subtitle language: " + vb.option.InputSubLang)
	vb.logger.Info("Target video's subtitle title: " + vb.option.InputSubTitle)

	if vb.option.FontsPath != "" {
		if err := vb.getFontsParams(vb.option.FontsPath); err != nil {
			return err
		}
		vb.logger.Info("Target video's font paths: " + vb.option.FontsPath)
		vb.logger.Info(fmt.Sprintf("Attach fonts parameters: %v", vb.fontsParams))
	} else {
		vb.logger.Info("Target video's font paths not set, skip.")
	}
	vb.logger.Info("Dest video directory: " + vb.option.OutputPath)

	template := `ffmpeg.exe -i "%s" -sub_charenc UTF-8 -i "%s" -map 0 -map 1 -metadata:s:s:%v language=%v -metadata:s:s:%v title="%v" -c copy %s "%v"`
	for _, v := range vb.videosList {
		sourceVideo := filepath.Join(vb.option.InputPath, v+vb.option.InputFormat)
		sourceSubtitle := filepath.Join(vb.option.InputPath, v+vb.option.InputSubSuffix)
		destVideo := filepath.Join(vb.option.OutputPath, v+vb.option.InputFormat)
		s := fmt.Sprintf(template,
			sourceVideo, sourceSubtitle, vb.option.InputSubNo,
			vb.option.InputSubLang, vb.option.InputSubNo, vb.option.InputSubTitle,
			vb.fontsParams, destVideo)
		vb.cmdBatch = append(vb.cmdBatch, s)
	}

	if vb.option.Exec {
		return nil
	}
	return nil
}

func (vb *videoBatch) StartAddFontsBatch() error {
	if err := vb.getVideosList(); err != nil {
		return err
	}

	if err := vb.getFontsParams(vb.option.FontsPath); err != nil {
		return err
	}

	vb.logger.Info("Source videos directory: " + vb.option.InputPath)
	vb.logger.Info("Get matching video count: " + strconv.Itoa(len(vb.videosList)))
	vb.logger.Info("Target video's font paths: " + vb.option.FontsPath)
	vb.logger.Info(fmt.Sprintf("Attach fonts parameters: %v", vb.fontsParams))
	vb.logger.Info("Dest video directory: " + vb.option.OutputPath)
	template := `ffmpeg.exe -i "%s" -c copy %s "%v"`
	for _, v := range vb.videosList {
		sourceVideo := filepath.Join(vb.option.InputPath, v+vb.option.InputFormat)
		destVideo := filepath.Join(vb.option.OutputPath, v+vb.option.InputFormat)
		s := fmt.Sprintf(template, sourceVideo, vb.fontsParams, destVideo)
		vb.cmdBatch = append(vb.cmdBatch, s)
	}

	if vb.option.Exec {
		return nil
	}
	return nil
}

func (vb *videoBatch) createDestDir() error {
	destDir := path.Join(vb.option.OutputPath)
	vb.logger.Info("Start creating destination directory: " + destDir)
	if fi, err := os.Stat(destDir); err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(destDir, os.ModePerm)
		} else {
			return err
		}
	} else {
		if fi.IsDir() {
			vb.logger.Info("Destination directory already exists")
		}
	}
	vb.logger.Info("Destination directory created")
	return nil
}

func (vb *videoBatch) getVideosList() error {
	err := filepath.Walk(vb.option.InputPath, func(path string, fi os.FileInfo, err error) error {
		if fi == nil {
			if err != nil {
				return err
			}
			return nil
		}
		if fi.IsDir() {
			return nil
		}
		filename := fi.Name()
		fileExt := filepath.Ext(filename)
		if fileExt == vb.option.InputFormat {
			fileName := strings.TrimSuffix(filename, fileExt)
			vb.videosList = append(vb.videosList, fileName)
			return nil
		}
		return nil
	})
	return err
}

func (vb *videoBatch) getFontsList() error {
	fontExts := []string{".ttf", ".otf", ".ttc"}

	err := filepath.Walk(vb.option.FontsPath, func(path string, fi os.FileInfo, err error) error {
		if fi == nil {
			if err != nil {
				return err
			}
			return nil
		}
		if fi.IsDir() {
			return nil
		}
		filename := fi.Name()
		fileExt := filepath.Ext(filename)

		for _, b := range fontExts {
			if fileExt == b {
				vb.fontsList = append(vb.fontsList, filename)
			}
		}
		return nil
	})
	return err
}

func (vb *videoBatch) getFontsParams(fontsPath string) error {
	if fontsPath == "" {
		return nil
	}

	var fontsList = make([]string, 0)
	fontExts := []string{".ttf", ".otf", ".ttc"}
	fontParamsTemplate := `-attach "%s" -metadata:s:t:%v mimetype=application/x-truetype-font`

	if err := filepath.Walk(fontsPath, func(path string, fi os.FileInfo, err error) error {
		if fi == nil {
			if err != nil {
				return err
			}
			return nil
		}
		if fi.IsDir() {
			return nil
		}
		filename := fi.Name()
		fileExt := filepath.Ext(filename)

		for _, b := range fontExts {
			if fileExt == b {
				fontsList = append(fontsList, filename)
			}
		}
		return nil
	}); err != nil {
		return err
	}
	for i, v := range fontsList {
		vb.fontsParams += fmt.Sprintf(fontParamsTemplate, filepath.Join(fontsPath, v), i) + " "
	}
	return nil
}
