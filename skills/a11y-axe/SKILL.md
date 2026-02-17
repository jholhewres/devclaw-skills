---
name: a11y-axe
description: "Accessibility audit with axe-core via browser injection"
---
# Accessibility Audit (axe-core)

Use the **bash** tool with npx to run accessibility audits.

## Quick Audit
```bash
npx @axe-core/cli <URL> --stdout 2>/dev/null | head -100
```

## With Specific Rules
```bash
npx @axe-core/cli <URL> --tags wcag2a,wcag2aa --stdout 2>/dev/null
```

## Specific Element
```bash
npx @axe-core/cli <URL> --include "#main-content" --stdout 2>/dev/null
```

## WCAG Levels
| Tag | Standard |
|-----|----------|
| wcag2a | WCAG 2.0 Level A |
| wcag2aa | WCAG 2.0 Level AA |
| wcag21a | WCAG 2.1 Level A |
| wcag21aa | WCAG 2.1 Level AA |
| best-practice | Best practices |

## Tips
- Focus on critical and serious violations first
- Use --include to scope audit to specific sections
- Run against localhost dev server for fastest results
- Combine with color-contrast skill for detailed contrast checks
