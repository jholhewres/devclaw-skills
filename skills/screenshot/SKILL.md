---
name: screenshot
version: 0.1.0
author: goclaw
description: "Capture website screenshots — full page, viewport, or element"
category: media
tags: [screenshot, capture, web, image, website]
requires:
  bins: [curl]
---
# Screenshot

Capture screenshots of websites using various methods.

## Option 1: Screenshot API Services

```bash
# apiflash.com (free tier)
export APIFLASH_KEY="your_key"
curl -s "https://api.apiflash.com/v1/urltoimage?url=https://example.com&access_key=$APIFLASH_KEY&full_page=true" -o screenshot.png

# screenshotapi.net
export SCREENSHOTAPI_KEY="your_key"
curl -s "https://api.screenshotapi.net/screenshot?token=$SCREENSHOTAPI_KEY&url=https://example.com&full=true" -o screenshot.png

# urlbox.io
curl -s "https://api.urlbox.io/v1/$URLBOX_KEY/png?url=example.com" -o screenshot.png
```

## Option 2: Puppeteer (Node.js)

```javascript
// screenshot.js
const puppeteer = require('puppeteer');

(async () => {
  const browser = await puppeteer.launch();
  const page = await browser.newPage();

  // Set viewport size
  await page.setViewport({ width: 1920, height: 1080 });

  await page.goto('https://example.com', { waitUntil: 'networkidle2' });

  // Full page screenshot
  await page.screenshot({ path: 'full.png', fullPage: true });

  // Viewport screenshot
  await page.screenshot({ path: 'viewport.png' });

  // Element screenshot
  const element = await page.$('.content');
  await element.screenshot({ path: 'element.png' });

  await browser.close();
})();
```

```bash
npm install puppeteer
node screenshot.js
```

## Option 3: Playwright

```bash
# One-liner
npx playwright screenshot https://example.com screenshot.png

# Full page
npx playwright screenshot --full-page https://example.com full.png

# Mobile viewport
npx playwright screenshot --viewport-size=375,667 https://example.com mobile.png

# Wait for selector
npx playwright screenshot --wait-for-selector=.content https://example.com loaded.png
```

## Option 4: wkhtmltoimage

```bash
# Install
sudo apt install wkhtmltopdf

# Basic screenshot
wkhtmltoimage https://example.com screenshot.png

# With quality options
wkhtmltoimage --quality 80 --width 1920 https://example.com screenshot.png
```

## Option 5: Cutycapt

```bash
# Install
sudo apt install cutycapt

# Capture
cutycapt --url=https://example.com --out=screenshot.png

# With dimensions
cutycapt --url=https://example.com --out=screenshot.png --min-width=1920 --min-height=1080
```

## Python Options

```python
# selenium_scraper.py
from selenium import webdriver
from selenium.webdriver.chrome.options import Options

options = Options()
options.add_argument('--headless')
options.add_argument('--window-size=1920,1080')

driver = webdriver.Chrome(options=options)
driver.get('https://example.com')
driver.save_screenshot('screenshot.png')
driver.quit()
```

## Advanced Options

```javascript
// With Playwright - advanced
const { chromium } = require('playwright');

(async () => {
  const browser = await chromium.launch();
  const page = await browser.newPage({
    viewport: { width: 1920, height: 1080 },
    deviceScaleFactor: 2 // Retina
  });

  // Authenticate
  await page.setExtraHTTPHeaders({
    'Authorization': 'Bearer token'
  });

  // Set cookies
  await page.context().addCookies([{
    name: 'session', value: 'xyz', domain: 'example.com'
  }]);

  await page.goto('https://example.com');
  await page.screenshot({ path: 'auth.png', fullPage: true });

  await browser.close();
})();
```

## Tips

- Use `fullPage: true` for complete page capture
- Set appropriate viewport size before capture
- Wait for `networkidle` for dynamic content
- Use `deviceScaleFactor: 2` for high-DPI screenshots
- Handle authentication with cookies or headers

## Triggers

screenshot, capture website, website screenshot, web capture,
full page screenshot, take screenshot
