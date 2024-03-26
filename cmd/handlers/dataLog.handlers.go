package handlers

import (
	"log"
	"net/http"
)

func FormatDataLogHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("YES")
}
