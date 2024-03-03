package main

import (
	"github.com/gsxhnd/garage/garage_jav"
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
	javMagnetFlag = &cli.BoolFlag{
		Name:  "magnet",
		Usage: "保存磁力链接,开启参数 --magnet",
		Value: false,
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
		javMagnetFlag,
		javOutputFlag,
	},
	Action: func(ctx *cli.Context) error {
		opt := &garage_jav.JavCrawlConfig{
			Code:           ctx.Args().Get(0),
			DownloadMagent: ctx.Bool("magnet"),
			Proxy:          ctx.String("proxy"),
			DestPath:       ctx.String("output"),
		}

		c, err := garage_jav.NewJavbusCrawl(logger, opt)
		if err != nil {
			logger.Panicw("client init error: " + err.Error())
			return err
		}

		if err := c.StartCrawlJavbusMovie(); err != nil {
			logger.Panicw("crawl error: " + err.Error())
			return err
		}

		// s := crawl.NewJavSave(logger, ctx.String("output"), nil)
		// s.Save(false, false)
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
		javMagnetFlag,
		&cli.StringFlag{Name: "prefix-code", Value: "EKDV", Usage: "番号前缀"},
		&cli.IntFlag{Name: "prefix-min", Value: 1, Usage: "番号开始编号"},
		&cli.IntFlag{Name: "prefix-max", Value: 5, Usage: "番号结束编号"},
	},
	Action: func(ctx *cli.Context) error {

		opt := &garage_jav.JavCrawlConfig{
			Proxy:          ctx.String("proxy"),
			DestPath:       ctx.String("output"),
			DownloadMagent: ctx.Bool("magnet"),
			PrefixCode:     ctx.String("prefix-code"),
			PrefixMinNo:    ctx.Uint("prefix-min"),
			PrefixMaxNo:    ctx.Uint("prefix-max"),
		}

		c, err := garage_jav.NewJavbusCrawl(logger, opt)
		if err != nil {
			logger.Panicw("client init error: " + err.Error())
			return err
		}

		if err := c.StartCrawlJavbusMovieByPrefix(); err != nil {
			logger.Panicw("crawl error: " + err.Error())
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
		javMagnetFlag,
		javOutputFlag,
		javProxyFlag,
		&cli.StringFlag{Name: "star-code", Value: "vfn", Usage: "演员番号"},
	},
	Action: func(ctx *cli.Context) error {

		opt := &garage_jav.JavCrawlConfig{
			DownloadMagent: ctx.Bool("magnet"),
			Proxy:          ctx.String("proxy"),
			DestPath:       ctx.String("output"),
			StarCode:       ctx.String("star-code"),
		}

		c, err := garage_jav.NewJavbusCrawl(logger, opt)
		if err != nil {
			logger.Panicw("client init error: " + err.Error())
			return err
		}

		if err := c.StartCrawlJavbusMovieByStar(); err != nil {
			logger.Panicw("crawl error: " + err.Error())
			return err
		}
		return nil
	},
}

var javStarCodeFromDirCmd = &cli.Command{
	Name:  "jav-code-from-dir",
	Usage: "根据演员ID爬取数据",
	Flags: []cli.Flag{
		javSiteFlag,
		javOutputFlag,
		javProxyFlag,
		&cli.StringFlag{Name: "input", Required: true},
	},
	Action: func(ctx *cli.Context) error {
		opt := &garage_jav.JavCrawlConfig{
			DownloadMagent: ctx.Bool("magnet"),
			Proxy:          ctx.String("proxy"),
			DestPath:       ctx.String("output"),
		}

		c, err := garage_jav.NewJavbusCrawl(logger, opt)
		if err != nil {
			logger.Panicw("client init error: " + err.Error())
			return err
		}

		if err := c.StartCrawlJavbusMovieByFilepath(ctx.String("input")); err != nil {
			logger.Panicw("crawl error: " + err.Error())
			return err
		}
		return nil
	},
}
