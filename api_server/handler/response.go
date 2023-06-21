package handler

import (
	"encoding/json"
	"github.com/just-lym/video_server/api_server/defs"
	"io"
	"net/http"
)

func SendErrorResponse(w http.ResponseWriter, response defs.ErrResponse) {
	w.WriteHeader(response.HttpSC)
	body, _ := json.Marshal(&response.Error)
	io.WriteString(w, string(body))
}

func SendNormalResponse(w http.ResponseWriter, resp string, sc int) {
	w.WriteHeader(sc)
	io.WriteString(w, resp)
}
