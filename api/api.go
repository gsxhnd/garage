package api

import (
	"fmt"
	"net/http"
	"os/exec"
	"runtime"
	"time"
)

func OpenBrowser(url string) {
	var commands = map[string]string{
		"windows": "start",
		"darwin":  "open",
		"linux":   "xdg-open",
	}
	run, ok := commands[runtime.GOOS]
	if !ok {
		fmt.Printf("don't know how to open things on %s platform\n", runtime.GOOS)
	}

	cmd := exec.Command(run, url)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}

func Run(port string) error {
	srv := &http.Server{
		Handler:      routes(),
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	err := srv.ListenAndServe()
	return err
}
