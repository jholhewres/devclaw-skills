---
name: bookmarks
description: "Save, organize, and search bookmarks and references"
---
# Bookmarks

Use **write_file**, **read_file**, **edit_file**, and **bash** to manage bookmarks.

## Save Bookmark
Append to `~/.devclaw/bookmarks.md`:

```markdown
## [Title](URL)
**Tags:** #tag1 #tag2
**Added:** YYYY-MM-DD
> Optional description or note
```

## Read Bookmarks
Use **read_file** on `~/.devclaw/bookmarks.md`

## Search
```bash
grep -i "TERM" ~/.devclaw/bookmarks.md
grep "#tag" ~/.devclaw/bookmarks.md
```

## Organize by Category
Alternative structure with separate files:
- `~/.devclaw/bookmarks/development.md`
- `~/.devclaw/bookmarks/articles.md`
- `~/.devclaw/bookmarks/tools.md`
- `~/.devclaw/bookmarks/reference.md`

## List Tags
```bash
grep -oh '#[a-zA-Z0-9_-]*' ~/.devclaw/bookmarks.md | sort | uniq -c | sort -rn
```

## Tips
- Use web_fetch to grab page title automatically when user shares a URL
- Group related bookmarks with markdown headers
- Use tags for cross-cutting categorization
- Combine with summarize skill to add summaries of saved articles
