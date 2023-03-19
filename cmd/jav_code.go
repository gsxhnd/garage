package cmd

import (
	"path"

	"github.com/gsxhnd/garage/crawl"
	"github.com/gsxhnd/garage/utils"
	"github.com/urfave/cli/v2"
)

var javCodeCmd = &cli.Command{
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
			logger  = utils.GetLogger()
			code    = ctx.Args().Get(0)
			proxy   = ctx.String("proxy")
			destDir = "./javs"
		)

		if err := utils.MkdirDestDir(path.Join(destDir, code)); err != nil {
			logger.Panic("目录创建失败， Error: " + err.Error())
			return err
		}

		c := crawl.NewCrawlClient(logger)
		_ = c.SetProxy(proxy)
		c.StarCrawlJavbusMovie(code)
		return nil
	},
}
