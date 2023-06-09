package jav

func (cc *CrawlClient) StarCrawlJavbusMovie(code string) {
	info, err := cc.DownloadInfo(code)
	if err != nil {
		return
	}
	err = cc.DownloadCover(info.Code, info.Cover)
	if err != nil {
		return
	}
}

func (cc *CrawlClient) StarCrawlJavbusMovieByPrefix(prefixCode string) {}
func (cc *CrawlClient) StarCrawlJavbusMovieByStar(starCode string)     {}
