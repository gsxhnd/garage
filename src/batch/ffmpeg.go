package batch

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"go.uber.org/zap"
)

type VideoBatcher interface {
	GetVideosList(sourceRootPath, sourceVideoType string) ([]string, error)
	// 获取字体
	GetFontsParams(fontPath string) (string, error)
	CreateDestDir(destPath string) error
	GetImportSubtitleBatch(videos []string, fonts string) []string
}
type videoBatch struct {
	logger *zap.Logger
}

func NewVideoBatch(l *zap.Logger) VideoBatcher {
	return &videoBatch{
		logger: l,
	}
}

func (vb *videoBatch) GetVideosList(sourceRootPath, sourceVideoType string) ([]string, error) {
	var vl = []string{}
	err := filepath.Walk(sourceRootPath, func(path string, fi os.FileInfo, err error) error {
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
		if fileExt == sourceVideoType {
			fileName := strings.TrimSuffix(filename, fileExt)
			vl = append(vl, fileName)
			return nil
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return vl, err
}

func (vb *videoBatch) GetFontsParams(fontPath string) (string, error) {
	if fontPath == "" {
		return "", nil
	}

	var fonts_list = make([]string, 0)
	font_exts := []string{".ttf", ".otf", ".ttc"}
	font_params_template := `-attach "%s" -metadata:s:t:%v mimetype=application/x-truetype-font`

	var font_params string
	if err := filepath.Walk(fontPath, func(path string, fi os.FileInfo, err error) error {
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
		return "", err
	}
	for i, v := range fonts_list {
		font_params += fmt.Sprintf(font_params_template, filepath.Join(fontPath, v), i) + " "
	}
	return font_params, nil
}

func (vb *videoBatch) CreateDestDir(destPath string) error {
	destDir := path.Join(destPath)
	vb.logger.Info("Start creating destination directory: " + destDir)
	time.Sleep(500 * time.Millisecond)
	if fi, err := os.Stat(destDir); err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(destDir, os.ModePerm)
		} else {
			return err
		}
	} else {
		if fi.IsDir() {
			vb.logger.Info("Destination directory already exists")
			return nil
		} else {
			os.Mkdir(destDir, os.ModePerm)
		}
	}
	vb.logger.Info("Destination directory created")
	return nil
}

func (vb *videoBatch) GetImportSubtitleBatch(videos []string, fonts string) []string {
	// t := `ffmpeg.exe -i "%s" -sub_charenc UTF-8 -i "%s" -map 0 -map 1 -metadata:s:s:%v language=%v -metadata:s:s:%v title="%v" -c copy %s %s "%v"`
	// var (
	// 	batch        = []string{}
	// 	fonts_params string
	// )
	// vb.Logger.Info("Source videos directory: " + vb.SourceRootPath)
	// vb.Logger.Info("Get matching video count: " + strconv.Itoa(len(videos)))
	// vb.Logger.Info("Target video's subtitle stream number: " + strconv.Itoa(vb.SourceSubtitleNumber))
	// vb.Logger.Info("Target video's subtitle language: " + vb.SourceSubtitleLanguage)
	// vb.Logger.Info("Target video's subtitle title: " + vb.SourceSubtitleTitle)
	// if vb.FontsPath != "" {
	// 	 vb.Logger.Info("Target video's font paths: " + vb.FontsPath)
	// 	fonts_params, _ = vb.GetFontsParams()
	// 	vb.Logger.Debug(fmt.Sprintf("Attach fonts parameters: %v", fonts_params))
	// }
	// vb.Logger.Info("Dest video directory: " + vb.DestPath)

	// for _, v := range videos {
	// 	sourceVideo := filepath.Join(vb.SourceRootPath, v+vb.SourceVideoType)
	// 	sourceSubtitle := filepath.Join(vb.SourceRootPath, v+vb.SourceSubtitleType)
	// 	destVideo := filepath.Join(vb.DestPath, v+vb.DestVideoType)
	// 	s := fmt.Sprintf(t,
	// 		sourceVideo, sourceSubtitle, vb.SourceSubtitleNumber,
	// 		vb.SourceSubtitleLanguage, vb.SourceSubtitleNumber, vb.SourceSubtitleTitle,
	// 		fonts_params, vb.Advance, destVideo)
	// 	batch = append(batch, s)
	// }
	// return batch
	return nil
}

type VideoBatch struct {
	SourceRootPath         string
	SourceVideoType        string
	SourceSubtitleType     string
	SourceSubtitleNumber   int
	SourceSubtitleLanguage string
	SourceSubtitleTitle    string
	FontsPath              string
	DestPath               string
	DestVideoType          string
	Advance                string
	Logger                 *zap.Logger
}

func (vb *VideoBatch) GetVideos() ([]string, error) {
	var vl = []string{}
	err := filepath.Walk(vb.SourceRootPath, func(path string, fi os.FileInfo, err error) error {
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
		if fileExt == vb.SourceVideoType {
			fileName := strings.TrimSuffix(filename, fileExt)
			vl = append(vl, fileName)
			return nil
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return vl, err
}

// 获取字体
func (vb *VideoBatch) GetFontsParams() (string, error) {
	var fonts_list = make([]string, 0)
	font_exts := []string{".ttf", ".otf", ".ttc"}
	font_params_template := `-attach "%s" -metadata:s:t:%v mimetype=application/x-truetype-font`
	var font_params string
	if err := filepath.Walk(vb.FontsPath, func(path string, fi os.FileInfo, err error) error {
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
		return "", err
	}
	for i, v := range fonts_list {
		font_params += fmt.Sprintf(font_params_template, filepath.Join(vb.FontsPath, v), i) + " "
	}
	return font_params, nil
}

func (vb *VideoBatch) GetConvertBatch(videos []string) []string {
	t := `ffmpeg.exe -i "%v" %v "%v"`
	var batch = []string{}
	for _, v := range videos {
		sourceVideo := filepath.Join(vb.SourceRootPath, v+vb.SourceVideoType)
		s := fmt.Sprintf(t, sourceVideo, vb.Advance, vb.DestPath+v+vb.DestVideoType)
		batch = append(batch, s)
	}
	return batch
}

func (vb *VideoBatch) CreateDestDir() error {
	destDir := path.Join(vb.DestPath)
	vb.Logger.Info("Start creating destination directory: " + destDir)
	time.Sleep(500 * time.Millisecond)
	if fi, err := os.Stat(destDir); err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(destDir, os.ModePerm)
		} else {
			return err
		}
	} else {
		if fi.IsDir() {
			vb.Logger.Info("Destination directory already exists")
			return nil
		} else {
			os.Mkdir(destDir, os.ModePerm)
		}
	}
	vb.Logger.Info("Destination directory created")
	return nil
}
