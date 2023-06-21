package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func registerHandlers() *httprouter.Router {
	router := httprouter.New()
	router.GET("/video/:vid-id", nil)
	router.POST("/upload/:vid-id", nil)
	return router
}

func main() {
	handlers := registerHandlers()
	err := http.ListenAndServe(":8765", handlers)
	if err != nil {
		return
	}
}
