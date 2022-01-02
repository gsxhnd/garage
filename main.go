package main

import (
	"os"

	"github.com/gsxhnd/garage/command"
)

func main() {
	err := command.RootCmd.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
