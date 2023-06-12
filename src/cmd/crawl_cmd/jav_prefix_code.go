package crawl_cmd

import (
	"github.com/gsxhnd/garage/src/crawl"
	"github.com/gsxhnd/garage/src/utils"
	"github.com/urfave/cli/v2"
)

var JavPrefixCmd = &cli.Command{
	Name:  "jav_prefix_code",
	Usage: "根据番号前缀爬取数据",
	Flags: []cli.Flag{
		proxyFlag,
		siteFlag,
		destDirFlag,
		prefixCodeFlag,
		prefixMinNoFlag,
		prefixMaxNoFlag,
	},
	Action: func(c *cli.Context) error {
		var logger = utils.GetLogger()

		client, err := crawl.NewCrawlClient(logger, crawl.CrawlOptions{
			Proxy:       c.String("proxy"),
			DestPath:    c.String("dest_dir"),
			PrefixCode:  c.String("prefix_code"),
			PrefixMinNo: c.Int("prefix_min_no"),
			PrefixMaxNo: c.Int("prefix_max_no"),
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
