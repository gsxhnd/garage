package main

import (
	"github.com/gsxhnd/garage/command"
	"os"
)

func main() {
	err := command.RootCmd.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
