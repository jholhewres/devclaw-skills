---
name: calendar
version: 0.3.0
author: devclaw
description: "Google Calendar — list, create, and manage events via gcal CLI"
category: productivity
tags: [google, calendar, scheduling, events, meetings, gcal]
requires:
  any_bins: [gcal, gcalcli]
---
# Google Calendar

Manage Google Calendar events from the command line using `gcal` (Node.js) or `gcalcli` (Python).

## Setup (gcal — recommended)

### 1. Install

```bash
sudo npm install -g gcal
```

### 2. Create Google OAuth credentials

1. Go to https://console.cloud.google.com/apis/credentials
2. Create a project (or select existing)
3. Enable **Google Calendar API**: https://console.cloud.google.com/apis/library/calendar-json.googleapis.com
4. Go to **Credentials** → **Create Credentials** → **OAuth client ID**
5. Application type: **Desktop app** (or "Other")
6. Download the JSON file
7. Save as `~/.client_secret.json`

### 3. Authorize (one-time)

```bash
# Generate authorization URL
gcal generateUrl

# Open the URL in your browser, authorize, and copy the code
gcal storeToken CODE_FROM_BROWSER
```

The token is saved at `~/calendar_api_token.json` with auto-refresh.

## Usage with gcal

### List events

```bash
# Today's events
gcal list

# Tomorrow
gcal list tomorrow

# Date range (natural language)
gcal list 'from Monday to Friday'
gcal list 'from 03/01/2026 to 03/07/2026'

# Specific dates
gcal list -f 2026-02-14 -t 2026-02-21
```

### Create events

```bash
# Natural language (preferred)
gcal insert 'Meeting with João tomorrow from 3pm to 4pm'
gcal insert 'Lunch on Friday at noon for 1 hour'
gcal insert 'Dentist appointment March 5 at 10am'

# Explicit parameters
gcal insert -s 'Team standup' -d 2026-02-15 -t 09:00 -D 30m
gcal insert -s 'Workshop' -d 2026-02-20 -t 14:00 -D 2h
```

### Bulk insert (from JSON)

```bash
cat > /tmp/events.json << 'EOF'
[{
  "calendarId": "primary",
  "resource": {
    "summary": "Sprint planning",
    "start": { "dateTime": "2026-02-17T10:00:00" },
    "end": { "dateTime": "2026-02-17T11:00:00" }
  }
}, {
  "calendarId": "primary",
  "resource": {
    "summary": "Retrospective",
    "start": { "dateTime": "2026-02-17T14:00:00" },
    "end": { "dateTime": "2026-02-17T15:00:00" }
  }
}]
EOF
gcal bulk -e /tmp/events.json
```

## Alternative: gcalcli (Python)

### Install

```bash
pip install gcalcli
```

### Usage

```bash
# List events
gcalcli agenda
gcalcli agenda --nostarted "2026-02-14" "2026-02-21"

# Create event
gcalcli quick "Meeting with team tomorrow 3pm 1hr"
gcalcli add --title "Meeting" --where "Room 1" --when "tomorrow 3pm" --duration 60

# Calendar view
gcalcli calw 2
```

## Tips

- Always confirm with the user before creating or deleting events.
- Parse natural language dates: "amanhã", "próxima segunda", "daqui a 2 horas".
- Use the user's timezone (from config or USER.md).
- For recurring events, mention the recurrence pattern to the user.
- When listing events, show: title, date/time, duration, location.
- `gcal insert` with natural language is the fastest way to create events.
- Token auto-refreshes — no need to re-authorize after initial setup.

## Triggers

calendar, what's on my calendar, schedule a meeting, create an event,
check my schedule, free time, available, minha agenda, agendar,
marcar reunião, calendário, próximos compromissos
