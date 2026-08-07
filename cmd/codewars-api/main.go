package main

import (
	"codewars-pretty-stats/internal/router"
	"log"
	"net/http"
	"time"
)

func main() {
	log.Println("Initializing...")

	var rout = router.New()

	server := &http.Server{
		Addr:         ":4322",
		Handler:      rout,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	log.Println("Server listening to localhost:4322...")
	log.Fatal(server.ListenAndServe())
}
