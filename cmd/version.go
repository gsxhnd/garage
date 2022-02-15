package cmd

import (
	"fmt"
	"time"

	"github.com/k0kubun/go-ansi"
	progressbar "github.com/schollz/progressbar/v3"
	"github.com/urfave/cli/v2"
)

var versionCmd = &cli.Command{
	Name: "version",
	Action: func(context *cli.Context) error {
		fmt.Println("version cmd")
		bar := progressbar.NewOptions(1000,
			progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
			progressbar.OptionEnableColorCodes(true),
			// progressbar.OptionShowBytes(true),
			// progressbar.OptionSetWidth(15),
			progressbar.OptionSetDescription(""),
			progressbar.OptionSetTheme(progressbar.Theme{
				Saucer:        "[green]=[reset]",
				SaucerHead:    "[green]>[reset]",
				SaucerPadding: " ",
				BarStart:      "[",
				BarEnd:        "]",
			}))
		for i := 0; i < 1000; i++ {
			bar.Add(1)
			time.Sleep(5 * time.Millisecond)
		}
		return nil
	},
}
