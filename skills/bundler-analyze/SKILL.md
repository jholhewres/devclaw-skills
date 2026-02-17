---
name: bundler-analyze
description: "Analyze JS bundle size with Vite/Webpack analyzers"
---
# Bundle Analyzer

Use the **bash** tool to analyze JavaScript bundle sizes.

## Vite
```bash
npx vite-bundle-visualizer 2>/dev/null
# generates stats.html

npx rollup-plugin-visualizer 2>/dev/null
```

## Webpack
```bash
npx webpack-bundle-analyzer stats.json 2>/dev/null
# First generate stats:
npx webpack --profile --json > stats.json
```

## Size Check (quick)
```bash
# Check build output sizes
ls -lhS dist/assets/*.js 2>/dev/null | head -20
ls -lhS dist/assets/*.css 2>/dev/null | head -10

# Total bundle size
du -sh dist/ 2>/dev/null
```

## Import Cost
```bash
npx import-cost <file.ts> 2>/dev/null
```

## Tips
- Run after npm run build to see actual production sizes
- Check for large dependencies with npm ls --depth=0
- Use dynamic import() for code splitting
- Compare sizes before/after changes
