// Package github implements the GitHub skill for GoClaw.
// Wraps the gh CLI to provide full GitHub integration:
// issues, PRs, CI runs, releases, API, and search.
package github

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

// GitHubSkill provides GitHub integration via the gh CLI.
type GitHubSkill struct {
	defaultOwner string
}

// New creates a new GitHubSkill instance.
func New(config map[string]any) (*GitHubSkill, error) {
	s := &GitHubSkill{}
	if owner, ok := config["default_owner"].(string); ok {
		s.defaultOwner = owner
	}
	return s, nil
}

// Init initializes the skill. Checks if gh is available.
func (s *GitHubSkill) Init(_ context.Context, config map[string]any) error {
	if owner, ok := config["default_owner"].(string); ok {
		s.defaultOwner = owner
	}
	if _, err := exec.LookPath("gh"); err != nil {
		return fmt.Errorf("gh CLI not found in PATH: %w", err)
	}
	return nil
}

// Execute runs a GitHub operation based on the input.
func (s *GitHubSkill) Execute(ctx context.Context, input string) (string, error) {
	// The agent should use tools directly; Execute is a fallback
	// for free-form gh commands.
	return s.runGH(ctx, strings.Fields(input)...)
}

// Shutdown releases resources.
func (s *GitHubSkill) Shutdown() error {
	return nil
}

// ---------- Issues ----------

// ListIssues lists issues in a repository.
func (s *GitHubSkill) ListIssues(ctx context.Context, args map[string]any) (any, error) {
	repo := s.resolveRepo(args)
	if repo == "" {
		return nil, fmt.Errorf("repo is required")
	}
	cmdArgs := []string{"issue", "list", "--repo", repo, "--json",
		"number,title,state,labels,assignees,createdAt,updatedAt"}

	if state, ok := args["state"].(string); ok && state != "" {
		cmdArgs = append(cmdArgs, "--state", state)
	}
	if labels, ok := args["labels"].(string); ok && labels != "" {
		cmdArgs = append(cmdArgs, "--label", labels)
	}
	if limit, ok := toInt(args["limit"]); ok {
		cmdArgs = append(cmdArgs, "--limit", fmt.Sprintf("%d", limit))
	}
	return s.runGHJSON(ctx, cmdArgs...)
}

// GetIssue gets details of a specific issue.
func (s *GitHubSkill) GetIssue(ctx context.Context, args map[string]any) (any, error) {
	repo := s.resolveRepo(args)
	number, ok := toInt(args["number"])
	if !ok {
		return nil, fmt.Errorf("number is required")
	}
	return s.runGHJSON(ctx, "issue", "view", fmt.Sprintf("%d", number),
		"--repo", repo, "--json",
		"number,title,body,state,labels,assignees,comments,createdAt,closedAt")
}

// CreateIssue creates a new issue.
func (s *GitHubSkill) CreateIssue(ctx context.Context, args map[string]any) (any, error) {
	repo := s.resolveRepo(args)
	title, _ := args["title"].(string)
	if title == "" {
		return nil, fmt.Errorf("title is required")
	}

	cmdArgs := []string{"issue", "create", "--repo", repo, "--title", title}
	if body, ok := args["body"].(string); ok && body != "" {
		cmdArgs = append(cmdArgs, "--body", body)
	}
	if labels, ok := args["labels"].(string); ok && labels != "" {
		cmdArgs = append(cmdArgs, "--label", labels)
	}
	if assignees, ok := args["assignees"].(string); ok && assignees != "" {
		cmdArgs = append(cmdArgs, "--assignee", assignees)
	}
	return s.runGH(ctx, cmdArgs...)
}

// CloseIssue closes an issue by number.
func (s *GitHubSkill) CloseIssue(ctx context.Context, args map[string]any) (any, error) {
	repo := s.resolveRepo(args)
	number, ok := toInt(args["number"])
	if !ok {
		return nil, fmt.Errorf("number is required")
	}
	return s.runGH(ctx, "issue", "close", fmt.Sprintf("%d", number), "--repo", repo)
}

// CommentIssue adds a comment to an issue.
func (s *GitHubSkill) CommentIssue(ctx context.Context, args map[string]any) (any, error) {
	repo := s.resolveRepo(args)
	number, ok := toInt(args["number"])
	if !ok {
		return nil, fmt.Errorf("number is required")
	}
	body, _ := args["body"].(string)
	if body == "" {
		return nil, fmt.Errorf("body is required")
	}
	return s.runGH(ctx, "issue", "comment", fmt.Sprintf("%d", number),
		"--repo", repo, "--body", body)
}

