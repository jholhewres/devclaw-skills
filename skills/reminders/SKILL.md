---
name: reminders
version: 0.1.0
author: devclaw
description: "Create, list, and manage reminders with date/time scheduling"
category: builtin
tags: [reminders, schedule, alerts, todo, due-dates]
---
# Reminders

You can create and manage time-based reminders for the user. Reminders are stored using the DevClaw scheduler and fire at the specified time.

## Creating reminders

Use the `cron_add` tool to schedule reminders at specific times:

```bash
# Reminder at a specific time today
# cron format: minute hour day month weekday
cron_add --id "rem-UNIQUE_ID" --schedule "0 15 14 2 *" --payload "📋 Reminder: client meeting at 3pm"

# Daily reminder at 9am
cron_add --id "daily-water" --schedule "0 9 * * *" --payload "💧 Time to drink water! Stay hydrated."

# Weekday reminder at 8:30am
cron_add --id "standup" --schedule "30 8 * * 1-5" --payload "🏃 Daily standup in 30 minutes!"

# Weekly reminder (every Monday at 10am)
cron_add --id "weekly-review" --schedule "0 10 * * 1" --payload "📊 Time for the weekly review!"

# Monthly reminder (1st of each month)
cron_add --id "bills" --schedule "0 10 1 * *" --payload "💰 Check monthly bills"
```

## Listing reminders

```bash
cron_list
```

## Removing reminders

```bash
cron_remove --id "rem-UNIQUE_ID"
```

## Date/time parsing guide

When the user says natural language, convert to cron:

| User says | Cron expression | Notes |
|-----------|----------------|-------|
| "amanhã às 9h" | `0 9 DD MM *` | Calculate tomorrow's date |
| "daqui a 2 horas" | Use `sleep` + background | Quick timer is better |
| "todo dia às 8h" | `0 8 * * *` | Daily |
| "segunda a sexta 9h" | `0 9 * * 1-5` | Weekdays |
| "toda segunda" | `0 9 * * 1` | Weekly |
| "dia 15 de cada mês" | `0 9 15 * *` | Monthly |
| "daqui a 30 minutos" | Use timer skill | Short-term = timer |

## Tips

- Generate unique IDs for reminders (e.g., "rem-" + timestamp or purpose).
- For reminders less than 1 hour away, use the **timer** skill instead (simpler).
- Always confirm the reminder time with the user before creating.
- Use the user's timezone (from config or USER.md) when interpreting times.
- Include a clear, descriptive message in the payload so the user understands the context.
- For one-shot reminders, suggest removing them after they fire.
- List existing reminders before creating duplicates.

## Triggers

remind me, reminder, set a reminder, lembrete, me lembre,
lembre-me amanhã, avise-me, agenda, scheduled reminder
