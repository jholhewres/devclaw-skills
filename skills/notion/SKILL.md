---
name: notion
version: 0.1.0
author: devclaw
description: "Notion API — create, read, search pages, databases, and blocks"
category: productivity
tags: [notion, notes, database, wiki, project-management]
requires:
  env: [NOTION_API_KEY]
---
# Notion

Manage Notion pages, databases, and blocks via the REST API.

## Setup

1. Create an integration at https://notion.so/my-integrations
2. Copy the API key (starts with `ntn_` or `secret_`)
3. Store it in the vault: `copilot config vault-set NOTION_API_KEY=ntn_your_key`
4. **Important:** Share target pages/databases with your integration (click "..." → "Connect to" → integration name)

## API Basics

All requests need these headers:

```bash
curl -s "https://api.notion.com/v1/..." \
  -H "Authorization: Bearer $NOTION_API_KEY" \
  -H "Notion-Version: 2025-09-03" \
  -H "Content-Type: application/json"
```

## Search pages and databases

```bash
curl -s -X POST "https://api.notion.com/v1/search" \
  -H "Authorization: Bearer $NOTION_API_KEY" \
  -H "Notion-Version: 2025-09-03" \
  -H "Content-Type: application/json" \
  -d '{"query": "SEARCH_TERM"}' | jq '.results[] | {id, object: .object, title: (.properties.title.title[0].plain_text // .title[0].plain_text // "untitled")}'
```

## Read a page

```bash
# Page properties
curl -s "https://api.notion.com/v1/pages/PAGE_ID" \
  -H "Authorization: Bearer $NOTION_API_KEY" \
  -H "Notion-Version: 2025-09-03" | jq '.properties'

# Page content (blocks)
curl -s "https://api.notion.com/v1/blocks/PAGE_ID/children" \
  -H "Authorization: Bearer $NOTION_API_KEY" \
  -H "Notion-Version: 2025-09-03" | jq '.results[] | {type, text: (.paragraph.rich_text[0].plain_text // .heading_1.rich_text[0].plain_text // .to_do.rich_text[0].plain_text // null)}'
```

## Create a page

```bash
# In a database
curl -s -X POST "https://api.notion.com/v1/pages" \
  -H "Authorization: Bearer $NOTION_API_KEY" \
  -H "Notion-Version: 2025-09-03" \
  -H "Content-Type: application/json" \
  -d '{
    "parent": {"database_id": "DATABASE_ID"},
    "properties": {
      "Name": {"title": [{"text": {"content": "Task title"}}]},
      "Status": {"select": {"name": "Todo"}}
    }
  }' | jq '{id, url}'

# Under another page
curl -s -X POST "https://api.notion.com/v1/pages" \
  -H "Authorization: Bearer $NOTION_API_KEY" \
  -H "Notion-Version: 2025-09-03" \
  -H "Content-Type: application/json" \
  -d '{
    "parent": {"page_id": "PARENT_PAGE_ID"},
    "properties": {
      "title": [{"text": {"content": "New Page Title"}}]
    },
    "children": [
      {"object": "block", "type": "paragraph", "paragraph": {"rich_text": [{"text": {"content": "Page content here."}}]}}
    ]
  }' | jq '{id, url}'
```

## Query a database

```bash
curl -s -X POST "https://api.notion.com/v1/databases/DATABASE_ID/query" \
  -H "Authorization: Bearer $NOTION_API_KEY" \
  -H "Notion-Version: 2025-09-03" \
  -H "Content-Type: application/json" \
  -d '{
    "filter": {"property": "Status", "select": {"equals": "In Progress"}},
    "sorts": [{"property": "Created", "direction": "descending"}]
  }' | jq '.results[] | {id, title: .properties.Name.title[0].plain_text, status: .properties.Status.select.name}'
```

## Add content to a page

```bash
curl -s -X PATCH "https://api.notion.com/v1/blocks/PAGE_ID/children" \
  -H "Authorization: Bearer $NOTION_API_KEY" \
  -H "Notion-Version: 2025-09-03" \
  -H "Content-Type: application/json" \
  -d '{
    "children": [
      {"object": "block", "type": "to_do", "to_do": {"rich_text": [{"text": {"content": "New task item"}}], "checked": false}}
    ]
  }'
```

## Tips

- The `Notion-Version` header is required — use `2025-09-03` (latest).
- Database IDs and page IDs can be found in the URL: `notion.so/workspace/PAGE_ID`.
- Always search first to find the right page/database before creating new ones.
- The integration only sees pages that have been explicitly shared with it.
- For complex queries, use filters and sorts to narrow results.

## Triggers

notion, add to notion, notion page, notion database, notion task,
adicionar no notion, buscar no notion, criar página
