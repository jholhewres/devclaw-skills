---
name: pomodoro
description: "Pomodoro technique: focus sessions, breaks, tracking"
---
# Pomodoro

Use **cron_add**, **cron_list**, **cron_remove**, and **memory_save** tools for Pomodoro sessions.

## Start Session
1. Use **cron_add** with type="at" to schedule end-of-focus alert (25 min)
2. Use **memory_save** to log session start: "pomodoro started for <task> at <time>"
3. When timer fires, notify user: "Focus session done! Take a 5-minute break."

## Cycle
| Phase | Duration | Alert |
|-------|----------|-------|
| Focus | 25 min | "Time's up! Take a short break." |
| Short break | 5 min | "Break over. Ready for another focus session?" |
| Long break (every 4) | 15-30 min | "Long break done. Start a new cycle?" |

## Track Sessions
Use **memory_save** to log completed pomodoros:
- "pomodoro completed: <task>, <date>, 25 min"

Use **memory_search** to retrieve history:
- "pomodoro" → list all sessions

## Daily Summary
```bash
# Count today's pomodoros from memory
```
Use **memory_search** with query "pomodoro completed today" to count.

## Tips
- Ask user what they'll work on before starting
- Track interruptions: if user messages during focus, note it
- After 4 pomodoros, suggest a long break
- Use memory_save to build a productivity log over time
- Combine with the timer skill for the actual countdown