// ---------- Pull Requests ----------

// ListPRs lists pull requests.
func (s *GitHubSkill) ListPRs(ctx context.Context, args map[string]any) (any, error) {
	repo := s.resolveRepo(args)
	cmdArgs := []string{"pr", "list", "--repo", repo, "--json",
		"number,title,state,author,headRefName,baseRefName,createdAt,updatedAt"}

	if state, ok := args["state"].(string); ok && state != "" {
		cmdArgs = append(cmdArgs, "--state", state)
	}
	if limit, ok := toInt(args["limit"]); ok {
		cmdArgs = append(cmdArgs, "--limit", fmt.Sprintf("%d", limit))
	}
	return s.runGHJSON(ctx, cmdArgs...)
}

// GetPR gets details of a specific pull request.
func (s *GitHubSkill) GetPR(ctx context.Context, args map[string]any) (any, error) {
	repo := s.resolveRepo(args)
	number, ok := toInt(args["number"])
	if !ok {
		return nil, fmt.Errorf("number is required")
	}
	return s.runGHJSON(ctx, "pr", "view", fmt.Sprintf("%d", number),
		"--repo", repo, "--json",
		"number,title,body,state,author,reviews,statusCheckRollup,headRefName,baseRefName,additions,deletions,changedFiles")
}

// PRDiff gets the diff of a pull request.
func (s *GitHubSkill) PRDiff(ctx context.Context, args map[string]any) (any, error) {
	repo := s.resolveRepo(args)
	number, ok := toInt(args["number"])
	if !ok {
		return nil, fmt.Errorf("number is required")
	}
	return s.runGH(ctx, "pr", "diff", fmt.Sprintf("%d", number), "--repo", repo)
}

// PRChecks checks CI status on a pull request.
func (s *GitHubSkill) PRChecks(ctx context.Context, args map[string]any) (any, error) {
	repo := s.resolveRepo(args)
	number, ok := toInt(args["number"])
	if !ok {
		return nil, fmt.Errorf("number is required")
	}
	return s.runGH(ctx, "pr", "checks", fmt.Sprintf("%d", number), "--repo", repo)
}

// PRMerge merges a pull request.
func (s *GitHubSkill) PRMerge(ctx context.Context, args map[string]any) (any, error) {
	repo := s.resolveRepo(args)
	number, ok := toInt(args["number"])
	if !ok {
		return nil, fmt.Errorf("number is required")
	}
	method := "merge"
	if m, ok := args["method"].(string); ok && m != "" {
		method = m
	}

	flag := "--merge"
	switch method {
	case "squash":
		flag = "--squash"
	case "rebase":
		flag = "--rebase"
	}
	return s.runGH(ctx, "pr", "merge", fmt.Sprintf("%d", number), "--repo", repo, flag)
}

// ---------- CI/CD ----------

// ListRuns lists recent workflow runs.
func (s *GitHubSkill) ListRuns(ctx context.Context, args map[string]any) (any, error) {
	repo := s.resolveRepo(args)
	cmdArgs := []string{"run", "list", "--repo", repo, "--json",
		"databaseId,displayTitle,status,conclusion,event,createdAt,headBranch"}

	if workflow, ok := args["workflow"].(string); ok && workflow != "" {
		cmdArgs = append(cmdArgs, "--workflow", workflow)
	}
	if limit, ok := toInt(args["limit"]); ok {
		cmdArgs = append(cmdArgs, "--limit", fmt.Sprintf("%d", limit))
	}
	return s.runGHJSON(ctx, cmdArgs...)
}

// ViewRun views details of a workflow run.
func (s *GitHubSkill) ViewRun(ctx context.Context, args map[string]any) (any, error) {
	repo := s.resolveRepo(args)
	runID, _ := args["run_id"].(string)
	if runID == "" {
		return nil, fmt.Errorf("run_id is required")
	}

	cmdArgs := []string{"run", "view", runID, "--repo", repo}
	if showLog, ok := args["log"].(bool); ok && showLog {
		cmdArgs = append(cmdArgs, "--log")
	}
	return s.runGH(ctx, cmdArgs...)
}

// RerunWorkflow re-runs a failed workflow.
func (s *GitHubSkill) RerunWorkflow(ctx context.Context, args map[string]any) (any, error) {
	repo := s.resolveRepo(args)
	runID, _ := args["run_id"].(string)
	if runID == "" {
		return nil, fmt.Errorf("run_id is required")
	}
	return s.runGH(ctx, "run", "rerun", runID, "--repo", repo, "--failed")
}

