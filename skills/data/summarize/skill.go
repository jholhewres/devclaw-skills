// Package summarize implements the summarize skill for GoClaw.
// Wraps the summarize CLI (https://summarize.sh) to extract and
// summarize content from URLs, videos, podcasts, and files.
package summarize

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
)

// SummarizeSkill wraps the summarize CLI for content extraction.
type SummarizeSkill struct {
	defaultModel    string
	defaultLanguage string
}

// New creates a new SummarizeSkill instance.
func New(config map[string]any) (*SummarizeSkill, error) {
	s := &SummarizeSkill{
		defaultLanguage: "en",
	}
	s.applyConfig(config)
	return s, nil
}

// Init initializes the skill. Checks if summarize is available.
func (s *SummarizeSkill) Init(_ context.Context, config map[string]any) error {
	s.applyConfig(config)
	if _, err := exec.LookPath("summarize"); err != nil {
		return fmt.Errorf("summarize CLI not found in PATH: %w", err)
	}
	return nil
}

// Execute summarizes from raw input (URL or text).
func (s *SummarizeSkill) Execute(ctx context.Context, input string) (string, error) {
	input = strings.TrimSpace(input)
	if strings.HasPrefix(input, "http://") || strings.HasPrefix(input, "https://") {
		result, err := s.SummarizeURL(ctx, map[string]any{"url": input})
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%v", result), nil
	}
	result, err := s.SummarizeText(ctx, map[string]any{"text": input})
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", result), nil
}

// Shutdown releases resources.
func (s *SummarizeSkill) Shutdown() error {
	return nil
}

// ---------- Tools ----------

// SummarizeURL summarizes content from a URL.
func (s *SummarizeSkill) SummarizeURL(ctx context.Context, args map[string]any) (any, error) {
	url, _ := args["url"].(string)
	if url == "" {
		return nil, fmt.Errorf("url is required")
	}

	cmdArgs := []string{url}
	cmdArgs = s.appendCommonFlags(cmdArgs, args)

	return s.run(ctx, cmdArgs...)
}

// TranscribeURL extracts transcript from a YouTube/podcast URL.
func (s *SummarizeSkill) TranscribeURL(ctx context.Context, args map[string]any) (any, error) {
	url, _ := args["url"].(string)
	if url == "" {
		return nil, fmt.Errorf("url is required")
	}

	cmdArgs := []string{url, "--transcript"}
	if lang, ok := args["language"].(string); ok && lang != "" {
		cmdArgs = append(cmdArgs, "--language", lang)
	}

	return s.run(ctx, cmdArgs...)
}

// SummarizeFile summarizes a local file.
func (s *SummarizeSkill) SummarizeFile(ctx context.Context, args map[string]any) (any, error) {
	path, _ := args["path"].(string)
	if path == "" {
		return nil, fmt.Errorf("path is required")
	}

	cmdArgs := []string{"--file", path}
	cmdArgs = s.appendCommonFlags(cmdArgs, args)

	return s.run(ctx, cmdArgs...)
}

// SummarizeText summarizes provided text content.
func (s *SummarizeSkill) SummarizeText(ctx context.Context, args map[string]any) (any, error) {
	text, _ := args["text"].(string)
	if text == "" {
		return nil, fmt.Errorf("text is required")
	}

	// Pass text via stdin using echo | summarize --stdin
	cmdArgs := []string{"--stdin"}
	cmdArgs = s.appendCommonFlags(cmdArgs, args)

	cmd := exec.CommandContext(ctx, "summarize", cmdArgs...)
	cmd.Stdin = strings.NewReader(text)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("summarize: %w\n%s", err, string(out))
	}
	return strings.TrimSpace(string(out)), nil
}

// ---------- Helpers ----------

func (s *SummarizeSkill) appendCommonFlags(args []string, params map[string]any) []string {
	lang := s.defaultLanguage
	if l, ok := params["language"].(string); ok && l != "" {
		lang = l
	}
	if lang != "" {
		args = append(args, "--language", lang)
	}

	if maxWords, ok := toInt(params["max_words"]); ok && maxWords > 0 {
		args = append(args, "--max-words", fmt.Sprintf("%d", maxWords))
	}

	if s.defaultModel != "" {
		args = append(args, "--model", s.defaultModel)
	}

	return args
}

func (s *SummarizeSkill) run(ctx context.Context, args ...string) (string, error) {
	cmd := exec.CommandContext(ctx, "summarize", args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("summarize %s: %w\n%s", strings.Join(args, " "), err, string(out))
	}
	return strings.TrimSpace(string(out)), nil
}

func (s *SummarizeSkill) applyConfig(config map[string]any) {
	if m, ok := config["default_model"].(string); ok {
		s.defaultModel = m
	}
	if l, ok := config["default_language"].(string); ok {
		s.defaultLanguage = l
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
	}
	return 0, false
}
