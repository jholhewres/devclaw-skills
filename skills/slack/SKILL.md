---
name: slack
version: 0.1.0
author: goclaw
description: "Slack integration — send messages, list channels, manage workspace"
category: communication
tags: [slack, messaging, teams, channels, notifications]
requires:
  bins: [curl, jq]
  env: [SLACK_BOT_TOKEN]
---
# Slack

Interact with Slack workspaces using the Slack Web API.

## Setup

1. Create a Slack App at https://api.slack.com/apps
2. Add Bot Token Scopes: `chat:write`, `channels:read`, `users:read`
3. Install to workspace and copy the Bot User OAuth Token
4. Set environment variable:
   ```bash
   export SLACK_BOT_TOKEN="xoxb-your-token"
   ```

## Send Messages

```bash
# Send message to channel
curl -s -X POST "https://slack.com/api/chat.postMessage" \
  -H "Authorization: Bearer $SLACK_BOT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"channel": "C12345678", "text": "Hello from goclaw!"}'

# Send to user (DM)
curl -s -X POST "https://slack.com/api/chat.postMessage" \
  -H "Authorization: Bearer $SLACK_BOT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"channel": "@username", "text": "Direct message"}'

# Send with blocks (rich formatting)
curl -s -X POST "https://slack.com/api/chat.postMessage" \
  -H "Authorization: Bearer $SLACK_BOT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "channel": "C12345678",
    "blocks": [
      {"type": "section", "text": {"type": "mrkdwn", "text": "*Bold text*"}}
    ]
  }'
```

## List Channels

```bash
# List public channels
curl -s "https://slack.com/api/conversations.list?types=public_channel" \
  -H "Authorization: Bearer $SLACK_BOT_TOKEN" | jq '.channels[] | {id, name, num_members}'

# List all channels (including private)
curl -s "https://slack.com/api/conversations.list?types=public_channel,private_channel" \
  -H "Authorization: Bearer $SLACK_BOT_TOKEN" | jq '.channels[] | {id, name}'
```

## List Users

```bash
# List workspace members
curl -s "https://slack.com/api/users.list" \
  -H "Authorization: Bearer $SLACK_BOT_TOKEN" | jq '.members[] | {id, name, real_name}'
```

## Read Messages

```bash
# Get recent messages from channel
curl -s "https://slack.com/api/conversations.history?channel=C12345678&limit=10" \
  -H "Authorization: Bearer $SLACK_BOT_TOKEN" | jq '.messages[] | {user, text, ts}'
```

## Thread Replies

```bash
# Reply in a thread
curl -s -X POST "https://slack.com/api/chat.postMessage" \
  -H "Authorization: Bearer $SLACK_BOT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"channel": "C12345678", "thread_ts": "1234567890.123456", "text": "Thread reply"}'
```

## Add Reactions

```bash
# Add emoji reaction
curl -s -X POST "https://slack.com/api/reactions.add" \
  -H "Authorization: Bearer $SLACK_BOT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"channel": "C12345678", "timestamp": "1234567890.123456", "name": "thumbsup"}'
```

## Tips

- Channel IDs start with `C`, user IDs start with `U`, DM channel IDs start with `D`
- Use `@username` to send DMs without knowing the channel ID
- Check response `ok` field for success: `.ok == true`
- Rate limit: ~1 request/second for most methods

## Triggers

slack, send slack message, slack channel, slack dm, slack notification,
notify team, post to slack