// ---------- Releases ----------

// ListReleases lists releases.
func (s *GitHubSkill) ListReleases(ctx context.Context, args map[string]any) (any, error) {
	repo := s.resolveRepo(args)
	cmdArgs := []string{"release", "list", "--repo", repo}
	if limit, ok := toInt(args["limit"]); ok {
		cmdArgs = append(cmdArgs, "--limit", fmt.Sprintf("%d", limit))
	}
	return s.runGH(ctx, cmdArgs...)
}

// CreateRelease creates a new release.
func (s *GitHubSkill) CreateRelease(ctx context.Context, args map[string]any) (any, error) {
	repo := s.resolveRepo(args)
	tag, _ := args["tag"].(string)
	title, _ := args["title"].(string)
	if tag == "" || title == "" {
		return nil, fmt.Errorf("tag and title are required")
	}

	cmdArgs := []string{"release", "create", tag, "--repo", repo, "--title", title}
	if notes, ok := args["notes"].(string); ok && notes != "" {
		cmdArgs = append(cmdArgs, "--notes", notes)
	}
	if draft, ok := args["draft"].(bool); ok && draft {
		cmdArgs = append(cmdArgs, "--draft")
	}
	if prerelease, ok := args["prerelease"].(bool); ok && prerelease {
		cmdArgs = append(cmdArgs, "--prerelease")
	}
	return s.runGH(ctx, cmdArgs...)
}

// ---------- API & Search ----------

// API makes a raw GitHub API request.
func (s *GitHubSkill) API(ctx context.Context, args map[string]any) (any, error) {
	endpoint, _ := args["endpoint"].(string)
	if endpoint == "" {
		return nil, fmt.Errorf("endpoint is required")
	}

	cmdArgs := []string{"api", endpoint}
	if method, ok := args["method"].(string); ok && method != "" && method != "GET" {
		cmdArgs = append(cmdArgs, "--method", method)
	}
	if body, ok := args["body"].(string); ok && body != "" {
		cmdArgs = append(cmdArgs, "--input", "-")
	}
	return s.runGHJSON(ctx, cmdArgs...)
}

// SearchRepos searches GitHub repositories.
func (s *GitHubSkill) SearchRepos(ctx context.Context, args map[string]any) (any, error) {
	query, _ := args["query"].(string)
	if query == "" {
		return nil, fmt.Errorf("query is required")
	}
	cmdArgs := []string{"search", "repos", query, "--json",
		"fullName,description,stargazersCount,language,updatedAt"}
	if limit, ok := toInt(args["limit"]); ok {
		cmdArgs = append(cmdArgs, "--limit", fmt.Sprintf("%d", limit))
	}
	return s.runGHJSON(ctx, cmdArgs...)
}

// SearchCode searches code across GitHub.
func (s *GitHubSkill) SearchCode(ctx context.Context, args map[string]any) (any, error) {
	query, _ := args["query"].(string)
	if query == "" {
		return nil, fmt.Errorf("query is required")
	}
	cmdArgs := []string{"search", "code", query}
	if repo, ok := args["repo"].(string); ok && repo != "" {
		cmdArgs = append(cmdArgs, "--repo", repo)
	}
	if limit, ok := toInt(args["limit"]); ok {
		cmdArgs = append(cmdArgs, "--limit", fmt.Sprintf("%d", limit))
	}
	return s.runGH(ctx, cmdArgs...)
}

// ---------- Helpers ----------

// resolveRepo returns the repo from args or builds it from default_owner.
func (s *GitHubSkill) resolveRepo(args map[string]any) string {
	if repo, ok := args["repo"].(string); ok && repo != "" {
		return repo
	}
	return ""
}

// runGH executes a gh command and returns stdout as string.
func (s *GitHubSkill) runGH(ctx context.Context, args ...string) (string, error) {
	cmd := exec.CommandContext(ctx, "gh", args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("gh %s: %w\n%s", strings.Join(args, " "), err, string(out))
	}
	return strings.TrimSpace(string(out)), nil
}

// runGHJSON executes a gh command and parses the JSON output.
func (s *GitHubSkill) runGHJSON(ctx context.Context, args ...string) (any, error) {
	output, err := s.runGH(ctx, args...)
	if err != nil {
		return nil, err
	}
	var result any
	if err := json.Unmarshal([]byte(output), &result); err != nil {
		// If JSON parsing fails, return raw output.
		return output, nil
	}
	return result, nil
}

// toInt converts an any value to int.
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
