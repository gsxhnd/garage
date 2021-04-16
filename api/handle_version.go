package api

import (
	"garage/utils"
	"net/http"
)

func GetVersion(w http.ResponseWriter, req *http.Request) {
	SendRes(w, nil, utils.GetVersionInfo())
}
