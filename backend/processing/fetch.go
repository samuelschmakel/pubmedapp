package processing

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
)

	
type ESearchResult struct {
	ESearchResult struct {
		IDlist         []string `json:"idlist"`
		Error string `json:"ERROR"` // only present when the API errors
	} `json:"esearchresult"`
}

type PubMedArticleSet struct {
	PubmedArticles []PubmedArticle `xml:"PubmedArticle"`
}

type PubmedArticle struct {
    MedlineCitation struct {
        PMID    string `xml:"PMID"`
        Article struct {
            ArticleTitle string `xml:"ArticleTitle"`
            Abstract     struct {
                AbstractText []string `xml:"AbstractText"`
            } `xml:"Abstract"`
        } `xml:"Article"`
    } `xml:"MedlineCitation"`
}

func FetchESearchResult(client *http.Client, query, num_articles string) (*ESearchResult, error) {
	url := os.Getenv("ESEARCH_URL")

	// Compiles regex to match one or more whitespace characters
	re := regexp.MustCompile(`\s+`)
	// Replaces all whitespace with "+" and starts the string with "&"
	cleaned := "&term=" + re.ReplaceAllString(strings.TrimSpace(query), "+")

	cleaned += "&retmax=" + num_articles
	// Change to make this not hardcoded
	url += cleaned + "&retmode=json&email=samuel.schmakel@gmail.com"
	//url += "&term=cancer+immunotherapy&retmax=2&retmode=json&email=samuel.schmakel@gmail.com"

	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response data: %w", err)
	}

	var result ESearchResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if result.ESearchResult.Error != "" {
		return nil, fmt.Errorf("NCBI Esearch Error: %s", result.ESearchResult.Error)
	}

	return &result, nil
}

func FetchEFetchResult(client *http.Client, idSlice []string) (*PubMedArticleSet, error) {
	url := os.Getenv("EFETCH_URL")
	// Change to make this not hardcoded
	ids := "&id=" + strings.Join(idSlice, ",")

	url += ids + "&rettype=abstract&retmode=xml&email=samuel.schmakel@gmail.com"
	//url += "&id=40601938,40601888&rettype=abstract&retmode=xml&email=samuel.schmakel@gmail.com"
	fmt.Printf("The url being used: %s", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	
	// Set headers similar to what a browser or Bruno might send
	// req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; MyApp/1.0)")
	// req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making HTTP request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var articles PubMedArticleSet
	if err := xml.Unmarshal(body, &articles); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &articles, nil
}