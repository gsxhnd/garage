package main

import (
	"github.com/gsxhnd/garage/garage_jav"
	"github.com/urfave/cli/v2"
)

var crawlJavbusCmd = &cli.Command{
	Name:        "crawl_javbus",
	Description: "",
	Flags: []cli.Flag{
		&cli.StringFlag{Name: "proxy", Usage: "代理配置,如: http://127.0.0.1:1080"},
		&cli.StringFlag{Name: "output", Usage: "设置下载目录", Value: "./javbus"},
		&cli.BoolFlag{
			Name:  "magnet",
			Usage: "保存磁力链接,开启参数 --magnet",
			Value: false,
		},
	},
	Subcommands: []*cli.Command{
		javbusCodeCmd,
		javbusPrefixCmd,
		javbusStarCodeCmd,
		javbusCodeFromDirCmd,
	},
}

var javbusCodeCmd = &cli.Command{
	Name:      "javbus_code",
	Aliases:   nil,
	Usage:     "根据指定番号爬取数据",
	UsageText: "jav_code --site [javbus/javlibrary] XXX-001",
	Flags:     []cli.Flag{},
	Action: func(ctx *cli.Context) error {
		opt := &garage_jav.JavbusCrawlConfig{
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
		_, err = c.GetJavbusMovie()
		if err != nil {
			return err
		}
		return c.SaveLocal()
	},
}

var javbusPrefixCmd = &cli.Command{
	Name:  "javbus_prefix_code",
	Usage: "根据番号前缀爬取数据",
	Flags: []cli.Flag{
		&cli.StringFlag{Name: "prefix_code", Value: "EKDV", Usage: "番号前缀"},
		&cli.Uint64Flag{Name: "prefix_min", Value: 1, Usage: "番号开始编号"},
		&cli.Uint64Flag{Name: "prefix_max", Value: 5, Usage: "番号结束编号"},
		&cli.Uint64Flag{Name: "prefix_zero", Value: 3, Usage: "番号结束编号"},
	},
	Action: func(ctx *cli.Context) error {
		opt := &garage_jav.JavbusCrawlConfig{
			Proxy:          ctx.String("proxy"),
			DestPath:       ctx.String("output"),
			DownloadMagent: ctx.Bool("magnet"),
			PrefixCode:     ctx.String("prefix_code"),
			PrefixMinNo:    ctx.Uint64("prefix_min"),
			PrefixMaxNo:    ctx.Uint64("prefix_max"),
			PrefixZero:     ctx.Uint64("prefix_zero"),
		}

		c, err := garage_jav.NewJavbusCrawl(logger, opt)
		if err != nil {
			logger.Panicw("client init error: " + err.Error())
			return err
		}
		if _, err := c.GetJavbusMovieByPrefix(); err != nil {
			return err
		}
		return c.SaveLocal()
	},
}

var javbusStarCodeCmd = &cli.Command{
	Name:  "javbus_star_code",
	Usage: "根据演员ID爬取数据",
	Flags: []cli.Flag{
		&cli.StringFlag{Name: "star_code", Value: "vfn", Usage: "演员番号"},
	},
	Action: func(ctx *cli.Context) error {
		opt := &garage_jav.JavbusCrawlConfig{
			DownloadMagent: ctx.Bool("magnet"),
			Proxy:          ctx.String("proxy"),
			DestPath:       ctx.String("output"),
			StarCode:       ctx.String("star_code"),
		}

		c, err := garage_jav.NewJavbusCrawl(logger, opt)
		if err != nil {
			logger.Panicw("client init error: " + err.Error())
			return err
		}
		if _, err := c.GetJavbusMovieByStar(); err != nil {
			return err
		}
		return c.SaveLocal()
	},
}

var javbusCodeFromDirCmd = &cli.Command{
	Name:  "javbus_code_from_dir",
	Usage: "根据演员ID爬取数据",
	Flags: []cli.Flag{
		&cli.StringFlag{Name: "input_path", Required: true},
	},
	Action: func(ctx *cli.Context) error {
		opt := &garage_jav.JavbusCrawlConfig{
			DownloadMagent: ctx.Bool("magnet"),
			Proxy:          ctx.String("proxy"),
			DestPath:       ctx.String("output"),
		}

		_, err := garage_jav.NewJavbusCrawl(logger, opt)
		if err != nil {
			logger.Panicw("client init error: " + err.Error())
			return err
		}

		// if err := c.StartCrawlJavbusMovieByFilepath(ctx.String("input")); err != nil {
		// 	logger.Panicw("crawl error: " + err.Error())
		// 	return err
		// }
		return nil
	},
}
