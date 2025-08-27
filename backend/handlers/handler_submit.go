package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"slices"

	"strings"

	"github.com/samuelschmakel/pubmedapp/backend/config"
	"github.com/samuelschmakel/pubmedapp/backend/processing"
)

// TODO: Call the Python HTTP API

type Handler struct {
	Cfg *config.ApiConfig
}

type ArticleInfo struct {
	Title string `json:"title"`
	Abstract string `json:"abstract"`
	URL string `json:"url"`
	Score float64 `json:"score"`
}

type DataFrameRow struct {
	Abstract string `json:"abstract"`
	SimilarityScore float64 `json:"similarity_score"`
}

type ArticlesSimilarity struct {
	Articles string `json:"articles"`
	Similarity float64 `json:"similarity"`
}

type PythonAPIInput struct {
	ArticleInfo []ArticleInfo `json:"articleInfo"`
	Context []string `json:"context"`
}

/*
type Response struct {
	Articles []ArticleInfo `json:"articles"`
	DataFrame []DataFrameRow `json:"data_frame,omitempty"`
}
	*/

func NewHandler(cfg *config.ApiConfig) *Handler {
	return &Handler{Cfg: cfg}
}

func (h *Handler) HandleSubmit(w http.ResponseWriter, req *http.Request) {
	 defer func() {
        if err := recover(); err != nil {
            log.Printf("Handler panicked: %v", err)
            http.Error(w, "Internal server error", http.StatusInternalServerError)
        }
    }()

    // CORS headers
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    if req.Method == "OPTIONS" {
        // Stop here for preflight requests
        w.WriteHeader(http.StatusOK)
        return
    }

    query := req.URL.Query().Get("query")
	numArticles := req.URL.Query().Get("num_articles")
	contextString := req.URL.Query().Get("context")
	var context []string

	if contextString != "" {
		context = strings.Split(contextString, ",")
		// Clean up whitespace if needed
		for i, item := range context {
			context[i] = strings.TrimSpace(item)
		}
	} else {
		context = []string{} // empty slice if no context provided
	}

	// Verify that the query exists:
	if query == "" {
		http.Error(w, "Missing requried query field", http.StatusBadRequest)
		return
	}

    fmt.Printf("query: %s, context: %s\n", query, context)

	eSearchResult, err := processing.FetchESearchResult(h.Cfg.HttpClient, query, numArticles)
	if err != nil {
		http.Error(w, "Error retrieving query UIDs from PubMed: "+err.Error(), http.StatusBadGateway)
		return
	}

	for _, v := range eSearchResult.ESearchResult.IDlist {
		fmt.Printf("ID: %s\n", v)
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

	var pd_frame []DataFrameRow

	// If context was provided, call the Python API to generate the scored pandas data frame
	// Then, add it to articles
	if len(context) != 0 {
		input := PythonAPIInput{
			ArticleInfo: articles,
			Context: context,
		}

		pd_frame, err = h.callPythonAPI("/process-list", input)
		fmt.Println("data frame: ", pd_frame)
		if err != nil {
			fmt.Println("error calling local Python API")
			fmt.Printf("error: %v\n", err)
			return
		}
		addArticleInfoScoreField(&articles, pd_frame)
		sortArticles(&articles)
	}

    w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(articles)
	if err != nil {
		http.Error(w, "Failed to encode articles", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) callPythonAPI(endpoint string, data PythonAPIInput) ([]DataFrameRow, error) {
	fmt.Println("calling Python API")
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal data: %w", err)
	}

	// TODO: use this to debug Python API. Also, test curl commands to test Python API
	fmt.Printf("JSON payload being sent: %s\n", string(jsonData))

	fmt.Printf("url sent to Python: %s\n", h.Cfg.PythonBaseURL+endpoint)
	resp, err := h.Cfg.PythonClient.Post(
		h.Cfg.PythonBaseURL+endpoint,
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return nil, fmt.Errorf("error calling Python client: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("python service returned status: %d", resp.StatusCode)
	}

	var result []DataFrameRow
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	return result, nil
}

func addArticleInfoScoreField(a *[]ArticleInfo, pd_frame []DataFrameRow) {
	for i := range *a {
		(*a)[i].Score = pd_frame[i].SimilarityScore
	}
}

func sortArticles(a *[]ArticleInfo) {
	slices.SortFunc(*a, func(a, b ArticleInfo) int {
		if a.Score > b.Score {
			return -1 // negative for descending
		}
		if a.Score < b.Score {
			return 1
		}
		return 0
	})
}