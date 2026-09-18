package main

import (
	"context"
	"path/filepath"
	"strings"

	"github.com/gsxhnd/garage/internal/jav"
	"github.com/gsxhnd/garage/utils"
	"github.com/urfave/cli/v2"
)

func crawlFlags(defaultOut string) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{Name: "proxy", Usage: "代理配置,如: http://127.0.0.1:1080"},
		&cli.StringFlag{Name: "cookie", Usage: "可选 Cookie, 如 JavDB 的 _jdb_session"},
		&cli.StringFlag{Name: "output", Usage: "设置下载目录", Value: defaultOut},
		&cli.BoolFlag{Name: "magnet", Usage: "保存磁力链接", Value: false},
		&cli.BoolFlag{Name: "cover", Usage: "下载封面", Value: true},
		&cli.IntFlag{Name: "parallel", Usage: "并发抓取数量", Value: 1},
		&cli.DurationFlag{Name: "delay", Usage: "请求间隔", Value: 0},
	}
}

func runCrawl(ctx *cli.Context, source jav.Source, query jav.Query) error {
	cctx := ctx.Context
	if cctx == nil {
		cctx = context.Background()
	}

	cfg := jav.Config{
		Proxy:          ctx.String("proxy"),
		Cookie:         ctx.String("cookie"),
		OutDir:         ctx.String("output"),
		DownloadMagnet: ctx.Bool("magnet"),
		DownloadCover:  ctx.Bool("cover"),
		Parallelism:    ctx.Int("parallel"),
		RandomDelay:    ctx.Duration("delay"),
	}

	crawler, err := jav.NewCrawler(source, logger, cfg)
	if err != nil {
		return err
	}
	movies, err := crawler.Crawl(cctx, query)
	if err != nil {
		return err
	}
	store, err := jav.NewLocalStore(logger, cfg)
	if err != nil {
		return err
	}
	return store.Save(cctx, movies)
}

func newSiteCrawlCommand(source jav.Source, name, usage, defaultOut, codeName, prefixName, starName string) *cli.Command {
	return &cli.Command{
		Name:  name,
		Usage: usage,
		Flags: crawlFlags(defaultOut),
		Before: func(ctx *cli.Context) error {
			return utils.MakeDir(filepath.Join(ctx.String("output"), "cover"))
		},
		Subcommands: []*cli.Command{
			{
				Name:      codeName,
				Aliases:   []string{"code"},
				Usage:     "根据指定番号爬取数据",
				UsageText: name + " " + codeName + " XXX-001",
				Action: func(ctx *cli.Context) error {
					code := strings.TrimSpace(ctx.Args().Get(0))
					if code == "" {
						return cli.Exit("missing code", 1)
					}
					return runCrawl(ctx, source, jav.Query{Codes: []string{code}})
				},
			},
			{
				Name:    prefixName,
				Aliases: []string{"prefix"},
				Usage:   "根据番号前缀爬取数据",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "prefix_code", Value: "EKDV", Usage: "番号前缀"},
					&cli.Uint64Flag{Name: "prefix_min", Value: 1, Usage: "番号开始编号"},
					&cli.Uint64Flag{Name: "prefix_max", Value: 5, Usage: "番号结束编号"},
					&cli.Uint64Flag{Name: "prefix_zero", Value: 3, Usage: "番号补零位数"},
				},
				Action: func(ctx *cli.Context) error {
					return runCrawl(ctx, source, jav.Query{
						Prefixes:   []string{ctx.String("prefix_code")},
						PrefixMin:  ctx.Uint64("prefix_min"),
						PrefixMax:  ctx.Uint64("prefix_max"),
						PrefixZero: ctx.Uint64("prefix_zero"),
					})
				},
			},
			{
				Name:    starName,
				Aliases: []string{"star"},
				Usage:   "根据演员ID爬取数据",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "star_code", Value: "vfn", Usage: "演员ID"},
				},
				Action: func(ctx *cli.Context) error {
					return runCrawl(ctx, source, jav.Query{
						StarIDs: []string{ctx.String("star_code")},
					})
				},
			},
			{
				Name:    "code_from_dir",
				Aliases: []string{"from-dir"},
				Usage:   "根据本地视频文件名爬取数据",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "input_path", Required: true},
				},
				Action: func(ctx *cli.Context) error {
					return runCrawl(ctx, source, jav.Query{
						VideosDir: ctx.String("input_path"),
					})
				},
			},
		},
	}
}
