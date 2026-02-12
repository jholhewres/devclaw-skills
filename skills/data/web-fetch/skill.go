// Package webfetch implements URL fetching and content extraction for GoClaw.
// Fetches web pages and extracts readable text, stripping HTML boilerplate.
package webfetch

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"
)

// WebFetchSkill fetches URLs and extracts readable content.
type WebFetchSkill struct {
	timeout      time.Duration
	maxBodyBytes int64
	userAgent    string
	client       *http.Client
}

// New creates a new WebFetchSkill instance.
func New(config map[string]any) (*WebFetchSkill, error) {
	s := &WebFetchSkill{
		timeout:      30 * time.Second,
		maxBodyBytes: 5 * 1024 * 1024, // 5MB
		userAgent:    "GoClaw/1.0",
	}
	s.applyConfig(config)
	s.client = &http.Client{Timeout: s.timeout}
	return s, nil
}

// Init initializes the skill.
func (s *WebFetchSkill) Init(_ context.Context, config map[string]any) error {
	s.applyConfig(config)
	s.client = &http.Client{Timeout: s.timeout}
	return nil
}

// Execute fetches a URL from raw input.
func (s *WebFetchSkill) Execute(ctx context.Context, input string) (string, error) {
	result, err := s.Fetch(ctx, map[string]any{"url": strings.TrimSpace(input)})
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", result), nil
}

// Shutdown releases resources.
func (s *WebFetchSkill) Shutdown() error {
	s.client.CloseIdleConnections()
	return nil
}

// ---------- Tools ----------

// Fetch fetches a URL and returns its content.
func (s *WebFetchSkill) Fetch(ctx context.Context, args map[string]any) (any, error) {
	rawURL, _ := args["url"].(string)
	if rawURL == "" {
		return nil, fmt.Errorf("url is required")
	}
	format, _ := args["format"].(string)
	if format == "" {
		format = "text"
	}

	body, contentType, err := s.doGet(ctx, rawURL)
	if err != nil {
		return nil, err
	}

	// If the response is not HTML, return raw content.
	if !strings.Contains(contentType, "text/html") {
		return body, nil
	}

	switch format {
	case "raw", "html":
		return body, nil
	case "markdown":
		return htmlToMarkdown(body), nil
	default: // "text"
		return htmlToText(body), nil
	}
}

// FetchHeaders performs a HEAD request and returns headers.
func (s *WebFetchSkill) FetchHeaders(ctx context.Context, args map[string]any) (any, error) {
	rawURL, _ := args["url"].(string)
	if rawURL == "" {
		return nil, fmt.Errorf("url is required")
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodHead, rawURL, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}
	req.Header.Set("User-Agent", s.userAgent)

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HEAD %s: %w", rawURL, err)
	}
	defer resp.Body.Close()

	headers := make(map[string]string)
	headers["status"] = resp.Status
	for k, v := range resp.Header {
		headers[k] = strings.Join(v, ", ")
	}
	return headers, nil
}

// FetchJSON fetches a JSON endpoint and parses the response.
func (s *WebFetchSkill) FetchJSON(ctx context.Context, args map[string]any) (any, error) {
	rawURL, _ := args["url"].(string)
	if rawURL == "" {
		return nil, fmt.Errorf("url is required")
	}

	body, _, err := s.doGet(ctx, rawURL)
	if err != nil {
		return nil, err
	}

	var result any
	if err := json.Unmarshal([]byte(body), &result); err != nil {
		return nil, fmt.Errorf("invalid JSON response: %w", err)
	}

	// Simple jq-like path extraction.
	if jqPath, ok := args["jq"].(string); ok && jqPath != "" {
		return extractJQPath(result, jqPath), nil
	}
	return result, nil
}

// ---------- HTTP ----------

