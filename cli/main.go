package main

import (
	"github.com/gsxhnd/garage/cli/command"
	"os"
)

func main() {
	err := command.RootCmd.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
