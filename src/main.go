package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/glebarez/go-sqlite"
	"github.com/gsxhnd/garage/src/cmd"
)

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

	err := cmd.RootCmd.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
