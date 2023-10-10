package crawl

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"time"

	"github.com/go-gota/gota/dataframe"
	"go.uber.org/zap"
)

type javSave struct {
	logger     *zap.Logger
	destPath   string
	javInfos   []JavMovie
	httpClient *http.Client
}

func NewJavSave(l *zap.Logger, dp string, infos []JavMovie) *javSave {
	return &javSave{
		logger:     l,
		destPath:   dp,
		javInfos:   infos,
		httpClient: &http.Client{},
	}
}

func (s *javSave) Save(cover, magnet bool) error {
	s.saveJavInfos()
	if magnet {
		s.saveMagents()
	}
	if cover {
		s.saveCovers("", "", "")
	}
	return nil
}

func (s *javSave) saveJavInfos() error {
	df := dataframe.LoadStructs(s.javInfos)
	f, err := os.OpenFile(path.Join(s.destPath, time.Now().Local().Format("2006-01-02-15-04-05")+"-jav_info.csv"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		s.logger.Error("Save jav info file failed error: %s" + err.Error())
		return err
	}
	defer f.Close()
	return df.WriteCSV(f)
}

func (s *javSave) saveMagents() error {
	if len(s.javInfos) == 0 {
		return nil
	}

	f, err := os.OpenFile(path.Join(s.destPath, time.Now().Local().Format("2006-01-02-15-04-05")+"-jav_magnet.text"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		s.logger.Error("Save jav info file failed error: %s" + err.Error())
		return err
	}
	defer f.Close()

	for _, v := range s.javInfos {
		f.WriteString(v.Code + "\n")
	}
	return nil
}

func (s *javSave) saveCovers(host, coverPath, code string) error {
	var urlImg = ""
	u, err := url.ParseRequestURI(coverPath)
	if err != nil {
		s.logger.Error("parse cover path failed error: %s" + err.Error())
		return err
	}
	if u.Host == "" {
		urlImg = host + coverPath
	} else {
		urlImg = coverPath
	}

	s.logger.Info("downloading coverage url: " + urlImg)
	ext := path.Ext(urlImg)
	resp, err := s.httpClient.Get(urlImg)
	if err != nil {
		s.logger.Error("downloading coverage error: " + err.Error())
		return err
	}
	body, _ := ioutil.ReadAll(resp.Body)

	f, err := os.OpenFile(path.Join(s.destPath, "cover", code+ext), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := io.Copy(f, bytes.NewReader(body)); err != nil {
		return err
	}
	return nil
}
