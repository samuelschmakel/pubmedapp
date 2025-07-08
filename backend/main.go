package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync/atomic"
	"time"

	"github.com/joho/godotenv"
	"github.com/samuelschmakel/pubmedapp/backend/config"
	"github.com/samuelschmakel/pubmedapp/backend/handlers"
)

func main() {   
    err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found, using environment variables")
    }
    cfg := &config.ApiConfig{
        FileserverHits: atomic.Int32{},
        HttpClient: &http.Client{
            Timeout: 10 * time.Second,
        },
    }

    h := handlers.NewHandler(cfg)

    http.HandleFunc("/api/data", h.HandleSubmit)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    // Run server on port (8080 by default)
    fmt.Printf("Starting server on port %s...\n", port)
    err = http.ListenAndServe(":"+port, nil)
    if err != nil {
        panic(err)
    }
}