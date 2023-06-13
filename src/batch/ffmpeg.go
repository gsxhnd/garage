package batch

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"go.uber.org/zap"
)

type VideoBatcher interface {
	// 创建输出后的文件夹
	createDestDir() error
	// 获取视频列表
	getVideosList() error
	// 获取字体
	getFontsParams(fontsPath string) error
	// get bactch
	GetAddSubtitleBatch(sourceSubtitleNumber int, sourceSubtitleType, sourceSubtitleLanguage, sourceSubtitleTitle, fontsPath string) ([]string, error)
	GetConvertBatch(advance, destVideoType string) ([]string, error)
}

type videoBatch struct {
	inputPath   string
	inputType   string
	videosList  []string
	fontsParams string
	outputPath  string
	cmdBatch    []string
	logger      *zap.Logger
}

func NewVideoBatch(l *zap.Logger, inputPath, inputType, outputPath string) (VideoBatcher, error) {
	client := &videoBatch{
		logger:      l,
		inputPath:   inputPath,
		inputType:   inputType,
		outputPath:  outputPath,
		videosList:  make([]string, 0),
		fontsParams: "",
		cmdBatch:    make([]string, 0),
	}
	if err := client.createDestDir(); err != nil {
		return nil, err
	}
	return client, nil
}

func (vb *videoBatch) createDestDir() error {
	destDir := path.Join(vb.outputPath)
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

func (vb *videoBatch) GetAddSubtitleBatch(inputSubNo int, inputSubSuffix, inputSubLang, inputSubTitle, fontsPath string) ([]string, error) {
	if err := vb.getVideosList(); err != nil {
		return nil, err
	}

	vb.logger.Info("Source videos directory: " + vb.inputPath)
	vb.logger.Info("Get matching video count: " + strconv.Itoa(len(vb.videosList)))
	vb.logger.Info("Target video's subtitle stream number: " + strconv.Itoa(inputSubNo))
	vb.logger.Info("Target video's subtitle language: " + inputSubLang)
	vb.logger.Info("Target video's subtitle title: " + inputSubTitle)
	if fontsPath != "" {
		if err := vb.getFontsParams(fontsPath); err != nil {
			return nil, err
		}
		vb.logger.Info("Target video's font paths: " + fontsPath)
		vb.logger.Info(fmt.Sprintf("Attach fonts parameters: %v", vb.fontsParams))
	} else {
		vb.logger.Info("Target video's font paths not set, skip.")
	}
	vb.logger.Info("Dest video directory: " + vb.outputPath)

	template := `ffmpeg.exe -i "%s" -sub_charenc UTF-8 -i "%s" -map 0 -map 1 -metadata:s:s:%v language=%v -metadata:s:s:%v title="%v" -c copy %s "%v"`
	for _, v := range vb.videosList {
		sourceVideo := filepath.Join(vb.inputPath, v+vb.inputType)
		sourceSubtitle := filepath.Join(vb.inputPath, v+inputSubSuffix)
		destVideo := filepath.Join(vb.outputPath, v+vb.inputType)
		s := fmt.Sprintf(template,
			sourceVideo, sourceSubtitle, inputSubNo,
			inputSubLang, inputSubNo, inputSubTitle,
			vb.fontsParams, destVideo)
		vb.cmdBatch = append(vb.cmdBatch, s)
	}
	return vb.cmdBatch, nil
}

func (vb *videoBatch) getVideosList() error {
	err := filepath.Walk(vb.inputPath, func(path string, fi os.FileInfo, err error) error {
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
		if fileExt == vb.inputType {
			fileName := strings.TrimSuffix(filename, fileExt)
			vb.videosList = append(vb.videosList, fileName)
			return nil
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

func (vb *videoBatch) GetConvertBatch(advance, outputType string) ([]string, error) {
	template := `ffmpeg.exe -i "%v" %v "%v"`
	if err := vb.getVideosList(); err != nil {
		return nil, err
	}

	for _, v := range vb.videosList {
		inputVideo := filepath.Join(vb.inputPath, v+vb.inputType)
		outputVideo := filepath.Join(vb.outputPath, v+outputType)
		s := fmt.Sprintf(template, inputVideo, advance, outputVideo)
		vb.cmdBatch = append(vb.cmdBatch, s)
	}
	return vb.cmdBatch, nil
}
