// Package websearch implements web search for GoClaw.
// Supports Brave Search API and SearXNG as providers.
package websearch

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// SearchResult represents a single search result.
type SearchResult struct {
	Title   string `json:"title"`
	URL     string `json:"url"`
	Snippet string `json:"snippet"`
	Age     string `json:"age,omitempty"`
}

// WebSearchSkill provides web search capabilities.
type WebSearchSkill struct {
	provider    string // "brave" or "searxng"
	braveAPIKey string
	searxngURL  string
	count       int
	safeSearch  string
	client      *http.Client
}

// New creates a new WebSearchSkill instance.
func New(config map[string]any) (*WebSearchSkill, error) {
	s := &WebSearchSkill{
		provider:   "brave",
		count:      5,
		safeSearch: "moderate",
		client: &http.Client{
			Timeout: 15 * time.Second,
		},
	}
	if err := s.applyConfig(config); err != nil {
		return nil, err
	}
	return s, nil
}

// Init initializes the skill with the provided configuration.
func (s *WebSearchSkill) Init(_ context.Context, config map[string]any) error {
	return s.applyConfig(config)
}

// Execute runs a search with the raw input string.
func (s *WebSearchSkill) Execute(ctx context.Context, input string) (string, error) {
	results, err := s.Search(ctx, map[string]any{"query": input})
	if err != nil {
		return "", err
	}
	b, _ := json.MarshalIndent(results, "", "  ")
	return string(b), nil
}

// Shutdown releases resources.
func (s *WebSearchSkill) Shutdown() error {
	s.client.CloseIdleConnections()
	return nil
}

// ---------- Tools ----------

// Search performs a web search.
func (s *WebSearchSkill) Search(ctx context.Context, args map[string]any) (any, error) {
	query, _ := args["query"].(string)
	if query == "" {
		return nil, fmt.Errorf("query is required")
	}

	count := s.count
	if c, ok := toInt(args["count"]); ok && c > 0 {
		count = c
	}
	freshness, _ := args["freshness"].(string)

	switch s.provider {
	case "searxng":
		return s.searchSearXNG(ctx, query, count, freshness, "general")
	default:
		return s.searchBrave(ctx, query, count, freshness, "search")
	}
}

// SearchNews searches for news articles.
func (s *WebSearchSkill) SearchNews(ctx context.Context, args map[string]any) (any, error) {
	query, _ := args["query"].(string)
	if query == "" {
		return nil, fmt.Errorf("query is required")
	}

	count := s.count
	if c, ok := toInt(args["count"]); ok && c > 0 {
		count = c
	}
	freshness, _ := args["freshness"].(string)
	if freshness == "" {
		freshness = "week"
	}

	switch s.provider {
	case "searxng":
		return s.searchSearXNG(ctx, query, count, freshness, "news")
	default:
		return s.searchBrave(ctx, query, count, freshness, "news")
	}
}

// ---------- Brave Search ----------

func (s *WebSearchSkill) searchBrave(ctx context.Context, query string, count int, freshness, searchType string) ([]SearchResult, error) {
	if s.braveAPIKey == "" {
		return nil, fmt.Errorf("BRAVE_API_KEY not configured")
	}

	endpoint := "https://api.search.brave.com/res/v1/web/search"
	if searchType == "news" {
		endpoint = "https://api.search.brave.com/res/v1/news/search"
	}

	params := url.Values{
		"q":     {query},
		"count": {fmt.Sprintf("%d", count)},
	}
	if freshness != "" {
		// Brave uses: pd (past day), pw (past week), pm (past month), py (past year)
		switch freshness {
		case "day":
			params.Set("freshness", "pd")
		case "week":
			params.Set("freshness", "pw")
		case "month":
			params.Set("freshness", "pm")
		case "year":
			params.Set("freshness", "py")
		}
	}
	if s.safeSearch != "" {
		params.Set("safesearch", s.safeSearch)
	}

	reqURL := endpoint + "?" + params.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Encoding", "gzip")
	req.Header.Set("X-Subscription-Token", s.braveAPIKey)

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("brave search: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("brave search returned %d: %s", resp.StatusCode, string(body))
	}

	if searchType == "news" {
		return parseBraveNews(body)
	}
	return parseBraveWeb(body)
}

func parseBraveWeb(body []byte) ([]SearchResult, error) {
	var resp struct {
		Web struct {
			Results []struct {
				Title       string `json:"title"`
				URL         string `json:"url"`
				Description string `json:"description"`
				Age         string `json:"age"`
			} `json:"results"`
		} `json:"web"`
	}
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("parsing brave response: %w", err)
	}

	results := make([]SearchResult, 0, len(resp.Web.Results))
	for _, r := range resp.Web.Results {
		results = append(results, SearchResult{
			Title:   r.Title,
			URL:     r.URL,
			Snippet: r.Description,
			Age:     r.Age,
		})
	}
	return results, nil
}

func parseBraveNews(body []byte) ([]SearchResult, error) {
	var resp struct {
		Results []struct {
			Title       string `json:"title"`
			URL         string `json:"url"`
			Description string `json:"description"`
			Age         string `json:"age"`
		} `json:"results"`
	}
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("parsing brave news response: %w", err)
	}

	results := make([]SearchResult, 0, len(resp.Results))
	for _, r := range resp.Results {
		results = append(results, SearchResult{
			Title:   r.Title,
			URL:     r.URL,
			Snippet: r.Description,
			Age:     r.Age,
		})
	}
	return results, nil
}

// ---------- SearXNG ----------

func (s *WebSearchSkill) searchSearXNG(ctx context.Context, query string, count int, freshness, category string) ([]SearchResult, error) {
	if s.searxngURL == "" {
		return nil, fmt.Errorf("SEARXNG_URL not configured")
	}

	params := url.Values{
		"q":          {query},
		"format":     {"json"},
		"categories": {category},
	}
	if freshness != "" {
		params.Set("time_range", freshness)
	}

	reqURL := strings.TrimRight(s.searxngURL, "/") + "/search?" + params.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("searxng search: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("searxng returned %d: %s", resp.StatusCode, string(body))
	}

	var searxResp struct {
		Results []struct {
			Title   string `json:"title"`
			URL     string `json:"url"`
			Content string `json:"content"`
		} `json:"results"`
	}
	if err := json.Unmarshal(body, &searxResp); err != nil {
		return nil, fmt.Errorf("parsing searxng response: %w", err)
	}

	results := make([]SearchResult, 0, min(count, len(searxResp.Results)))
	for i, r := range searxResp.Results {
		if i >= count {
			break
		}
		results = append(results, SearchResult{
			Title:   r.Title,
			URL:     r.URL,
			Snippet: r.Content,
		})
	}
	return results, nil
}

// ---------- Helpers ----------

func (s *WebSearchSkill) applyConfig(config map[string]any) error {
	if p, ok := config["provider"].(string); ok {
		s.provider = p
	}
	if key, ok := config["brave_api_key"].(string); ok {
		s.braveAPIKey = key
	}
	if u, ok := config["searxng_url"].(string); ok {
		s.searxngURL = u
	}
	if c, ok := toInt(config["default_count"]); ok {
		s.count = c
	}
	if ss, ok := config["safe_search"].(string); ok {
		s.safeSearch = ss
	}
	return nil
}

func toInt(v any) (int, bool) {
	switch n := v.(type) {
	case int:
		return n, true
	case float64:
		return int(n), true
	case int64:
		return int(n), true
	case json.Number:
		i, err := n.Int64()
		return int(i), err == nil
	}
	return 0, false
}
