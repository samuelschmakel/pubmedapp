package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"strings"

	"github.com/samuelschmakel/pubmedapp/backend/config"
	"github.com/samuelschmakel/pubmedapp/backend/processing"
)

type Handler struct {
	Cfg *config.ApiConfig
}

type ArticleInfo struct {
	Title string `json:"title"`
	Abstract string `json:"abstract"`
	URL string `json:"url"`
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

	eSearchResult, err := processing.FetchESearchResult(h.Cfg.HttpClient)
	if err != nil {
		http.Error(w, "Error retrieving query UIDs from PubMed: "+err.Error(), http.StatusBadGateway)
		return
	}

	articleSet, err := processing.FetchEFetchResult(h.Cfg.HttpClient, eSearchResult.ESearchResult.IDlist)
	if err != nil {
		http.Error(w, "Error retreiving abstracts from UIDs: "+err.Error(), http.StatusBadGateway)
	}

	if articleSet == nil {
		fmt.Println("the result from efetch was nil")
		return
	}

	var articles []ArticleInfo

	for _, a := range articleSet.PubmedArticles {
		pmid := a.MedlineCitation.PMID
		title := a.MedlineCitation.Article.ArticleTitle
		abstract := strings.Join(a.MedlineCitation.Article.Abstract.AbstractText, " ")
		url := fmt.Sprintf("https://pubmed.ncbi.nlm.nih.gov/%s/", pmid)
	
		articles = append(articles, ArticleInfo{
			Title:    title,
			Abstract: abstract,
			URL:      url,
		})
	}

    w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(articles)
	if err != nil {
		http.Error(w, "Failed to encode articles", http.StatusInternalServerError)
		return
	}
    //fmt.Fprint(w, `{"message": "Handling data, hello from Go backend!"}`)
}