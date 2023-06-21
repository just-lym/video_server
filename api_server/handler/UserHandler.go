package handler

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/just-lym/video_server/api_server/dbops"
	"github.com/just-lym/video_server/api_server/defs"
	"github.com/just-lym/video_server/api_server/session"
	"io"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	all, _ := io.ReadAll(r.Body)
	userCredential := &defs.UserCredential{}
	if err := json.Unmarshal(all, userCredential); err != nil {
		SendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}
	if err := dbops.AddCredential(userCredential.Username, userCredential.Pwd); err != nil {
		SendErrorResponse(w, defs.ErrorDBError)
		return
	}
	sessionId := session.GenerateNewSessionId(userCredential.Username)
	signedUp := &defs.SignedUp{
		Success:   true,
		SessionId: sessionId,
	}
	if res, err := json.Marshal(signedUp); err != nil {
		SendErrorResponse(w, defs.ErrorInternalFaults)
	} else {
		SendNormalResponse(w, string(res), 200)
	}
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("user_name")
	io.WriteString(w, uname)
}
