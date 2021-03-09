package api

import (
	"encoding/json"
	"net/http"
)

func SendRes(w http.ResponseWriter, err error, data interface{}) {
	w.Header().Set("Content-Type", "application/json;   charset=UTF-8")
	var res = struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}{}
	w.WriteHeader(http.StatusOK)

	if err != nil {
		res.Code = 10010
		res.Message = err.Error()
		_ = json.NewEncoder(w).Encode(res)
	} else {
		res.Code = 0
		res.Message = "ok"
		res.Data = data
		_ = json.NewEncoder(w).Encode(res)
	}
}
