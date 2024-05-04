package handler

import "github.com/gsxhnd/garage/utils"

type CrawlHandler interface{}

type crawlHnadler struct {
	logger utils.Logger
}

func NewCrawlHandler(l utils.Logger) CrawlHandler {
	return crawlHnadler{
		logger: l,
	}
}
