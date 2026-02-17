---
name: trello
version: 0.1.0
author: devclaw
description: "Manage Trello boards, lists, and cards via REST API"
category: productivity
tags: [trello, kanban, tasks, boards, project-management]
requires:
  env: [TRELLO_API_KEY, TRELLO_TOKEN]
---
# Trello

Manage Trello boards, lists, and cards via the REST API.

## Setup

1. Get your API key: https://trello.com/app-key
2. Generate a token (click "Token" link on that page)
3. Store in vault:
   ```bash
   copilot config vault-set TRELLO_API_KEY=your_key
   copilot config vault-set TRELLO_TOKEN=your_token
   ```

## List boards

```bash
curl -s "https://api.trello.com/1/members/me/boards?key=$TRELLO_API_KEY&token=$TRELLO_TOKEN&fields=name,id,url" | jq '.[] | {name, id, url}'
```

## List lists in a board

```bash
curl -s "https://api.trello.com/1/boards/BOARD_ID/lists?key=$TRELLO_API_KEY&token=$TRELLO_TOKEN" | jq '.[] | {name, id}'
```

## List cards in a list

```bash
curl -s "https://api.trello.com/1/lists/LIST_ID/cards?key=$TRELLO_API_KEY&token=$TRELLO_TOKEN" | jq '.[] | {name, id, desc, due, labels: [.labels[].name]}'
```

## Create a card

```bash
curl -s -X POST "https://api.trello.com/1/cards?key=$TRELLO_API_KEY&token=$TRELLO_TOKEN" \
  -d "idList=LIST_ID" \
  -d "name=Card Title" \
  -d "desc=Card description" \
  -d "due=2026-02-20T10:00:00.000Z" | jq '{id, name, url}'
```

## Move a card

```bash
curl -s -X PUT "https://api.trello.com/1/cards/CARD_ID?key=$TRELLO_API_KEY&token=$TRELLO_TOKEN" \
  -d "idList=NEW_LIST_ID"
```

## Add a comment

```bash
curl -s -X POST "https://api.trello.com/1/cards/CARD_ID/actions/comments?key=$TRELLO_API_KEY&token=$TRELLO_TOKEN" \
  -d "text=Your comment here"
```

## Archive a card

```bash
curl -s -X PUT "https://api.trello.com/1/cards/CARD_ID?key=$TRELLO_API_KEY&token=$TRELLO_TOKEN" \
  -d "closed=true"
```

## Add a label

```bash
# List available labels
curl -s "https://api.trello.com/1/boards/BOARD_ID/labels?key=$TRELLO_API_KEY&token=$TRELLO_TOKEN" | jq '.[] | {id, name, color}'

# Add label to card
curl -s -X POST "https://api.trello.com/1/cards/CARD_ID/idLabels?key=$TRELLO_API_KEY&token=$TRELLO_TOKEN" \
  -d "value=LABEL_ID"
```

## Add a checklist

```bash
# Create checklist
curl -s -X POST "https://api.trello.com/1/checklists?key=$TRELLO_API_KEY&token=$TRELLO_TOKEN" \
  -d "idCard=CARD_ID" \
  -d "name=Checklist" | jq '{id, name}'

# Add item
curl -s -X POST "https://api.trello.com/1/checklists/CHECKLIST_ID/checkItems?key=$TRELLO_API_KEY&token=$TRELLO_TOKEN" \
  -d "name=Task item"
```

## Search

```bash
curl -s "https://api.trello.com/1/search?key=$TRELLO_API_KEY&token=$TRELLO_TOKEN&query=SEARCH_TERM&modelTypes=cards&cards_limit=10" | jq '.cards[] | {name, id, list: .idList}'
```

## Tips

- Board/List/Card IDs are in the Trello URL or via list commands.
- Rate limits: 300 requests/10s per API key, 100/10s per token.
- Always list boards first to find the right IDs.
- Use search to find cards quickly across all boards.
- Confirm before archiving or moving cards.

## Triggers

trello, add to trello, move card, create card, trello board,
mover card, criar card, quadro trello
