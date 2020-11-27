package main

import (
	"garage/command"
	"os"
)

func main() {
	err := command.App.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
