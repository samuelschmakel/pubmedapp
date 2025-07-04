package handlers

import (
	"fmt"
	"net/http"

	"github.com/samuelschmakel/pubmedapp/backend/config"
	"github.com/samuelschmakel/pubmedapp/backend/processing"
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

    // TODO: Use these parameters in helper function to query dataset
    query := req.URL.Query().Get("query")
    context := req.URL.Query().Get("context")

	// Verify that the query exists:
	if query == "" {
		http.Error(w, "Missing requried query field", http.StatusBadRequest)
		return
	}

    fmt.Printf("query: %s, context: %s\n", query, context)
	eSearchResult, err := processing.FetchPapers()
	if err != nil {
		fmt.Printf("error returned from FetchPapers(): %s", error.Error(err))
		http.Error(w, "Error retrieving query results from Pubmed", http.StatusInternalServerError)
		return
	}

	IDlist := eSearchResult.ESearchResult.IDlist
	if len(IDlist) == 0 {
		http.Error(w, "No articles found for that query", http.StatusBadRequest)
		return
	}

	for _, v := range IDlist {
		fmt.Printf("ID: %s\n", v)
	}

    w.Header().Set("Content-Type", "application/json")
    fmt.Fprint(w, `{"message": "Handling data, hello from Go backend!"}`)
}