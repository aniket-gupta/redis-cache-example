package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aniket-gupta/redis-cache-example/cmd/handler"
)

func main() {
	// We create a simple httpserver
	server := http.Server{
		Addr:    fmt.Sprint(":8000"),
		Handler: handler.NewHandler(),
	}

	// Run server
	log.Printf("Starting HTTP Server. Listening at %q", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Printf("%v", err)
	} else {
		log.Println("Server closed ! ")
	}

}
