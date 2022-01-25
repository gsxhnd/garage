package command

import (
	"github.com/urfave/cli/v2"
)

var (
	proxyFlag = &cli.StringFlag{Name: "proxy", Usage: "代理配置"}
)
