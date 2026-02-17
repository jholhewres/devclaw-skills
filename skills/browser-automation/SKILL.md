---
name: browser-automation
version: 0.1.0
author: devclaw
description: "Browser automation with Playwright — navigate, click, fill forms, scrape"
category: automation
tags: [browser, playwright, automation, scraping, web]
requires:
  bins: [npx, node]
---
# Browser Automation

Automate web browsers using Playwright for testing, scraping, and automation.

## Setup

```bash
# Install Playwright
npm install -g playwright
npx playwright install
```

## Quick Scripts

```bash
# Take screenshot
npx playwright screenshot https://example.com screenshot.png

# Generate PDF
npx playwright pdf https://example.com page.pdf
```

## Navigate & Screenshot

```javascript
// script.js
const { chromium } = require('playwright');

(async () => {
  const browser = await chromium.launch();
  const page = await browser.newPage();

  await page.goto('https://example.com');
  await page.screenshot({ path: 'screenshot.png', fullPage: true });

  await browser.close();
})();
```

```bash
node script.js
```

## Click & Fill Forms

```javascript
const { chromium } = require('playwright');

(async () => {
  const browser = await chromium.launch({ headless: false });
  const page = await browser.newPage();

  await page.goto('https://example.com/login');

  // Fill form
  await page.fill('#email', 'user@example.com');
  await page.fill('#password', 'secretpassword');

  // Click button
  await page.click('button[type="submit"]');

  // Wait for navigation
  await page.waitForURL('**/dashboard');

  await browser.close();
})();
```

## Extract Data

```javascript
const { chromium } = require('playwright');

(async () => {
  const browser = await chromium.launch();
  const page = await browser.newPage();

  await page.goto('https://example.com/products');

  // Extract text
  const title = await page.textContent('h1');

  // Extract multiple elements
  const products = await page.$$eval('.product', items =>
    items.map(item => ({
      name: item.querySelector('.name')?.textContent,
      price: item.querySelector('.price')?.textContent
    }))
  );

  console.log(JSON.stringify(products, null, 2));

  await browser.close();
})();
```

## Handle Dialogs & Alerts

```javascript
page.on('dialog', async dialog => {
  console.log(dialog.message());
  await dialog.accept();
});
```

## Wait Strategies

```javascript
// Wait for element
await page.waitForSelector('.loaded');

// Wait for specific text
await page.waitForSelector('text=Welcome');

// Wait for network idle
await page.waitForLoadState('networkidle');

// Wait with timeout
await page.waitForSelector('.element', { timeout: 5000 });
```

## Authentication

```javascript
// Save auth state
const context = await browser.newContext();
await context.storageState({ path: 'auth.json' });

// Load auth state
const context = await browser.newContext({ storageState: 'auth.json' });
```

## Command Line

```bash
# Run with specific browser
npx playwright test --project=chromium
npx playwright test --project=firefox
npx playwright test --project=webkit

# Run headed (visible browser)
npx playwright test --headed

# Debug mode
npx playwright test --debug

# Generate code from actions
npx playwright codegen https://example.com
```

## Tips

- Use `headless: false` to see browser during development
- Use `codegen` to record actions and generate code
- Use `waitForLoadState('networkidle')` for dynamic content
- Use `page.pause()` for debugging
- Screenshots: PNG, JPEG supported

## Triggers

browser automation, playwright, web scraping, automate browser,
click, fill form, screenshot, web testing
