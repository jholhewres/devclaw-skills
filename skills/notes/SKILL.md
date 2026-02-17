---
name: notes
version: 0.1.0
author: devclaw
description: "Quick notes, lists, and ideas — stored locally as markdown files"
category: builtin
tags: [notes, memo, lists, ideas, journal, quick-capture]
---
# Notes

You can save, search, read, and manage notes for the user. Notes are stored as markdown files in the `~/.devclaw/notes/` directory.

## Creating notes

```bash
# Ensure directory exists
mkdir -p ~/.devclaw/notes

# Quick note (timestamped)
cat > ~/.devclaw/notes/$(date +%Y%m%d-%H%M%S)-note.md << 'EOF'
# Quick note

Content of the note here.

Tags: #idea #work
EOF

# Named note
cat > ~/.devclaw/notes/shopping-list.md << 'EOF'
# Shopping List

- [ ] Milk
- [ ] Bread
- [ ] Eggs
- [ ] Fruits
EOF

# Append to existing note
echo "- [ ] Coffee" >> ~/.devclaw/notes/shopping-list.md

# Journal entry (daily)
cat >> ~/.devclaw/notes/journal-$(date +%Y-%m-%d).md << EOF

## $(date +%H:%M) — Entry

What happened today...

EOF
```

## Reading notes

```bash
# List all notes
ls -lt ~/.devclaw/notes/ | head -20

# Read a specific note
cat ~/.devclaw/notes/shopping-list.md

# Search across all notes
grep -rl "SEARCH_TERM" ~/.devclaw/notes/ 2>/dev/null
grep -n "SEARCH_TERM" ~/.devclaw/notes/*.md 2>/dev/null

# Recent notes (last 7 days)
find ~/.devclaw/notes/ -name "*.md" -mtime -7 -exec ls -lt {} +
```

## Editing notes

```bash
# Replace content
cat > ~/.devclaw/notes/FILENAME.md << 'EOF'
# Updated title

New content here.
EOF

# Mark todo as done (replace "- [ ]" with "- [x]")
sed -i 's/- \[ \] Milk/- [x] Milk/' ~/.devclaw/notes/shopping-list.md
```

## Deleting notes

```bash
# Delete a note (always confirm first!)
rm ~/.devclaw/notes/FILENAME.md

# Archive old notes
mkdir -p ~/.devclaw/notes/archive
mv ~/.devclaw/notes/old-note.md ~/.devclaw/notes/archive/
```

## Note types

| Type | Naming convention | Example |
|------|------------------|---------|
| Quick note | `YYYYMMDD-HHMMSS-note.md` | `20260214-153000-note.md` |
| Named note | `descriptive-name.md` | `shopping-list.md` |
| Daily journal | `journal-YYYY-MM-DD.md` | `journal-2026-02-14.md` |
| Project note | `project-NAME.md` | `project-devclaw.md` |
| Todo list | `todo-NAME.md` | `todo-this-week.md` |

## Tips

- Always confirm before deleting notes.
- Use descriptive names so notes are easy to find later.
- For todo lists, use `- [ ]` / `- [x]` markdown checkboxes.
- Use tags at the bottom of notes (e.g., `Tags: #work #urgent`) for easier searching.
- When the user says "note this", "save this", "anota isso", or "salva isso", create a quick note with timestamp.
- Read the note back to the user after creating it for confirmation.
- For large collections, suggest organizing by folders/categories.

## Triggers

note, save this, note this, anota isso, salva isso, create a note, my notes,
remember this, lembrar disso, shopping list, lista de compras, todo list, journal, diário
