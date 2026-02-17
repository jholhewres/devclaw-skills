---
name: journal
description: "Daily journaling with structured entries and reflections"
---
# Journal

Use **write_file**, **read_file**, **edit_file**, and **memory_save** for daily journaling.

## Daily Entry
Use **write_file** to create `~/.devclaw/journal/YYYY-MM-DD.md`:

```markdown
# 2026-02-17

## How I'm feeling
(mood, energy level)

## What I accomplished
- ...

## What I learned
- ...

## What I'm grateful for
- ...

## Tomorrow's focus
- ...
```

## Append to Today
Use **edit_file** to add entries to today's journal throughout the day.

## Review
```bash
ls -lt ~/.devclaw/journal/ | head -14    # last 2 weeks
cat ~/.devclaw/journal/$(date +%Y-%m-%d).md 2>/dev/null   # today
```

## Weekly Reflection
```bash
# Read last 7 entries
for i in $(seq 0 6); do
  D=$(date -d "$i days ago" +%Y-%m-%d)
  [ -f ~/.devclaw/journal/$D.md ] && echo "=== $D ===" && head -20 ~/.devclaw/journal/$D.md
done
```

## Search
```bash
grep -rl "TERM" ~/.devclaw/journal/ --include="*.md"
```

## Tips
- Create journal directory on first use: mkdir -p ~/.devclaw/journal
- Save key insights with memory_save for long-term recall
- Don't force structure — adapt to user's style
- Ask reflective questions to prompt deeper entries
