package batch

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

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
		vs := strings.Split(fi.Name(), ".")
		if vs[1] == vb.SourceVideoType {
			vl = append(vl, vs[0])
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
		sourceVideo := filepath.Join(vb.SourceRootPath, v+"."+vb.SourceVideoType)
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
		sourceVideo := filepath.Join(vb.SourceRootPath, v+"."+vb.SourceVideoType)
		s := fmt.Sprintf(t, sourceVideo, vb.Advance, vb.DestPath+v+vb.DestVideoType)
		batch = append(batch, s)
	}
	return batch
}
