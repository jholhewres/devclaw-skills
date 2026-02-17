---
name: sprint-report
description: "Generate sprint reports from git + issue tracker"
---
# Sprint Report

Use **bash** (git + gh/jira CLI) to generate sprint reports.

## Collect Sprint Data
```bash
# Commits in sprint period
git log --since="2 weeks ago" --pretty=format:"%h %s (%an)" --no-merges | head -50
git shortlog --since="2 weeks ago" -sn

# PRs merged in sprint
gh pr list -R OWNER/REPO --state merged --search "merged:>=$(date -d '2 weeks ago' +%Y-%m-%d)" --json number,title,author --limit 50

# Issues closed in sprint
gh issue list -R OWNER/REPO --state closed --search "closed:>=$(date -d '2 weeks ago' +%Y-%m-%d)" --json number,title,labels --limit 50

# Open issues (carry-over)
gh issue list -R OWNER/REPO --state open --label "sprint-current" --json number,title,assignees --limit 30
```

## Report Sections
- **Velocity**: issues completed vs planned
- **Completed**: features, fixes, improvements
- **Carry-over**: unfinished items
- **Blockers**: items that were blocked
- **Metrics**: PRs merged, commits, contributors

## Tips
- Compare with previous sprint for trends
- Highlight scope changes (added/removed mid-sprint)
- Use memory_save to track velocity history
