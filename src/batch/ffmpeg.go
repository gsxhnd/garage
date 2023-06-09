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
	CreateDestDir(destPath string) error
	// 获取视频列表
	getVideosList() error
	// 获取字体
	getFontsParams(fontsPath string) error
	// get bactch
	GetAddSubtitleBatch(sourceSubtitleNumber int, sourceSubtitleType, sourceSubtitleLanguage, sourceSubtitleTitle, fontsPath string) ([]string, error)
	GetConvertBatch(advance, destVideoType string) ([]string, error)
}

type videoBatch struct {
	sourceRootPath  string
	sourceVideoType string
	videosList      []string
	fontsParams     string
	destPath        string
	cmdBatch        []string
	logger          *zap.Logger
}

func NewVideoBatch(l *zap.Logger, sourceRootPath, sourceVideoType string) VideoBatcher {
	return &videoBatch{
		logger:          l,
		sourceRootPath:  sourceRootPath,
		sourceVideoType: sourceVideoType,

		destPath:    "",
		videosList:  make([]string, 0),
		fontsParams: "",
		cmdBatch:    make([]string, 0),
	}
}

func (vb *videoBatch) CreateDestDir(destPath string) error {
	destDir := path.Join(destPath)
	vb.logger.Info("Start creating destination directory: " + destDir)
	if fi, err := os.Stat(destDir); err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(destDir, os.ModePerm)
		} else {
			return err
		}
	} else {
		if fi.IsDir() {
			vb.logger.Info("Destination directory already exists")
		} else {
			os.Mkdir(destDir, os.ModePerm)
		}
	}
	vb.destPath = destDir
	vb.logger.Info("Destination directory created")
	return nil
}

func (vb *videoBatch) GetAddSubtitleBatch(sourceSubtitleNumber int, sourceSubtitleType, sourceSubtitleLanguage, sourceSubtitleTitle, fontsPath string) ([]string, error) {
	if err := vb.getVideosList(); err != nil {
		return nil, err
	}

	vb.logger.Info("Source videos directory: " + vb.sourceRootPath)
	vb.logger.Info("Get matching video count: " + strconv.Itoa(len(vb.videosList)))
	vb.logger.Info("Target video's subtitle stream number: " + strconv.Itoa(sourceSubtitleNumber))
	vb.logger.Info("Target video's subtitle language: " + sourceSubtitleLanguage)
	vb.logger.Info("Target video's subtitle title: " + sourceSubtitleTitle)
	if fontsPath != "" {
		if err := vb.getFontsParams(fontsPath); err != nil {
			return nil, err
		}
		vb.logger.Info("Target video's font paths: " + fontsPath)
		vb.logger.Info(fmt.Sprintf("Attach fonts parameters: %v", vb.fontsParams))
	} else {
		vb.logger.Info("Target video's font paths not set, skip.")
	}
	vb.logger.Info("Dest video directory: " + vb.destPath)

	template := `ffmpeg.exe -i "%s" -sub_charenc UTF-8 -i "%s" -map 0 -map 1 -metadata:s:s:%v language=%v -metadata:s:s:%v title="%v" -c copy %s "%v"`
	for _, v := range vb.videosList {
		sourceVideo := filepath.Join(vb.sourceRootPath, v+vb.sourceVideoType)
		sourceSubtitle := filepath.Join(vb.sourceRootPath, v+sourceSubtitleType)
		destVideo := filepath.Join(vb.destPath, v+vb.sourceVideoType)
		s := fmt.Sprintf(template,
			sourceVideo, sourceSubtitle, sourceSubtitleNumber,
			sourceSubtitleLanguage, sourceSubtitleNumber, sourceSubtitleTitle,
			vb.fontsParams, destVideo)
		vb.cmdBatch = append(vb.cmdBatch, s)
	}
	return vb.cmdBatch, nil
}

func (vb *videoBatch) getVideosList() error {
	err := filepath.Walk(vb.sourceRootPath, func(path string, fi os.FileInfo, err error) error {
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
		if fileExt == vb.sourceVideoType {
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

	var fonts_list = make([]string, 0)
	font_exts := []string{".ttf", ".otf", ".ttc"}
	font_params_template := `-attach "%s" -metadata:s:t:%v mimetype=application/x-truetype-font`

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

		for _, b := range font_exts {
			if fileExt == b {
				fonts_list = append(fonts_list, filename)
			}
		}
		return nil
	}); err != nil {
		return err
	}
	for i, v := range fonts_list {
		vb.fontsParams += fmt.Sprintf(font_params_template, filepath.Join(fontsPath, v), i) + " "
	}
	return nil
}

func (vb *videoBatch) GetConvertBatch(advance, destVideoType string) ([]string, error) {
	template := `ffmpeg.exe -i "%v" %v "%v"`
	if err := vb.getVideosList(); err != nil {
		return nil, err
	}

	for _, v := range vb.videosList {
		sourceVideo := filepath.Join(vb.sourceRootPath, v+vb.sourceVideoType)
		destVideo := filepath.Join(vb.destPath, v+destVideoType)
		s := fmt.Sprintf(template, sourceVideo, advance, destVideo)
		vb.cmdBatch = append(vb.cmdBatch, s)
	}
	return vb.cmdBatch, nil
}
