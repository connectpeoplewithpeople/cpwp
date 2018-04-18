package api

import (
	"fmt"
	"common"
	"net/http"
)

func GetStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, "{\"version\":\"v1\",\"status\":0,\"port\":%v}", common.HttpPort)
}