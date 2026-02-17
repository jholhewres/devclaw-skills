---
name: time-tracker
description: "Track time spent on tasks and projects"
---
# Time Tracker

Use **memory_save**, **memory_search**, **write_file**, and **bash** to track time.

## Start Tracking
Use **memory_save**: "timetrack start: <task> at <ISO timestamp>"

## Stop Tracking
Use **memory_save**: "timetrack stop: <task> at <ISO timestamp>, duration: Xh Xm"

## Log File
Append entries to `~/.devclaw/timetrack/YYYY-MM-DD.csv`:
```
start,end,duration_min,project,task
2026-02-17T09:00:00,2026-02-17T10:30:00,90,devclaw,skills refactor
2026-02-17T10:45:00,2026-02-17T12:00:00,75,devclaw,frontend bugfix
```

## Reports
```bash
# Today's time
cat ~/.devclaw/timetrack/$(date +%Y-%m-%d).csv 2>/dev/null

# This week total by project
for i in $(seq 0 6); do
  D=$(date -d "$i days ago" +%Y-%m-%d)
  cat ~/.devclaw/timetrack/$D.csv 2>/dev/null
done | grep -v "^start" | awk -F, '{sum[$4]+=$3} END {for(p in sum) printf "%s: %.1fh\n", p, sum[p]/60}'

# This month
cat ~/.devclaw/timetrack/$(date +%Y-%m)-*.csv 2>/dev/null | grep -v "^start" | awk -F, '{sum[$4]+=$3} END {for(p in sum) printf "%s: %.1fh\n", p, sum[p]/60}'
```

## Tips
- Create dir on first use: mkdir -p ~/.devclaw/timetrack
- Ask what user is working on when they start a session
- Remind to stop timer when task changes
- Use memory_search "timetrack" to find recent entries
- Combine with pomodoro for focused tracking
