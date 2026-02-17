---
name: release-notes
description: "Generate release notes from issues + commits"
---
# Release Notes Generator

Use **bash** (git + gh CLI) to generate release notes.

## Collect Data
```bash
# Commits between tags
git log v1.0.0..v1.1.0 --pretty=format:"- %s (%h)" --no-merges

# Closed issues since last release (GitHub)
gh issue list -R OWNER/REPO --state closed --search "closed:>2026-01-01" --limit 50 --json number,title

# Merged PRs since last release
gh pr list -R OWNER/REPO --state merged --search "merged:>2026-01-01" --limit 50 --json number,title,labels
```

## Format
Group by: Features, Bug Fixes, Improvements, Breaking Changes.
Include: PR numbers, issue links, contributor mentions.

## Publish
```bash
gh release create TAG -R OWNER/REPO --title "TITLE" --notes-file release-notes.md
```

## Tips
- Use git tags to bound the release
- Cross-reference issues mentioned in commits
- Highlight breaking changes at the top
