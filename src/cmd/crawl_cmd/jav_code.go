package crawl_cmd

import (
	"github.com/gsxhnd/garage/src/crawl"
	"github.com/gsxhnd/garage/src/utils"
	"github.com/urfave/cli/v2"
)

var JavCodeCmd = &cli.Command{
	Name:      "jav_code",
	Aliases:   nil,
	Usage:     "根据指定番号爬取数据",
	UsageText: "jav_code --site [javbus/javlibrary] XXX-001",
	Flags: []cli.Flag{
		proxyFlag,
		siteFlag,
		destDirFlag,
	},
	Action: func(ctx *cli.Context) error {
		var (
			logger = utils.GetLogger()
			code   = ctx.Args().Get(0)
		)

		c, err := crawl.NewCrawlClient(logger, crawl.CrawlOptions{
			Proxy:    ctx.String("proxy"),
			DestPath: ctx.String("dest_dir"),
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
