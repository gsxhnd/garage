package main

import (
	"github.com/gsxhnd/garage/cmd"
	"os"
)

func main() {
	err := cmd.App.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
