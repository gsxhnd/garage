package main

import (
	"os"

	"github.com/gsxhnd/garage/src/cmd"
)

func main() {
	err := cmd.RootCmd.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
