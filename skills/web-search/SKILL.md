---
name: web-search
version: 0.2.0
author: devclaw
description: "Search the web for current information using Brave Search API or DuckDuckGo"
category: data
tags: [search, web, brave, duckduckgo, news]
requires:
  any_env: [BRAVE_API_KEY]
---
# Web Search

You can search the web for current information, news, and recent events.

## Using Brave Search API (preferred — if `BRAVE_API_KEY` is available)

```bash
# General web search
curl -s "https://api.search.brave.com/res/v1/web/search?q=URL_ENCODED_QUERY&count=5" \
  -H "Accept: application/json" \
  -H "X-Subscription-Token: $BRAVE_API_KEY" | jq '.web.results[] | {title, url, description}'

# News search (past week)
curl -s "https://api.search.brave.com/res/v1/web/search?q=URL_ENCODED_QUERY&count=5&freshness=pw" \
  -H "Accept: application/json" \
  -H "X-Subscription-Token: $BRAVE_API_KEY" | jq '.web.results[] | {title, url, description}'

# Time-filtered search (today, past week, past month, past year)
# freshness values: pd (day), pw (week), pm (month), py (year)
curl -s "https://api.search.brave.com/res/v1/web/search?q=URL_ENCODED_QUERY&count=10&freshness=pd" \
  -H "Accept: application/json" \
  -H "X-Subscription-Token: $BRAVE_API_KEY" | jq '.web.results[] | {title, url, description}'
```

## Using DuckDuckGo (fallback — no API key needed)

```bash
# DuckDuckGo Instant Answer API (structured results)
curl -s "https://api.duckduckgo.com/?q=URL_ENCODED_QUERY&format=json&no_html=1" | jq '{
  heading: .Heading,
  abstract: .AbstractText,
  source: .AbstractSource,
  url: .AbstractURL,
  related: [.RelatedTopics[:5][] | {text: .Text, url: .FirstURL}]
}'

# DuckDuckGo HTML scrape (more results but less structured)
curl -s "https://html.duckduckgo.com/html/?q=URL_ENCODED_QUERY" \
  | grep -oP 'class="result__a"[^>]*href="\K[^"]+' | head -5
```

## Tips

- **URL-encode** the query: replace spaces with `+` or `%20`.
- Use Brave API if `BRAVE_API_KEY` is set — results are significantly better.
- For news and current events, add freshness parameters.
- Be specific in queries for better results.
- Combine with the **web-fetch** skill to read full content of interesting results.
- For local searches, include the location in the query.

## Triggers

search, search the web, google, look up, find information, what is, news about,
latest news, pesquisar, buscar na web, notícias
