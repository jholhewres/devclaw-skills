---
name: web-fetch
version: 0.2.0
author: devclaw
description: "Fetch URL content and extract readable text/markdown"
category: data
tags: [web, fetch, scrape, readability, markdown, curl]
---
# Web Fetch

You can fetch and read the content of any web page, API endpoint, or file URL.

## Fetching web pages (readable text)

```bash
# Fetch page content — strip HTML, get readable text
curl -sL "URL" | sed 's/<script[^>]*>.*<\/script>//g' | sed 's/<style[^>]*>.*<\/style>//g' | sed 's/<[^>]*>//g' | sed '/^[[:space:]]*$/d' | head -300

# Using readability-cli if installed (much better extraction)
readable "URL" 2>/dev/null || curl -sL "URL" | sed 's/<[^>]*>//g' | sed '/^$/d' | head -300

# Fetch only the <article> or <main> content (more targeted)
curl -sL "URL" | grep -oP '<(article|main)[^>]*>.*?</(article|main)>' | sed 's/<[^>]*>//g' | head -200
```

## Fetching JSON APIs

```bash
# Fetch and parse JSON
curl -s "API_URL" -H "Accept: application/json" | jq '.'

# Extract specific fields
curl -s "API_URL" | jq '.data[] | {id, name, value}'

# POST request with JSON body
curl -s -X POST "API_URL" -H "Content-Type: application/json" -d '{"key": "value"}' | jq '.'
```

## Checking headers only (HEAD request)

```bash
# Check content type, size, redirects without downloading
curl -sI "URL" | head -15
```

## Downloading files

```bash
# Download to a specific path
curl -sLo /tmp/output_file "URL"

# Download and check file type
curl -sL "URL" -o /tmp/downloaded && file /tmp/downloaded
```

## Tips

- Always use `-sL` (silent + follow redirects).
- For large pages, pipe through `head -N` to limit output.
- Strip HTML with `sed 's/<[^>]*>//g'` for plain text.
- Check `Content-Type` header first to decide parsing strategy.
- Respect robots.txt and rate limits.
- Combine with **web-search** skill: search first, then fetch top results.
- For paywalled or JS-heavy sites, results may be limited.

## Triggers

fetch, read this URL, open this link, what does this page say,
get content from, scrape, ler esta página, abrir link
