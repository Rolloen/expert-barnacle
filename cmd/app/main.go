package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load(".env")

	routerMux := http.NewServeMux()
	Routes(routerMux)

	server := &http.Server{
		Addr:    os.Getenv("HOST") + ":" + os.Getenv("PORT"),
		Handler: routerMux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalln("oups", err)
	}
}
