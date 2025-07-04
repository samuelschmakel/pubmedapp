package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"

	"github.com/joho/godotenv"
	"github.com/samuelschmakel/pubmedapp/backend/config"
	"github.com/samuelschmakel/pubmedapp/backend/handlers"
)

func main() {   
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    cfg := &config.ApiConfig{
        FileserverHits: atomic.Int32{},
    }

    h := handlers.NewHandler(cfg)

    http.HandleFunc("/api/data", h.HandleSubmit)

    // Run server on port 8080
    fmt.Println("Starting server on port 8080...")
    err = http.ListenAndServe(":8080", nil)
    if err != nil {
        panic(err)
    }
}