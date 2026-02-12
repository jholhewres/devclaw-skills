// Package gog implements Google Workspace integration for GoClaw.
// Wraps the gog CLI to provide access to Gmail, Calendar, and Drive.
package gog

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

// GogSkill provides Google Workspace integration via the gog CLI.
type GogSkill struct {
	defaultAccount string
	calendarID     string
}

// New creates a new GogSkill instance.
func New(config map[string]any) (*GogSkill, error) {
	s := &GogSkill{
		calendarID: "primary",
	}
	s.applyConfig(config)
	return s, nil
}

// Init initializes the skill. Checks if gog is available.
func (s *GogSkill) Init(_ context.Context, config map[string]any) error {
	s.applyConfig(config)
	if _, err := exec.LookPath("gog"); err != nil {
		return fmt.Errorf("gog CLI not found in PATH: %w", err)
	}
	return nil
}

// Execute runs a gog command from raw input.
func (s *GogSkill) Execute(ctx context.Context, input string) (string, error) {
	return s.run(ctx, strings.Fields(input)...)
}

// Shutdown releases resources.
func (s *GogSkill) Shutdown() error {
	return nil
}

// ---------- Gmail Tools ----------

// GmailList lists recent emails.
func (s *GogSkill) GmailList(ctx context.Context, args map[string]any) (any, error) {
	cmdArgs := []string{"gmail", "list"}
	if query, ok := args["query"].(string); ok && query != "" {
		cmdArgs = append(cmdArgs, "--query", query)
	}
	if limit, ok := toInt(args["limit"]); ok {
		cmdArgs = append(cmdArgs, "--limit", fmt.Sprintf("%d", limit))
	}
	cmdArgs = append(cmdArgs, "--json")
	return s.runJSON(ctx, cmdArgs...)
}

// GmailRead reads a specific email.
func (s *GogSkill) GmailRead(ctx context.Context, args map[string]any) (any, error) {
	id, _ := args["id"].(string)
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	return s.runJSON(ctx, "gmail", "read", id, "--json")
}

// GmailSend sends an email.
func (s *GogSkill) GmailSend(ctx context.Context, args map[string]any) (any, error) {
	to, _ := args["to"].(string)
	subject, _ := args["subject"].(string)
	body, _ := args["body"].(string)
	if to == "" || subject == "" || body == "" {
		return nil, fmt.Errorf("to, subject, and body are required")
	}

	cmdArgs := []string{"gmail", "send", "--to", to, "--subject", subject, "--body", body}
	if cc, ok := args["cc"].(string); ok && cc != "" {
		cmdArgs = append(cmdArgs, "--cc", cc)
	}
	return s.run(ctx, cmdArgs...)
}

// GmailReply replies to an email.
func (s *GogSkill) GmailReply(ctx context.Context, args map[string]any) (any, error) {
	id, _ := args["id"].(string)
	body, _ := args["body"].(string)
	if id == "" || body == "" {
		return nil, fmt.Errorf("id and body are required")
	}
	return s.run(ctx, "gmail", "reply", id, "--body", body)
}

// ---------- Calendar Tools ----------

// CalendarList lists upcoming events.
func (s *GogSkill) CalendarList(ctx context.Context, args map[string]any) (any, error) {
	cmdArgs := []string{"calendar", "list"}
	if days, ok := toInt(args["days"]); ok {
		cmdArgs = append(cmdArgs, "--days", fmt.Sprintf("%d", days))
	}
	cal := s.calendarID
	if c, ok := args["calendar"].(string); ok && c != "" {
		cal = c
	}
	cmdArgs = append(cmdArgs, "--calendar", cal, "--json")
	return s.runJSON(ctx, cmdArgs...)
}

// CalendarCreate creates a calendar event.
func (s *GogSkill) CalendarCreate(ctx context.Context, args map[string]any) (any, error) {
	title, _ := args["title"].(string)
	start, _ := args["start"].(string)
	if title == "" || start == "" {
		return nil, fmt.Errorf("title and start are required")
	}

	cmdArgs := []string{"calendar", "create", "--title", title, "--start", start}
	if dur, ok := args["duration"].(string); ok && dur != "" {
		cmdArgs = append(cmdArgs, "--duration", dur)
	}
	if desc, ok := args["description"].(string); ok && desc != "" {
		cmdArgs = append(cmdArgs, "--description", desc)
	}
	if loc, ok := args["location"].(string); ok && loc != "" {
		cmdArgs = append(cmdArgs, "--location", loc)
	}
	if attendees, ok := args["attendees"].(string); ok && attendees != "" {
		cmdArgs = append(cmdArgs, "--attendees", attendees)
	}
	return s.run(ctx, cmdArgs...)
}

// CalendarDelete deletes a calendar event.
func (s *GogSkill) CalendarDelete(ctx context.Context, args map[string]any) (any, error) {
	id, _ := args["id"].(string)
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	return s.run(ctx, "calendar", "delete", id)
}

// ---------- Drive Tools ----------

// DriveList lists files in Google Drive.
func (s *GogSkill) DriveList(ctx context.Context, args map[string]any) (any, error) {
	cmdArgs := []string{"drive", "list"}
	if query, ok := args["query"].(string); ok && query != "" {
		cmdArgs = append(cmdArgs, "--query", query)
	}
	if folder, ok := args["folder"].(string); ok && folder != "" {
		cmdArgs = append(cmdArgs, "--folder", folder)
	}
	if limit, ok := toInt(args["limit"]); ok {
		cmdArgs = append(cmdArgs, "--limit", fmt.Sprintf("%d", limit))
	}
	cmdArgs = append(cmdArgs, "--json")
	return s.runJSON(ctx, cmdArgs...)
}

// DriveDownload downloads a file from Google Drive.
func (s *GogSkill) DriveDownload(ctx context.Context, args map[string]any) (any, error) {
	id, _ := args["id"].(string)
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	cmdArgs := []string{"drive", "download", id}
	if output, ok := args["output"].(string); ok && output != "" {
		cmdArgs = append(cmdArgs, "--output", output)
	}
	return s.run(ctx, cmdArgs...)
}

// DriveUpload uploads a file to Google Drive.
func (s *GogSkill) DriveUpload(ctx context.Context, args map[string]any) (any, error) {
	path, _ := args["path"].(string)
	if path == "" {
		return nil, fmt.Errorf("path is required")
	}
	cmdArgs := []string{"drive", "upload", path}
	if folder, ok := args["folder"].(string); ok && folder != "" {
		cmdArgs = append(cmdArgs, "--folder", folder)
	}
	if name, ok := args["name"].(string); ok && name != "" {
		cmdArgs = append(cmdArgs, "--name", name)
	}
	return s.run(ctx, cmdArgs...)
}

// ---------- Helpers ----------

func (s *GogSkill) run(ctx context.Context, args ...string) (string, error) {
	cmd := exec.CommandContext(ctx, "gog", args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("gog %s: %w\n%s", strings.Join(args, " "), err, string(out))
	}
	return strings.TrimSpace(string(out)), nil
}

func (s *GogSkill) runJSON(ctx context.Context, args ...string) (any, error) {
	output, err := s.run(ctx, args...)
	if err != nil {
		return nil, err
	}
	var result any
	if err := json.Unmarshal([]byte(output), &result); err != nil {
		return output, nil
	}
	return result, nil
}

func (s *GogSkill) applyConfig(config map[string]any) {
	if acc, ok := config["default_account"].(string); ok {
		s.defaultAccount = acc
	}
	if cal, ok := config["calendar_id"].(string); ok {
		s.calendarID = cal
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
