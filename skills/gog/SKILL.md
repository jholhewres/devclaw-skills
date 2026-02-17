---
name: gog
version: 0.2.0
author: devclaw
description: "Google Workspace integration — Gmail, Calendar, Drive via gog CLI"
category: productivity
tags: [google, gmail, calendar, drive, workspace, email]
requires:
  bins: [gog]
---
# Google Workspace (gog)

Full Google Workspace access via the `gog` CLI: Gmail, Calendar, and Drive.

## Gmail

```bash
# List recent emails
gog gmail list --limit 10

# Search emails
gog gmail list --query "from:boss is:unread" --limit 10
gog gmail list --query "subject:report after:2026/01/01"

# Read a specific email
gog gmail read --id EMAIL_ID

# Send email
gog gmail send --to "user@example.com" --subject "Subject" --body "Message body"

# Reply to email
gog gmail reply --id EMAIL_ID --body "Reply text"

# Send with CC
gog gmail send --to "user@example.com" --cc "other@example.com" --subject "Subject" --body "Body"
```

## Calendar

```bash
# List upcoming events
gog calendar list --days 7

# Create event
gog calendar create --title "Meeting" --start "2026-02-15T15:00:00" --duration "1h"
gog calendar create --title "Lunch" --start "tomorrow 12:00" --duration "1h" --location "Restaurant" --attendees "friend@example.com"

# Delete event
gog calendar delete --id EVENT_ID
```

## Drive

```bash
# List files
gog drive list --limit 20
gog drive list --query "type:document name:report"

# Search in folder
gog drive list --folder FOLDER_ID

# Download file
gog drive download --id FILE_ID --output /tmp/downloaded_file

# Upload file
gog drive upload --path ./local_file.pdf --folder FOLDER_ID
gog drive upload --path ./report.pdf --name "Q1 Report 2026"
```

## Tips

- **Always confirm** before sending emails or deleting events/files.
- Use search queries to narrow down results (Gmail supports full search syntax).
- For calendar, parse natural language: "tomorrow at 3pm", "next Friday".
- Check `gog auth status` if you get permission errors.
- For large file lists, use `--limit` to cap output.
- Combine with other skills: search web for info, then email a summary.

## Triggers

email, gmail, send email, check email, calendar, schedule, meeting,
google drive, drive, upload, download, minha agenda, enviar email, verificar emails
