package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/just-lym/video_server/stream_server/handlers"
	"github.com/just-lym/video_server/stream_server/limit"
	"net/http"
)

func registerHandlers() *httprouter.Router {
	router := httprouter.New()
	router.GET("/video/:vid-id", handlers.StreamHandler)
	router.POST("/upload/:vid-id", handlers.UploadHandler)
	return router
}

type middleWareHandler struct {
	r *httprouter.Router
	c *limit.ConnLimiter
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	m.c = limit.NewConnLimiter(10)
	return &m
}

func (m *middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if m.c.GetConn() {
		handlers.SendErrorResponse(w, http.StatusTooManyRequests, "too many request")
		return
	}
	m.r.ServeHTTP(w, r)
	defer m.c.ReleaseConn()
}

func main() {
	router := registerHandlers()
	middleWareHandler := NewMiddleWareHandler(router)
	err := http.ListenAndServe(":8765", middleWareHandler)
	if err != nil {
		return
	}
}
