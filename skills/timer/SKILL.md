---
name: timer
version: 0.1.0
author: devclaw
description: "Set timers, alarms, and Pomodoro sessions — runs in background"
category: builtin
tags: [timer, alarm, pomodoro, countdown, reminder]
---
# Timer

You can set timers and alarms that run in the background. Use the DevClaw scheduler (`cron_add`) for precise time-based alerts or bash `sleep` for quick countdowns.

## Quick timers (using sleep)

```bash
# Simple countdown — sends a message when done
sleep 300 && echo "⏰ Timer de 5 minutos finalizado!"

# With a custom message
sleep 600 && echo "⏰ Hora de verificar o forno!"

# 30 seconds
sleep 30 && echo "⏰ 30 segundos! Tempo esgotado."
```

> **Importante:** Use `bash background:true` para que o timer rode em segundo plano sem bloquear.

## Using the scheduler (for precise times)

```bash
# Reminder at a specific time (one-shot cron)
# Format: minute hour day month weekday
cron_add --id "reminder-123" --schedule "30 14 * * *" --payload "Lembrete: reunião às 15h"

# Remove after it fires
cron_remove --id "reminder-123"
```

## Pomodoro technique

```bash
# Work session (25 min)
sleep 1500 && echo "🍅 Pomodoro finalizado! Hora de uma pausa de 5 minutos."

# Short break (5 min)
sleep 300 && echo "🔔 Pausa acabou! Volte ao trabalho."

# Long break (15 min, after 4 pomodoros)
sleep 900 && echo "☕ Pausa longa acabou! Pronto para mais um ciclo?"
```

## Time conversion reference

| Input | Seconds |
|-------|---------|
| 30s | 30 |
| 1m | 60 |
| 5m | 300 |
| 10m | 600 |
| 15m | 900 |
| 25m | 1500 |
| 30m | 1800 |
| 1h | 3600 |
| 2h | 7200 |

## Tips

- Always run timers in background mode so the user can keep chatting.
- Convert natural language times: "5 minutos" → `sleep 300`, "meia hora" → `sleep 1800`.
- For recurring timers (e.g., "every 30 minutes"), use the scheduler with cron expressions.
- When a timer completes, notify the user with a clear message including the original purpose.
- For Pomodoro, track the cycle number if the user wants a full session.

## Triggers

timer, set a timer, alarm, set alarm, pomodoro, countdown,
temporizador, alarme, cronômetro, me avise em, lembrete em
