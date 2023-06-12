package cmd

import (
	"github.com/gsxhnd/garage/src/crawl"
	"github.com/gsxhnd/garage/src/utils"
	"github.com/urfave/cli/v2"
)

var (
	javProxyFlag  = &cli.StringFlag{Name: "proxy", Usage: "代理配置,如: http://127.0.0.1:1080"}
	javOutputFlag = &cli.StringFlag{Name: "output", Usage: "设置下载目录", Value: "./javbus"}
	javSiteFlag   = &cli.StringFlag{
		Name:        "site",
		Usage:       "选择爬取数据的网站,支持网站(javbus)",
		Destination: nil,
		HasBeenSet:  false,
		Value:       "javbus",
	}
)

var javCodeCmd = &cli.Command{
	Name:      "jav-code",
	Aliases:   nil,
	Usage:     "根据指定番号爬取数据",
	UsageText: "jav_code --site [javbus/javlibrary] XXX-001",
	Flags: []cli.Flag{
		javProxyFlag,
		javSiteFlag,
		javOutputFlag,
	},
	Action: func(ctx *cli.Context) error {
		var (
			logger = utils.GetLogger()
			code   = ctx.Args().Get(0)
		)

		c, err := crawl.NewCrawlClient(logger, crawl.CrawlOptions{
			Proxy:    ctx.String("proxy"),
			DestPath: ctx.String("output"),
		})
		if err != nil {
			logger.Panic("client init error: " + err.Error())
			return err
		}

		if err := c.StartCrawlJavbusMovie(code); err != nil {
			logger.Panic("crawl error: " + err.Error())
			return err
		}
		return nil
	},
}

var javPrefixCmd = &cli.Command{
	Name:  "jav-prefix-code",
	Usage: "根据番号前缀爬取数据",
	Flags: []cli.Flag{
		javProxyFlag,
		javSiteFlag,
		javOutputFlag,
		&cli.StringFlag{Name: "prefix-code", Required: true},
		&cli.IntFlag{Name: "prefix-min", Required: true, Value: 1},
		&cli.IntFlag{Name: "prefix-max", Required: true, Value: 5},
	},
	Action: func(c *cli.Context) error {
		var logger = utils.GetLogger()

		client, err := crawl.NewCrawlClient(logger, crawl.CrawlOptions{
			Proxy:       c.String("proxy"),
			DestPath:    c.String("output"),
			PrefixCode:  c.String("prefix-code"),
			PrefixMinNo: c.Int("prefix-min"),
			PrefixMaxNo: c.Int("prefix-max"),
		})

		if err != nil {
			logger.Panic("client init error: " + err.Error())
			return err
		}

		if err := client.StartCrawlJavbusMovieByPrefix(); err != nil {
			logger.Panic("crawl error: " + err.Error())
			return err
		}
		return nil
	},
}

var javStarCodeCmd = &cli.Command{
	Name:  "jav-star-code",
	Usage: "根据演员ID爬取数据",
	Flags: []cli.Flag{
		javSiteFlag,
		javOutputFlag,
		javProxyFlag,
	},
	Action: func(c *cli.Context) error {
		return nil
	},
}
