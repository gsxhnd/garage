package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	rootCmd := cli.NewApp()

	err := rootCmd.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
