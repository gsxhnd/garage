package main

import "github.com/gsxhnd/garage/internal/jav"

var crawlJavbusCmd = newSiteCrawlCommand(
	jav.SourceJavbus,
	"crawl_javbus",
	"从 javbus 爬取番号信息、封面和磁力链接",
	"./javbus",
	"javbus_code",
	"javbus_prefix_code",
	"javbus_star_code",
)
