package handlers

import (
	"fmt"
	"net/http"

	"github.com/samuelschmakel/pubmedapp/backend/config"
)

type Handler struct {
	Cfg *config.ApiConfig
}

func NewHandler(cfg *config.ApiConfig) *Handler {
	return &Handler{Cfg: cfg}
}

func (h *Handler) HandleSubmit(w http.ResponseWriter, req *http.Request) {

    // CORS headers
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    if req.Method == "OPTIONS" {
        // Stop here for preflight requests
        w.WriteHeader(http.StatusOK)
        return
    }

    // TO DO: Use these parameters in helper function to 
    query := req.URL.Query().Get("query")
    context := req.URL.Query().Get("context")
    fmt.Printf("query: %s, context: %s\n", query, context)

    w.Header().Set("Content-Type", "application/json")
    fmt.Fprint(w, `{"message": "Handling data, hello from Go backend!"}`)
}