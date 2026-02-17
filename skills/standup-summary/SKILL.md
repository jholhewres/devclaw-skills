---
name: standup-summary
description: "Generate daily standup summary from git + issues"
---
# Standup Summary

Use **bash** (git + gh) to generate daily standup summaries.

## Collect Yesterday's Activity
```bash
# Commits in last 24h
git log --since="24 hours ago" --pretty=format:"- %s (%an, %ar)" --no-merges

# PRs merged yesterday
gh pr list -R OWNER/REPO --state merged --search "merged:>=$(date -d 'yesterday' +%Y-%m-%d)" --json number,title,author --limit 20

# Issues closed yesterday
gh issue list -R OWNER/REPO --state closed --search "closed:>=$(date -d 'yesterday' +%Y-%m-%d)" --json number,title,assignees --limit 20

# Issues in progress
gh issue list -R OWNER/REPO --label "in-progress" --json number,title,assignees --limit 20
```

## Format
Structure as:
- **Done**: what was completed
- **In Progress**: what's being worked on
- **Blocked**: any blockers

## Tips
- Run via cron_add as a daily job
- Post to Slack/Discord channel automatically
- Group by team member when in team mode
