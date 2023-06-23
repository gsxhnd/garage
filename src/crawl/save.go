package crawl

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/url"
	"os"
	"path"
	"time"

	"github.com/go-gota/gota/dataframe"
)

func (cc *crawlClient) saveJavInfos() error {
	df := dataframe.LoadStructs(cc.javInfos)
	f, err := os.OpenFile(path.Join(cc.destPath, time.Now().Local().Format("2006-01-02-15-04-05")+"-jav_info.csv"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		cc.logger.Error("Save jav info file failed error: %s" + err.Error())
		return err
	}
	defer f.Close()
	return df.WriteCSV(f)
}

func (cc *crawlClient) saveCovers(coverPath, code string) error {
	var urlImg = ""
	u, err := url.ParseRequestURI(coverPath)
	if err != nil {
		cc.logger.Error("parse cover path failed error: %s" + err.Error())
		return err
	}
	if u.Host == "" {
		urlImg = cc.javbusUrl + coverPath
	} else {
		urlImg = coverPath
	}

	cc.logger.Info("downloading coverage url: " + urlImg)
	ext := path.Ext(urlImg)
	resp, err := cc.httpClient.Get(urlImg)
	if err != nil {
		cc.logger.Error("downloading coverage error: " + err.Error())
		return err
	}
	body, _ := ioutil.ReadAll(resp.Body)

	f, err := os.OpenFile(path.Join(cc.destPath, "cover", code+ext), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := io.Copy(f, bytes.NewReader(body)); err != nil {
		return err
	}
	return nil
}

func (cc *crawlClient) saveMagents() error {
	if len(cc.javMagnets) == 0 {
		return nil
	}

	f, err := os.OpenFile(path.Join(cc.destPath, time.Now().Local().Format("2006-01-02-15-04-05")+"-jav_magnet.text"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		cc.logger.Error("Save jav info file failed error: %s" + err.Error())
		return err
	}
	defer f.Close()

	for _, v := range cc.javMagnets {
		f.WriteString(v + "\n")
	}
	return nil
}
