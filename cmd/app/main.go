package main

import (
	"fmt"
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
		Addr:    ":" + os.Getenv("PORT"),
		Handler: routerMux,
	}
	fmt.Println("LAUNCHING SERVER on " + server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalln("Error has occured when launching the server", err)
	}
}
