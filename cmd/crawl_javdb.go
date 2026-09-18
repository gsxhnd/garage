package main

import "github.com/gsxhnd/garage/internal/jav"

var crawlJavDBCmd = newSiteCrawlCommand(
	jav.SourceJavdb,
	"crawl_javdb",
	"从 javdb 爬取番号信息、封面和磁力链接",
	"./javdb",
	"javdb_code",
	"javdb_prefix_code",
	"javdb_star_code",
)
