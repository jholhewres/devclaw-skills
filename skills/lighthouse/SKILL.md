---
name: lighthouse
description: "Web performance audit with Google Lighthouse"
metadata: {"openclaw":{"always":false,"emoji":"🏠"}}
---
# Lighthouse

Audit performance, accessibility, SEO and best practices for websites.

## Setup

1. **Check if installed:**
   ```bash
   command -v node && node --version
   command -v npx && npx lighthouse --version 2>/dev/null || npm list -g lighthouse 2>/dev/null
   ```

2. **Install:**
   ```bash
   # Node.js (required for npx / lighthouse)
   # macOS
   brew install node

   # Ubuntu / Debian
   sudo apt update && sudo apt install -y nodejs npm

   # Lighthouse (use npx, no install needed)
   npx lighthouse <url> --output=json --chrome-flags="--headless --no-sandbox"

   # Or install globally
   npm install -g lighthouse
   ```

## Run Audit

```bash
# Full audit (output to terminal)
npx lighthouse <url> --output=json --chrome-flags="--headless --no-sandbox" 2>/dev/null | jq '{
  performance: .categories.performance.score,
  accessibility: .categories.accessibility.score,
  bestPractices: .categories["best-practices"].score,
  seo: .categories.seo.score
}'

# Save HTML report
npx lighthouse <url> --output=html --output-path=./lighthouse-report.html --chrome-flags="--headless --no-sandbox"

# Performance only
npx lighthouse <url> --only-categories=performance --output=json --chrome-flags="--headless --no-sandbox"
```

## Categories

| Category | What it measures |
|----------|------------------|
| Performance | LCP, FID, CLS, TTFB, speed index |
| Accessibility | Contrast, alt text, ARIA, keyboard navigation |
| Best Practices | HTTPS, console errors, deprecated APIs |
| SEO | Meta tags, structured data, mobile-friendly |

## Interpret Scores

- **90-100**: Good (green)
- **50-89**: Needs Improvement (orange)
- **0-49**: Poor (red)

## Key Metrics (Core Web Vitals)

| Metric | Good | Needs Improvement | Poor |
|--------|------|-------------------|------|
| LCP (Largest Contentful Paint) | < 2.5s | 2.5s - 4s | > 4s |
| FID (First Input Delay) | < 100ms | 100ms - 300ms | > 300ms |
| CLS (Cumulative Layout Shift) | < 0.1 | 0.1 - 0.25 | > 0.25 |

## Compare Results

```bash
# Run multiple times and compare
for i in 1 2 3; do
  npx lighthouse <url> --output=json --chrome-flags="--headless --no-sandbox" 2>/dev/null | \
    jq -r '"Run '$i': perf=\(.categories.performance.score) a11y=\(.categories.accessibility.score)"'
done
```

## Tips

- Run 3+ times and average (results vary)
- Use `--chrome-flags="--headless --no-sandbox"` for servers
- Scores are multiplied by 100 (0.95 = 95 points)
- Use `--preset=desktop` to simulate desktop (default is mobile)
- Install Chrome/Chromium on server if not available
