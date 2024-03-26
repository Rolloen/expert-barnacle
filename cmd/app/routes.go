package app

import (
	"net/http"
	"techTest/cmd/handlers"
)

func Routes(router *http.ServeMux) {
	router.HandleFunc("POST /test", handlers.FormatDataLogHandler)
}
