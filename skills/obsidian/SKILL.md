---
name: obsidian
description: "Manage Obsidian vaults: notes, links, tags, daily notes, search"
---
# Obsidian

Use **read_file**, **write_file**, **edit_file**, and **bash** to manage Obsidian vaults.

## Locate Vault
```bash
# Common vault locations
ls -d ~/Documents/Obsidian* ~/Obsidian* ~/.obsidian* 2>/dev/null
# Or find .obsidian folders
find ~/ -maxdepth 4 -name ".obsidian" -type d 2>/dev/null | head -5
```

## Daily Notes
```bash
# Create/open today's daily note
VAULT="$HOME/Documents/ObsidianVault"
TODAY=$(date +%Y-%m-%d)
cat "$VAULT/Daily/$TODAY.md" 2>/dev/null || echo "# $TODAY" > "$VAULT/Daily/$TODAY.md"
```
Use **write_file** to create: `<vault>/Daily/YYYY-MM-DD.md`

## Search Notes
```bash
VAULT="$HOME/Documents/ObsidianVault"
grep -rl "TERM" "$VAULT" --include="*.md" | head -20
grep -rl "#tag" "$VAULT" --include="*.md" | head -20
```

## Create Note
Use **write_file** with path `<vault>/<folder>/<name>.md`:
- Use `[[wikilinks]]` for internal links
- Use `#tags` for categorization
- Use YAML frontmatter for metadata

## List & Organize
```bash
VAULT="$HOME/Documents/ObsidianVault"
find "$VAULT" -name "*.md" -not -path "*/.obsidian/*" | wc -l              # total notes
find "$VAULT" -name "*.md" -not -path "*/.obsidian/*" -newer "$VAULT/Daily/$(date -d '7 days ago' +%Y-%m-%d).md" | head -20  # recent
ls -lt "$VAULT"/**/*.md 2>/dev/null | head -20                              # recently modified
grep -roh '\[\[.*\]\]' "$VAULT" --include="*.md" | sort | uniq -c | sort -rn | head -20  # most linked
grep -roh '#[a-zA-Z0-9_-]*' "$VAULT" --include="*.md" | sort | uniq -c | sort -rn | head -20  # top tags
```

## Templates
Use **write_file** to create templates in `<vault>/Templates/`:
- Meeting: frontmatter + attendees + agenda + action items
- Project: goals + tasks + references
- Fleeting: quick capture with timestamp

## Tips
- Always ask user for vault path if unknown, or use memory_search to recall
- Use memory_save to remember vault path after first use
- Respect existing folder structure
- Use wikilinks [[Note Name]] for cross-references
- Daily notes go in Daily/ folder by convention
