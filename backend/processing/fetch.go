package processing

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

	
type ESearchResult struct {
	ESearchResult struct {
		IDlist         []string `json:"idlist"`
		Error string `json:"ERROR"` // only present when the API errors
	} `json:"esearchresult"`
}

func FetchPapers() (*ESearchResult, error) {
	fmt.Println("running fetchPapers()")
	url := os.Getenv("ESEARCH_URL")
	// Change to make this not hardcoded
	url += "&term=cancer+immunotherapy&retmax=2&retmode=json&email=samuel.schmakel@gmail.com"
	fmt.Printf("url passed into Get request: %s\n", url)
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response data: %w", err)
	}

	var result ESearchResult
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if result.ESearchResult.Error != "" {
		return nil, fmt.Errorf("API Error: %s", result.ESearchResult.Error)
	}

	return &result, nil
}

func FetchAbstracts(url string) (string, error) {
	// url := os.Getenv("EFETCH_URL")
	// Change to make this not hardcoded
	// url += "&id=40601938, 40601888&rettype=abstract&retmode=text&email=samuel.schmakel@gmail.com"

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}

	// Set headers similar to what a browser or Bruno might send
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; MyApp/1.0)")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")

	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error making HTTP request: %w", err)
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	return string(bodyBytes), nil
}