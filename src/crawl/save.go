package crawl

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path"
	"time"

	"github.com/go-gota/gota/dataframe"
)

func (cc *crawlClient) DownloadCover(code, cover string) error {
	resp, _ := cc.httpClient.Get(cc.javbusUrl + cover)
	body, _ := ioutil.ReadAll(resp.Body)
	out, _ := os.Create("./javs/" + code + "/" + code + ".jpg")
	io.Copy(out, bytes.NewReader(body))
	return nil
}

func (cc *crawlClient) saveJavInfos() error {
	df := dataframe.LoadStructs(cc.javInfos)
	f, err := os.OpenFile(path.Join(cc.destPath, time.Now().Local().Format("2006-01-02-15-04-05")+"-jav_info.csv"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		// cc.logger.Error("")
		return err
	}
	defer f.Close()
	return df.WriteCSV(f)
}

func (cc *crawlClient) saveCovers(coverPath, code string) error {
	ext := path.Ext(coverPath)
	resp, _ := cc.httpClient.Get(cc.javbusUrl + coverPath)
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
	return nil
}
