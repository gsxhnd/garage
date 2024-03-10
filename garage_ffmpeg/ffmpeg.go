package garage_ffmpeg

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"

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
	createDestDir() error                    // 创建输出后的文件夹
	GetVideosList() ([]string, error)        // 获取视频列表s
	GetFontsList() ([]string, error)         // 获取字体列表
	GetFontsParams() (string, error)         // 获取字体列表
	GetConvertBatch() ([]string, error)      // 获取转换视频命令
	StartConvertBatch() error                // 转换视频
	GetAddFontsBatch() ([]string, error)     // 获取添加字体命令
	StartAddFontsBatch() error               // 添加字体
	GetAddSubtittleBatch() ([]string, error) //
	StartAddSubtittleBatch() error           // 添加字幕
	ExecuteBatch() error
	GetExecBatch() rxgo.Observable
}

type videoBatch struct {
	option   *VideoBatchOption
	cmdBatch []string
	Ob       *ObWriter
}

var FONT_EXT = []string{".ttf", ".otf", ".ttc"}

const CONVERT_TEMPLATE = `ffmpeg.exe -i "%v" %v "%v"`
const ADD_SUB_TEMPLATE = `ffmpeg.exe -i "%s" -sub_charenc UTF-8 -i "%s" -map 0 -map 1 -metadata:s:s:%v language=%v -metadata:s:s:%v title="%v" -c copy %s "%v"`
const ADD_FONT_TEMPLATE = `ffmpeg.exe -i "%s" -c copy %s "%v"`
const FONT_TEMPLATE = `-attach "%s" -metadata:s:t:%v mimetype=application/x-truetype-font `

func NewVideoBatch(opt *VideoBatchOption) (VideoBatcher, error) {
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
		// vb.logger.Debug("get video filename: " + filename)

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
				fontsList = append(fontsList, filename)
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
		return "", nil
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
		inputVideo := filepath.Join(vb.option.InputPath, v+vb.option.InputFormat)
		outputVideo := filepath.Join(vb.option.OutputPath, v+vb.option.OutputFormat)
		s := fmt.Sprintf(CONVERT_TEMPLATE, inputVideo, vb.option.Advance, outputVideo)
		vb.cmdBatch = append(vb.cmdBatch, s)
	}

	return vb.cmdBatch, nil
}

func (vb *videoBatch) StartConvertBatch() error {
	_, err := vb.GetConvertBatch()
	if err != nil {
		return err
	}
	return vb.ExecuteBatch()
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
		sourceVideo := filepath.Join(vb.option.InputPath, v+vb.option.InputFormat)
		destVideo := filepath.Join(vb.option.OutputPath, v+vb.option.InputFormat)
		s := fmt.Sprintf(ADD_FONT_TEMPLATE, sourceVideo, fontsParams, destVideo)
		vb.cmdBatch = append(vb.cmdBatch, s)
	}

	return vb.cmdBatch, nil
}

func (vb *videoBatch) StartAddFontsBatch() error {
	// if fontsList, err := vb.GetFontsList(); err != nil {
	// 	return err
	// }

	// vb.logger.Info("Source videos directory: " + vb.option.InputPath)
	// vb.logger.Info("Get matching video count: " + strconv.Itoa(len(vb.videosList)))
	// vb.logger.Info("Target video's font paths: " + vb.option.FontsPath)
	// vb.logger.Info(fmt.Sprintf("Attach fonts parameters: %v", vb.fontsParams))
	// vb.logger.Info("Dest video directory: " + vb.option.OutputPath)

	if !vb.option.Exec {
		return nil
	} else {
		_, err := vb.GetConvertBatch()
		if err != nil {
			return err
		}
		return vb.ExecuteBatch()
	}
}

func (vb *videoBatch) GetAddSubtittleBatch() ([]string, error) {
	videosList, err := vb.GetVideosList()
	if err != nil {
		return nil, err
	}

	// if vb.option.FontsPath != "" {
	// }

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

func (vb *videoBatch) StartAddSubtittleBatch() error {

	// vb.logger.Debug("Source videos directory: " + vb.option.InputPath)
	// vb.logger.Debug("Get matching video count: " + strconv.Itoa(len(vb.videosList)))
	// vb.logger.Debug("Target video's subtitle stream number: " + strconv.Itoa(vb.option.InputSubNo))
	// vb.logger.Debug("Target video's subtitle language: " + vb.option.InputSubLang)
	// vb.logger.Debug("Target video's subtitle title: " + vb.option.InputSubTitle)

	// vb.logger.Info("Target video's font paths: " + vb.option.FontsPath)
	// vb.logger.Info(fmt.Sprintf("Attach fonts parameters: %v", vb.fontsParams))
	// vb.logger.Info("Target video's font paths not set, skip.")
	// vb.logger.Info("Dest video directory: " + vb.option.OutputPath)

	// vb.logger.Info("Get all videos, starting convert")

	if !vb.option.Exec {
		return nil
	} else {
		_, err := vb.GetAddSubtittleBatch()
		if err != nil {
			return err
		}
		return vb.ExecuteBatch()
	}
}

func (vb *videoBatch) createDestDir() error {
	destDir := path.Join(vb.option.OutputPath)
	// vb.logger.Info("Start creating destination directory: " + destDir)
	if fi, err := os.Stat(destDir); err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(destDir, os.ModePerm)
		} else {
			return err
		}
	} else {
		if fi.IsDir() {
			return errors.New("destination directory already exists")
			// vb.logger.Info("Destination directory already exists")
		}
	}
	// vb.logger.Info("Destination directory created")
	return nil
}

func (vb *videoBatch) ExecuteBatch() error {
	if len(vb.cmdBatch) == 0 {
		vb.cmdBatch = append(vb.cmdBatch, "-al")
		vb.cmdBatch = append(vb.cmdBatch, "-al")
		vb.cmdBatch = append(vb.cmdBatch, "-al")
		vb.cmdBatch = append(vb.cmdBatch, "-al")
	}

	fmt.Println(vb.cmdBatch, vb.option.Exec)

	for _, cmd := range vb.cmdBatch {
		if !vb.option.Exec {
			// vb.logger.Sugar().Infof("cmd: %v", cmd)
			return nil
		}

		// startTime := time.Now()
		// vb.logger.Sugar().Infof("Start convert video cmd: %v", cmd)
		cmd := exec.Command("ls", cmd)
		cmd.Stdout = vb.Ob
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			// vb.logger.Sugar().Errorf("cmd error: %v", err)
			return err
		}
		// vb.logger.Sugar().Infof("Finished convert video, spent time: %v sec", time.Since(startTime).Seconds())
	}
	return nil
}

func (vb *videoBatch) GetExecBatch() rxgo.Observable {
	return vb.Ob.ob
}

type ObWriter struct {
	ob rxgo.Observable
	ch chan rxgo.Item
}

func ObWriterNew() *ObWriter {
	ch := make(chan rxgo.Item)
	return &ObWriter{
		ob: rxgo.FromChannel(ch),
		ch: ch,
	}
}

func (o *ObWriter) Write(p []byte) (int, error) {
	o.ch <- rxgo.Of(p)
	return len(p), nil
}
