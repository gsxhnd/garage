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

type VideoBatch struct {
	SourceRootPath     string
	SourceVideoType    string
	SourceSubtitleType string
	DestPath           string
	DestVideoType      string
	Advance            string
	Logger             *zap.Logger
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

func (vb *VideoBatch) GetSubtitleBatch(videos []string) []string {
	t := `ffmpeg.exe -i "%v" -sub_charenc UTF-8 -i "%v" -metadata:s:s:0 language=chi -metadata:s:s:0 title="Chinese" -c copy %v "%v"`
	var batch = []string{}
	for _, v := range videos {
		sourceVideo := filepath.Join(vb.SourceRootPath, v+vb.SourceVideoType)
		sourceSubtitle := filepath.Join(vb.SourceRootPath, v+vb.SourceSubtitleType)
		s := fmt.Sprintf(t, sourceVideo, sourceSubtitle, vb.Advance, vb.DestPath+v+vb.DestVideoType)
		batch = append(batch, s)
	}
	return batch
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
	destDir := path.Join(vb.SourceRootPath, vb.DestPath)
	vb.Logger.Info("Start creating destination directory" + destDir)
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
