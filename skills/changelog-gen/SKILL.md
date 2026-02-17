---
name: changelog-gen
description: "Generate CHANGELOG.md from conventional commits"
---
# Changelog Generator

Use the **bash** tool to generate changelogs from Git history.

## Generate from Conventional Commits
```bash
# List commits since last tag
git log $(git describe --tags --abbrev=0 2>/dev/null || echo "HEAD~50")..HEAD --pretty=format:"%s (%h)" | head -50

# Group by type
git log $(git describe --tags --abbrev=0 2>/dev/null || echo "HEAD~50")..HEAD --pretty=format:"%s" | grep -E "^feat" | head -20
git log $(git describe --tags --abbrev=0 2>/dev/null || echo "HEAD~50")..HEAD --pretty=format:"%s" | grep -E "^fix" | head -20
git log $(git describe --tags --abbrev=0 2>/dev/null || echo "HEAD~50")..HEAD --pretty=format:"%s" | grep -E "^(refactor|perf|docs|chore|ci|test)" | head -20
```

## Conventional Commit Types
| Prefix | Section |
|--------|---------|
| feat: | Features |
| fix: | Bug Fixes |
| perf: | Performance |
| refactor: | Refactoring |
| docs: | Documentation |
| chore: | Maintenance |

## Workflow
1. Read current CHANGELOG.md with **read_file**
2. Collect commits since last release with **bash**
3. Group and format entries
4. Prepend new section with **edit_file**

## Tips
- Use git tags to determine release boundaries
- Include breaking changes prominently (BREAKING CHANGE: or !)
- Link to issues/PRs when available