func (s *WebFetchSkill) doGet(ctx context.Context, rawURL string) (string, string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, rawURL, nil)
	if err != nil {
		return "", "", fmt.Errorf("creating request: %w", err)
	}
	req.Header.Set("User-Agent", s.userAgent)
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/json,text/plain;q=0.9,*/*;q=0.8")

	resp, err := s.client.Do(req)
	if err != nil {
		return "", "", fmt.Errorf("GET %s: %w", rawURL, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return "", "", fmt.Errorf("GET %s returned %d", rawURL, resp.StatusCode)
	}

	lr := io.LimitReader(resp.Body, s.maxBodyBytes)
	bodyBytes, err := io.ReadAll(lr)
	if err != nil {
		return "", "", fmt.Errorf("reading body: %w", err)
	}

	contentType := resp.Header.Get("Content-Type")
	return string(bodyBytes), contentType, nil
}

// ---------- HTML Processing ----------

// htmlToText strips HTML tags and extracts readable text content.
// This is a lightweight implementation; for production use consider
// a proper readability library.
func htmlToText(html string) string {
	// Remove script and style blocks.
	reScript := regexp.MustCompile(`(?is)<script[^>]*>.*?</script>`)
	html = reScript.ReplaceAllString(html, "")
	reStyle := regexp.MustCompile(`(?is)<style[^>]*>.*?</style>`)
	html = reStyle.ReplaceAllString(html, "")
	reNav := regexp.MustCompile(`(?is)<nav[^>]*>.*?</nav>`)
	html = reNav.ReplaceAllString(html, "")
	reFooter := regexp.MustCompile(`(?is)<footer[^>]*>.*?</footer>`)
	html = reFooter.ReplaceAllString(html, "")

	// Convert block elements to newlines.
	reBlock := regexp.MustCompile(`(?i)</(p|div|h[1-6]|li|tr|br|hr)[^>]*>`)
	html = reBlock.ReplaceAllString(html, "\n")
	reBR := regexp.MustCompile(`(?i)<br\s*/?>`)
	html = reBR.ReplaceAllString(html, "\n")

	// Strip remaining tags.
	reTags := regexp.MustCompile(`<[^>]+>`)
	text := reTags.ReplaceAllString(html, "")

	// Decode common HTML entities.
	text = strings.ReplaceAll(text, "&amp;", "&")
	text = strings.ReplaceAll(text, "&lt;", "<")
	text = strings.ReplaceAll(text, "&gt;", ">")
	text = strings.ReplaceAll(text, "&quot;", "\"")
	text = strings.ReplaceAll(text, "&#39;", "'")
	text = strings.ReplaceAll(text, "&nbsp;", " ")

	// Collapse whitespace.
	reSpaces := regexp.MustCompile(`[ \t]+`)
	text = reSpaces.ReplaceAllString(text, " ")
	reNewlines := regexp.MustCompile(`\n{3,}`)
	text = reNewlines.ReplaceAllString(text, "\n\n")

	return strings.TrimSpace(text)
}

// htmlToMarkdown converts HTML to a basic markdown representation.
func htmlToMarkdown(html string) string {
	// Remove script/style/nav/footer.
	reScript := regexp.MustCompile(`(?is)<script[^>]*>.*?</script>`)
	html = reScript.ReplaceAllString(html, "")
	reStyle := regexp.MustCompile(`(?is)<style[^>]*>.*?</style>`)
	html = reStyle.ReplaceAllString(html, "")

	// Convert headings.
	for i := 1; i <= 6; i++ {
		prefix := strings.Repeat("#", i)
		re := regexp.MustCompile(fmt.Sprintf(`(?is)<h%d[^>]*>(.*?)</h%d>`, i, i))
		html = re.ReplaceAllString(html, "\n"+prefix+" $1\n")
	}

	// Convert links.
	reLink := regexp.MustCompile(`(?is)<a[^>]*href="([^"]*)"[^>]*>(.*?)</a>`)
	html = reLink.ReplaceAllString(html, "[$2]($1)")

	// Convert bold/italic.
	reBold := regexp.MustCompile(`(?is)<(strong|b)[^>]*>(.*?)</(strong|b)>`)
	html = reBold.ReplaceAllString(html, "**$2**")
	reItalic := regexp.MustCompile(`(?is)<(em|i)[^>]*>(.*?)</(em|i)>`)
	html = reItalic.ReplaceAllString(html, "*$2*")

	// Convert code.
	reCode := regexp.MustCompile(`(?is)<code[^>]*>(.*?)</code>`)
	html = reCode.ReplaceAllString(html, "`$1`")

	// Convert lists.
	reLI := regexp.MustCompile(`(?is)<li[^>]*>(.*?)</li>`)
	html = reLI.ReplaceAllString(html, "- $1\n")

	// Convert paragraphs and divs.
	reP := regexp.MustCompile(`(?is)</(p|div)>`)
	html = reP.ReplaceAllString(html, "\n\n")
	reBR := regexp.MustCompile(`(?i)<br\s*/?>`)
	html = reBR.ReplaceAllString(html, "\n")

	// Strip remaining tags.
	reTags := regexp.MustCompile(`<[^>]+>`)
	md := reTags.ReplaceAllString(html, "")

	// Decode entities.
	md = strings.ReplaceAll(md, "&amp;", "&")
	md = strings.ReplaceAll(md, "&lt;", "<")
	md = strings.ReplaceAll(md, "&gt;", ">")
	md = strings.ReplaceAll(md, "&quot;", "\"")
	md = strings.ReplaceAll(md, "&#39;", "'")
	md = strings.ReplaceAll(md, "&nbsp;", " ")

	// Clean up whitespace.
	reSpaces := regexp.MustCompile(`[ \t]+`)
	md = reSpaces.ReplaceAllString(md, " ")
	reNewlines := regexp.MustCompile(`\n{3,}`)
	md = reNewlines.ReplaceAllString(md, "\n\n")

	return strings.TrimSpace(md)
}

// ---------- Helpers ----------

// extractJQPath does basic dot-notation path extraction on parsed JSON.
// Supports: .field, .field.nested, .field[0]
func extractJQPath(data any, path string) any {
	path = strings.TrimPrefix(path, ".")
	if path == "" {
		return data
	}

	parts := strings.SplitN(path, ".", 2)
	key := parts[0]

	// Handle array index: field[0]
	if idx := strings.Index(key, "["); idx >= 0 {
		field := key[:idx]
		indexStr := strings.TrimSuffix(key[idx+1:], "]")

		if field != "" {
			if m, ok := data.(map[string]any); ok {
				data = m[field]
			} else {
				return nil
			}
		}

		if arr, ok := data.([]any); ok {
			var i int
			if _, err := fmt.Sscanf(indexStr, "%d", &i); err == nil && i < len(arr) {
				data = arr[i]
			} else {
				return nil
			}
		} else {
			return nil
		}
	} else {
		if m, ok := data.(map[string]any); ok {
			data = m[key]
		} else {
			return nil
		}
	}

	if len(parts) == 2 {
		return extractJQPath(data, parts[1])
	}
	return data
}

func (s *WebFetchSkill) applyConfig(config map[string]any) {
	if t, ok := toInt(config["timeout_seconds"]); ok {
		s.timeout = time.Duration(t) * time.Second
	}
	if m, ok := toInt(config["max_body_bytes"]); ok {
		s.maxBodyBytes = int64(m)
	}
	if ua, ok := config["user_agent"].(string); ok {
		s.userAgent = ua
	}
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
