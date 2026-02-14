---
name: telegram
version: 0.1.0
author: goclaw
description: "Telegram Bot API — send messages, photos, files to users and groups"
category: communication
tags: [telegram, messaging, bot, notifications]
requires:
  bins: [curl, jq]
  env: [TELEGRAM_BOT_TOKEN]
---
# Telegram

Interact with Telegram using the Bot API.

## Setup

1. Talk to @BotFather on Telegram
2. Send `/newbot` and follow instructions
3. Copy the bot token
4. Set environment variable:
   ```bash
   export TELEGRAM_BOT_TOKEN="123456789:ABCdefGHIjklMNOpqrsTUVwxyz"
   ```

## Get Chat ID

```bash
# Send a message to your bot first, then:
curl -s "https://api.telegram.org/bot$TELEGRAM_BOT_TOKEN/getUpdates" | jq '.result[-1].message.chat.id'

# For groups/channels, add bot as admin and send a message
```

## Send Messages

```bash
# Send text message
curl -s -X POST "https://api.telegram.org/bot$TELEGRAM_BOT_TOKEN/sendMessage" \
  -H "Content-Type: application/json" \
  -d '{"chat_id": "CHAT_ID", "text": "Hello from goclaw!"}'

# Markdown formatting
curl -s -X POST "https://api.telegram.org/bot$TELEGRAM_BOT_TOKEN/sendMessage" \
  -H "Content-Type: application/json" \
  -d '{
    "chat_id": "CHAT_ID",
    "text": "*Bold* _italic_ `code`",
    "parse_mode": "MarkdownV2"
  }'

# HTML formatting
curl -s -X POST "https://api.telegram.org/bot$TELEGRAM_BOT_TOKEN/sendMessage" \
  -H "Content-Type: application/json" \
  -d '{
    "chat_id": "CHAT_ID",
    "text": "<b>Bold</b> <i>italic</i> <code>code</code>",
    "parse_mode": "HTML"
  }'
```

## Send Photos & Files

```bash
# Send photo by URL
curl -s -X POST "https://api.telegram.org/bot$TELEGRAM_BOT_TOKEN/sendPhoto" \
  -H "Content-Type: application/json" \
  -d '{"chat_id": "CHAT_ID", "photo": "https://example.com/image.png"}'

# Send document
curl -s -X POST "https://api.telegram.org/bot$TELEGRAM_BOT_TOKEN/sendDocument" \
  -F "chat_id=CHAT_ID" \
  -F "document=@/path/to/file.pdf"
```

## Interactive Messages

```bash
# Message with inline keyboard buttons
curl -s -X POST "https://api.telegram.org/bot$TELEGRAM_BOT_TOKEN/sendMessage" \
  -H "Content-Type: application/json" \
  -d '{
    "chat_id": "CHAT_ID",
    "text": "Choose an option:",
    "reply_markup": {
      "inline_keyboard": [
        [{"text": "Yes", "callback_data": "yes"}, {"text": "No", "callback_data": "no"}]
      ]
    }
  }'
```

## Get Updates (Polling)

```bash
# Get new messages
curl -s "https://api.telegram.org/bot$TELEGRAM_BOT_TOKEN/getUpdates?offset=-1" | jq '.result'

# Get all pending updates
curl -s "https://api.telegram.org/bot$TELEGRAM_BOT_TOKEN/getUpdates" | jq '.result[]'
```

## Tips

- Chat ID for users: positive number (e.g., 123456789)
- Chat ID for groups: negative number (e.g., -100123456789)
- Channel posts need `@channelname` or channel ID
- Use `parse_mode: "MarkdownV2"` for rich formatting
- Escape special chars in MarkdownV2: _ * [ ] ( ) ~ ` > # + - = | { } . !

## Triggers

telegram, send telegram message, telegram bot, telegram notification,
notify telegram, tg message
