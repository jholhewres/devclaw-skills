---
name: gmail
version: 0.1.0
author: devclaw
description: "Gmail API — read, send, and manage emails via Gmail"
category: communication
tags: [gmail, email, google, mail, inbox]
requires:
  bins: [curl, jq]
  env: [GMAIL_ACCESS_TOKEN]
---
# Gmail

Interact with Gmail using the Gmail API.

## Setup

1. Go to Google Cloud Console: https://console.cloud.google.com
2. Create project, enable Gmail API
3. Configure OAuth 2.0 consent screen
4. Create OAuth credentials (Desktop app)
5. Get access token via OAuth flow
6. Set environment variable:
   ```bash
   export GMAIL_ACCESS_TOKEN="ya29.a0..."
   ```

## List Messages

```bash
# List recent emails
curl -s "https://gmail.googleapis.com/gmail/v1/users/me/messages?maxResults=10" \
  -H "Authorization: Bearer $GMAIL_ACCESS_TOKEN" | jq '.messages'

# Search emails
curl -s "https://gmail.googleapis.com/gmail/v1/users/me/messages?q=is:unread+from:boss@company.com" \
  -H "Authorization: Bearer $GMAIL_ACCESS_TOKEN" | jq '.messages'

# Common search queries:
# is:unread, is:starred, has:attachment, larger:5M
# from:email@example.com, to:me, subject:urgent
# newer_than:7d, older_than:1m
```

## Read Message

```bash
# Get message details
curl -s "https://gmail.googleapis.com/gmail/v1/users/me/messages/MESSAGE_ID?format=full" \
  -H "Authorization: Bearer $GMAIL_ACCESS_TOKEN" | jq '{
    id, snippet,
    from: (.payload.headers[] | select(.name == "From") | .value),
    subject: (.payload.headers[] | select(.name == "Subject") | .value),
    date: (.payload.headers[] | select(.name == "Date") | .value)
  }'

# Get raw message (including body)
curl -s "https://gmail.googleapis.com/gmail/v1/users/me/messages/MESSAGE_ID?format=raw" \
  -H "Authorization: Bearer $GMAIL_ACCESS_TOKEN" | jq -r '.raw' | base64 -d
```

## Send Email

```bash
# Create raw email (RFC 2822 format)
RAW_EMAIL=$(cat <<EOF | base64 -w0
From: me@gmail.com
To: recipient@example.com
Subject: Test Email from DevClaw

This is the email body.
Can have multiple lines.
EOF
)

# Send
curl -s -X POST "https://gmail.googleapis.com/gmail/v1/users/me/messages/send" \
  -H "Authorization: Bearer $GMAIL_ACCESS_TOKEN" \
  -H "Content-Type: application/json" \
  -d "{\"raw\": \"$RAW_EMAIL\"}"
```

## Manage Labels

```bash
# List labels
curl -s "https://gmail.googleapis.com/gmail/v1/users/me/labels" \
  -H "Authorization: Bearer $GMAIL_ACCESS_TOKEN" | jq '.labels[]'

# Add label to message
curl -s -X POST "https://gmail.googleapis.com/gmail/v1/users/me/messages/MESSAGE_ID/modify" \
  -H "Authorization: Bearer $GMAIL_ACCESS_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"addLabelIds": ["Label_123"]}'
```

## Mark as Read/Star

```bash
# Mark as read
curl -s -X POST "https://gmail.googleapis.com/gmail/v1/users/me/messages/MESSAGE_ID/modify" \
  -H "Authorization: Bearer $GMAIL_ACCESS_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"removeLabelIds": ["UNREAD"]}'

# Star email
curl -s -X POST "https://gmail.googleapis.com/gmail/v1/users/me/messages/MESSAGE_ID/modify" \
  -H "Authorization: Bearer $GMAIL_ACCESS_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"addLabelIds": ["STARRED"]}'
```

## Tips

- Access tokens expire after ~1 hour, refresh using refresh token
- Use `format=metadata` for headers only (faster)
- Max results per request: 500
- Label IDs: INBOX, SENT, DRAFT, SPAM, TRASH, UNREAD, STARRED

## Triggers

gmail, email, check email, send email, read gmail, inbox, gmail api
