---
name: web-scraper
version: 0.1.0
author: goclaw
description: "Web scraping — extract data from websites using curl, pup, or Python"
category: data
tags: [scraper, web, extraction, crawling, html]
requires:
  bins: [curl, jq]
---
# Web Scraper

Extract data from websites using various tools.

## Option 1: curl + HTML Parsing

```bash
# Get page and extract with grep/sed
curl -sL "https://example.com" | grep -oP '<title>.*?</title>'

# Extract all links
curl -sL "https://example.com" | grep -oP 'href="[^"]*"' | cut -d'"' -f2

# Extract text between tags
curl -sL "https://example.com" | sed -n '/<p>/,/<\/p>/p'
```

## Option 2: Pup (HTML Parser)

```bash
# Install
go install github.com/ericchiang/pup@latest

# Extract elements
curl -sL "https://example.com" | pup 'title text{}'

# Extract all links
curl -sL "https://example.com" | pup 'a attr{href}'

# Extract structured data
curl -sL "https://example.com/products" | pup '.product json{}'

# Extract with CSS selectors
curl -sL "https://example.com" | pup 'div.content h1 text{}'

# Multiple attributes
curl -sL "https://example.com" | pup 'a attr{href} attr{title}'
```

## Option 3: Python + BeautifulSoup

```python
# scraper.py
import requests
from bs4 import BeautifulSoup
import json

url = "https://example.com"
response = requests.get(url)
soup = BeautifulSoup(response.text, 'html.parser')

# Extract title
title = soup.find('title').text

# Extract all links
links = [a.get('href') for a in soup.find_all('a', href=True)]

# Extract by class
products = soup.find_all(class_='product')
data = [{'name': p.find(class_='name').text,
         'price': p.find(class_='price').text} for p in products]

print(json.dumps(data, indent=2))
```

```bash
pip install beautifulsoup4 requests
python scraper.py
```

## Option 4: Python + Playwright

```python
# dynamic_scraper.py
from playwright.sync_api import sync_playwright
import json

with sync_playwright() as p:
    browser = p.chromium.launch()
    page = browser.new_page()
    page.goto('https://example.com')

    # Wait for dynamic content
    page.wait_for_selector('.loaded')

    # Extract data
    products = page.eval_on_selector_all('.product', '''
        items => items.map(item => ({
            name: item.querySelector('.name')?.textContent,
            price: item.querySelector('.price')?.textContent
        }))
    ''')

    print(json.dumps(products, indent=2))
    browser.close()
```

## Handling JavaScript Sites

```bash
# Use curl with embedded JS rendering (if available)
# Or use Playwright/Puppeteer for JS-heavy sites

# Basic: check if content is in initial HTML
curl -sL "https://example.com" | grep "product"
```

## Rate Limiting & Ethics

```bash
# Add delay between requests
for url in "${urls[@]}"; do
  curl -sL "$url" | pup 'title text{}'
  sleep 2
done

# Set user agent
curl -sL -A "Mozilla/5.0" "https://example.com"

# Respect robots.txt
curl -sL "https://example.com/robots.txt"
```

## API-based Scraping (When Available)

```bash
# Many sites have JSON APIs - check network tab
curl -sL "https://api.example.com/v1/products" | jq '.data[]'

# Common API patterns
curl -sL "https://example.com/api/products?page=1&limit=50" | jq '.'
```

## Tips

- Check if an API exists before scraping HTML
- Respect `robots.txt` and rate limits
- Use `--compressed` with curl for gzip support
- Add `-L` to follow redirects
- Use `-H "Accept: application/json"` for API endpoints
- Handle pagination with `&page=N` parameters

## Triggers

scrape, scraper, web scraping, extract data, crawl website,
html parsing, web extraction
