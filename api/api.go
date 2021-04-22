package api

import (
	"net/http"
	"time"
)

func Run(port, imgDir string) error {
	srv := &http.Server{
		Handler:      routes(imgDir),
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	err := srv.ListenAndServe()
	return err
}
