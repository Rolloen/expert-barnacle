package main

import (
	"net/http"
	"techTest/cmd/handlers"
)

func Routes(router *http.ServeMux) {
	router.HandleFunc("GET /analysis", handlers.GetDataLogHandler)
}
