package handlers

import (
	"io"
	"net/http"
)

func SendErrorResponse(w http.ResponseWriter, httpSc int, msg string) {
	w.WriteHeader(httpSc)
	io.WriteString(w, msg)
}
