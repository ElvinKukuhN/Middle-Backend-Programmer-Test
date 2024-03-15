package main

import (
	"github.com/ElvinKukuhN/Middle-Backend-Programmer-Test/routes"
	"log"
	"net/http"
	"os"
)

func main() {
	routes.SetupRoutes()

	port := ":8080" // Port default
	if p := os.Getenv("PORT"); p != "" {
		port = ":" + p
	}

	log.Printf("Server starting on port %s...", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
