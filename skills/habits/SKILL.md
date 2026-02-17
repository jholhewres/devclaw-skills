---
name: habits
description: "Track daily habits and streaks"
---
# Habit Tracker

Use **write_file**, **read_file**, **edit_file**, **memory_save**, and **bash** to track habits.

## Setup Habits
Use **write_file** to create `~/.devclaw/habits/config.md`:

```markdown
# My Habits
- exercise
- read 30min
- meditate
- journal
- no social media before noon
```

## Daily Check-in
Use **edit_file** to update `~/.devclaw/habits/YYYY-MM.md`:

```markdown
# Habits — February 2026

| Day | exercise | read | meditate | journal |
|-----|----------|------|----------|---------|
| 01  | x        | x    |          | x       |
| 02  | x        |      | x        | x       |
| 03  |          | x    | x        |         |
```

## Streaks
```bash
# Count current streak for a habit (consecutive 'x' from today backwards)
cat ~/.devclaw/habits/$(date +%Y-%m).md 2>/dev/null
```

## Monthly Summary
```bash
cat ~/.devclaw/habits/$(date +%Y-%m).md 2>/dev/null | grep "| [0-9]" | awk -F'|' '{
  for(i=3;i<=NF-1;i++) if($i ~ /x/) count[i]++; total++
} END {
  for(i in count) printf "Habit %d: %d/%d (%.0f%%)\n", i-2, count[i], total, count[i]*100/total
}'
```

## Tips
- Create dirs on first use: mkdir -p ~/.devclaw/habits
- Ask user about habits at start of day (combine with cron_add)
- Celebrate streaks and milestones
- Use memory_save for streak records
- Keep it simple — 3-5 habits max is sustainable
