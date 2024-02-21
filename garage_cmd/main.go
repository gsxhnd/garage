package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"

	_ "github.com/glebarez/go-sqlite"
)

var RootCmd = cli.NewApp()

func init() {
	RootCmd.HideVersion = true
	RootCmd.Usage = "命令行工具"
	RootCmd.Flags = []cli.Flag{}
	RootCmd.Commands = []*cli.Command{
		crawlCmd,
		ffmpegBatchCmd,
		versionCmd,
		serverCmd,
	}
	// RootCmd.CommandNotFound = func(ctx *cli.Context, s string) {
	// 	fmt.Println(s)
	// }
}

func main() {
	if db, err := sql.Open("sqlite", "./testdb/billfish.db"); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("db open successfully")
		row := db.QueryRow("select name from bf_file limit 1")
		fmt.Println(row.Err())
		var data interface{}
		fmt.Println(row.Scan(&data))
		fmt.Println(data)
	}

	err := RootCmd.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
