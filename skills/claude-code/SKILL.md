---
name: claude-code
description: "Full-stack coding assistant powered by Claude Code CLI"
---
# Claude Code

Use the **bash** tool with the Claude Code CLI for advanced coding tasks.

## Requirements
```bash
npm install -g @anthropic-ai/claude-code
claude setup-token   # or claude login
```

## Usage
```bash
claude -p "fix the authentication bug in auth.ts" --allowedTools bash,read,write
claude -p "review this code for security issues" --permission-mode plan
```

## Tips
- Be specific in prompts
- For read-only analysis, use --permission-mode plan
- Check auth: claude status
