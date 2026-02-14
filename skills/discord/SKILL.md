---
name: discord
version: 0.1.0
author: goclaw
description: "Discord integration — send messages via webhooks and bots"
category: communication
tags: [discord, messaging, gaming, communities, webhooks]
requires:
  bins: [curl, jq]
  env: [DISCORD_WEBHOOK_URL]
---
# Discord

Interact with Discord using webhooks or the Bot API.

## Setup (Webhook - Easiest)

1. Go to Server Settings > Integrations > Webhooks
2. Click "New Webhook", name it, select channel
3. Copy the Webhook URL
4. Set environment variable:
   ```bash
   export DISCORD_WEBHOOK_URL="https://discord.com/api/webhooks/ID/TOKEN"
   ```

## Send Messages (Webhook)

```bash
# Simple message
curl -s -X POST "$DISCORD_WEBHOOK_URL" \
  -H "Content-Type: application/json" \
  -d '{"content": "Hello from goclaw!"}'

# With username and avatar override
curl -s -X POST "$DISCORD_WEBHOOK_URL" \
  -H "Content-Type: application/json" \
  -d '{
    "content": "Custom message",
    "username": "goclaw Bot",
    "avatar_url": "https://example.com/avatar.png"
  }'

# With embed (rich formatting)
curl -s -X POST "$DISCORD_WEBHOOK_URL" \
  -H "Content-Type: application/json" \
  -d '{
    "embeds": [{
      "title": "Notification",
      "description": "Something happened!",
      "color": 3447003,
      "fields": [
        {"name": "Status", "value": "Success", "inline": true},
        {"name": "Time", "value": "Now", "inline": true}
      ]
    }]
  }'
```

## Bot API (More Features)

```bash
# Requires DISCORD_BOT_TOKEN
export DISCORD_BOT_TOKEN="your-bot-token"

# Send message to channel
curl -s -X POST "https://discord.com/api/v10/channels/CHANNEL_ID/messages" \
  -H "Authorization: Bot $DISCORD_BOT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"content": "Hello!"}'

# Get channel messages
curl -s "https://discord.com/api/v10/channels/CHANNEL_ID/messages?limit=10" \
  -H "Authorization: Bot $DISCORD_BOT_TOKEN" | jq '.[] | {author: .author.username, content}'

# List guilds (servers)
curl -s "https://discord.com/api/v10/users/@me/guilds" \
  -H "Authorization: Bot $DISCORD_BOT_TOKEN" | jq '.[] | {id, name}'
```

## Embed Examples

```bash
# Status embed with color coding
curl -s -X POST "$DISCORD_WEBHOOK_URL" \
  -H "Content-Type: application/json" \
  -d '{
    "embeds": [{
      "title": "Deploy Status",
      "description": "Production deployment completed",
      "color": 3066993,
      "footer": {"text": "goclaw Automation"}
    }]
  }'

# Colors: Red=15158332, Green=3066993, Blue=3447003, Yellow=16776960
```

## Tips

- Webhooks are easiest for simple notifications
- Bot API requires creating an app at https://discord.com/developers/applications
- Embed color is a decimal integer (not hex)
- Max message length: 2000 characters
- Rate limit: ~5 requests/second per channel

## Triggers

discord, send discord message, discord webhook, discord notification,
notify discord, discord bot
