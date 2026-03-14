package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type SearchRequest struct {
	Search       string `json:"search"`
	QueryType    string `json:"queryType,omitempty"`
	SearchMode   string `json:"searchMode,omitempty"`
	SearchFields string `json:"searchFields,omitempty"`
	Select       string `json:"select,omitempty"`
	Top          int    `json:"top,omitempty"`
	Count        bool   `json:"count,omitempty"`
}

type SearchResponse struct {
	ODataContext string                   `json:"@odata.context"`
	ODataCount   *int                     `json:"@odata.count,omitempty"`
	Value        []map[string]interface{} `json:"value"`
}

func SearchAI(serviceName, indexName, apiKey, searchText string) (*SearchResponse, error) {
	url := fmt.Sprintf("https://%s.search.windows.net/indexes/%s/docs/search?api-version=2025-09-01",
		serviceName, indexName)

	// Build request with recommended defaults
	reqBody := SearchRequest{
		Search:     searchText,
		QueryType:  "simple",
		SearchMode: "any",
		Count:      true,
		Top:        10,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("api-key", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	// Check for HTTP errors
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search failed with status %d: %s", resp.StatusCode, string(body))
	}

	var searchResp SearchResponse
	if err := json.Unmarshal(body, &searchResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &searchResp, nil
}

func main() {
	serviceName := os.Getenv("SEARCH_SERVICE_NAME")
	indexName := os.Getenv("SEARCH_INDEX_NAME")
	apiKey := os.Getenv("SEARCH_API_KEY")

	if serviceName == "" || indexName == "" || apiKey == "" {
		fmt.Fprintln(os.Stderr, "Error: SEARCH_SERVICE_NAME, SEARCH_INDEX_NAME, and SEARCH_API_KEY must be set")
		os.Exit(1)
	}

	var searchText string
	if len(os.Args) > 1 {
		searchText = os.Args[1]
	} else {
		// Default to wildcard search to test connection
		searchText = "*"
		fmt.Println("No search query provided, using '*' to return all documents")
	}

	result, err := SearchAI(serviceName, indexName, apiKey, searchText)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if result.ODataCount != nil {
		fmt.Printf("Total documents matching: %d\n", *result.ODataCount)
	}
	fmt.Printf("Results returned: %d\n\n", len(result.Value))

	if len(result.Value) == 0 {
		fmt.Println("No results found. Troubleshooting tips:")
		fmt.Println("1. Verify your index has documents (try search='*')")
		fmt.Println("2. Check if the search fields contain the terms you're searching for")
		fmt.Println("3. Try a simpler search query (single word)")
		fmt.Println("4. Verify your API key has query permissions")
	} else {
		for i, doc := range result.Value {
			fmt.Printf("Result %d:\n", i+1)
			jsonDoc, _ := json.MarshalIndent(doc, "", "  ")
			fmt.Println(string(jsonDoc))
			fmt.Println()
		}
	}
}
