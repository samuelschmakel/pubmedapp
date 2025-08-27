package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/samuelschmakel/pubmedapp/backend/config"
	"github.com/samuelschmakel/pubmedapp/backend/handlers"
)

func main() {   
    err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found, using environment variables")
    }
    cfg := config.CreateConfig()
    h := handlers.NewHandler(cfg)

    http.Handle("/api/", http.HandlerFunc(h.HandleSubmit))       // API routes
    http.Handle("/", http.FileServer(http.Dir("./frontend/")))   // Frontend

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"  // default for local dev
    }

    // Run server on port (8080 by default)
    fmt.Printf("Starting server on port %s...\n", port)
    err = http.ListenAndServe(":"+port, nil)
    if err != nil {
        panic(err)
    }
}