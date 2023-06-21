package handler

import (
	"github.com/just-lym/video_server/api_server/defs"
	"github.com/just-lym/video_server/api_server/session"
	"net/http"
)

var (
	HEADER_FIELD_SESSION = "X-Session-Id"
	HEADER_FIELD_UNAME   = "X-User-Name"
)

func ValidUserSession(r *http.Request) bool {
	sid := r.Header.Get(HEADER_FIELD_SESSION)
	if len(sid) == 0 {
		return false
	}
	userName, b := session.IsSessionExpired(sid)
	if b {
		return false
	}
	r.Header.Add(HEADER_FIELD_UNAME, userName)
	return true
}

func ValidateUser(w http.ResponseWriter, r *http.Request) bool {
	uname := r.Header.Get(HEADER_FIELD_UNAME)
	if len(uname) == 0 {
		SendErrorResponse(w, defs.ErrorNotAuthUser)
		return false
	}
	return true
}
