package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/just-lym/video_server/api_server/handler"
	"net/http"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", handler.CreateUser)
	router.GET("/user/:user_name", handler.Login)
	return router
}

type middleWareHandler struct {
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return &m
}

func (m *middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler.ValidUserSession(r)
	m.r.ServeHTTP(w, r)
}

func main() {
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r)
	err := http.ListenAndServe(":8080", mh)
	if err != nil {
		return
	}
}
