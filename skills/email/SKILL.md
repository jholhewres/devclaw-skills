---
name: email
version: 0.1.0
author: devclaw
description: "Send emails via SMTP or API services (SendGrid, Mailgun, Resend)"
category: communication
tags: [email, smtp, sendgrid, mailgun, resend, mail]
requires:
  bins: [curl, jq]
---
# Email

Send emails using various email services.

## Option 1: Resend (Recommended - Free tier)

```bash
# Setup
export RESEND_API_KEY="re_xxx"

# Send email
curl -s -X POST "https://api.resend.com/emails" \
  -H "Authorization: Bearer $RESEND_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "from": "onboarding@resend.dev",
    "to": "recipient@example.com",
    "subject": "Hello from DevClaw",
    "html": "<h1>Hello!</h1><p>This is a test email.</p>"
  }'
```

## Option 2: SendGrid

```bash
# Setup
export SENDGRID_API_KEY="SG.xxx"

# Send email
curl -s -X POST "https://api.sendgrid.com/v3/mail/send" \
  -H "Authorization: Bearer $SENDGRID_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "personalizations": [{"to": [{"email": "recipient@example.com"}]}],
    "from": {"email": "sender@example.com"},
    "subject": "Hello from DevClaw",
    "content": [{"type": "text/plain", "value": "Email body text"}]
  }'
```

## Option 3: Mailgun

```bash
# Setup
export MAILGUN_API_KEY="key-xxx"
export MAILGUN_DOMAIN="mg.example.com"

# Send email
curl -s -X POST "https://api.mailgun.net/v3/$MAILGUN_DOMAIN/messages" \
  -u "api:$MAILGUN_API_KEY" \
  -d from="sender@$MAILGUN_DOMAIN" \
  -d to="recipient@example.com" \
  -d subject="Hello from DevClaw" \
  -d text="Email body text"
```

## Option 4: Postmark

```bash
# Setup
export POSTMARK_SERVER_TOKEN="xxx-xxx-xxx"

# Send email
curl -s -X POST "https://api.postmarkapp.com/email" \
  -H "X-Postmark-Server-Token: $POSTMARK_SERVER_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "From": "sender@example.com",
    "To": "recipient@example.com",
    "Subject": "Hello from DevClaw",
    "TextBody": "Email body text"
  }'
```

## Option 5: SMTP via curl (Basic)

```bash
# Requires smtp server access
curl -s --url "smtp://smtp.example.com:587" \
  --ssl-reqd \
  --mail-from "sender@example.com" \
  --mail-rcpt "recipient@example.com" \
  --user "username:password" \
  -T - <<EOF
From: sender@example.com
To: recipient@example.com
Subject: Hello from DevClaw

Email body text
EOF
```

## With Attachments (Resend)

```bash
# Send with attachment (base64 encoded)
curl -s -X POST "https://api.resend.com/emails" \
  -H "Authorization: Bearer $RESEND_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "from": "onboarding@resend.dev",
    "to": "recipient@example.com",
    "subject": "Document attached",
    "html": "<p>Please find attached document.</p>",
    "attachments": [{
      "filename": "document.pdf",
      "content": "'$(base64 -w0 document.pdf)'"
    }]
  }'
```

## Tips

- Resend: Free 3000 emails/month, easiest API
- SendGrid: Free 100 emails/day forever
- Mailgun: Free 5000 emails/month for 3 months
- Use HTML + plain text for better deliverability
- Always verify sender domain for production

## Triggers

send email, email, mail, smtp, sendgrid, mailgun, resend, email api
