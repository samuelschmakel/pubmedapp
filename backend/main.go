package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

type apiconfig struct {
    fileserverHits atomic.Int32
    platform string
    secretkey string
}

func main() {
    http.HandleFunc("/api/data", handleData)

    // Run server on port 8080
    fmt.Println("Starting server on port 8080...")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        panic(err)
    }

    /*
    const filepathRoot = "."
    port := os.Getenv("PORT")
    fmt.Printf("port: %s", port)

    godotenv.Load()
    dbURL := os.Getenv("DB_URL")
    if dbURL == "" {
        log.Fatal("PLATFORM must be set")
    }

    mux := http.NewServeMux()

    mux.Handle("GET /api/test", handlerPass)
    */
}

func handleData(w http.ResponseWriter, r *http.Request) {

    // CORS headers
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    if r.Method == "OPTIONS" {
        // Stop here for preflight requests
        w.WriteHeader(http.StatusOK)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    fmt.Fprint(w, `{"message": "Handling data, hello from Go backend!"}`)
}