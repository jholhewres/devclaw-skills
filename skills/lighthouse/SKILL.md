---
name: lighthouse
description: "Auditoria de performance web com Google Lighthouse"
metadata: {"openclaw":{"always":false,"emoji":"🏠"}}
---
# Lighthouse

Auditoria de performance, acessibilidade, SEO e best practices para sites.

## Rodar Auditoria

```bash
# Auditoria completa (output no terminal)
npx lighthouse <url> --output=json --chrome-flags="--headless --no-sandbox" 2>/dev/null | jq '{
  performance: .categories.performance.score,
  accessibility: .categories.accessibility.score,
  bestPractices: .categories["best-practices"].score,
  seo: .categories.seo.score
}'

# Salvar report HTML
npx lighthouse <url> --output=html --output-path=./lighthouse-report.html --chrome-flags="--headless --no-sandbox"

# Só performance
npx lighthouse <url> --only-categories=performance --output=json --chrome-flags="--headless --no-sandbox"
```

## Categorias

| Categoria | O que mede |
|-----------|------------|
| Performance | LCP, FID, CLS, TTFB, speed index |
| Accessibility | Contraste, alt text, ARIA, navegação por teclado |
| Best Practices | HTTPS, console errors, deprecated APIs |
| SEO | Meta tags, structured data, mobile-friendly |

## Interpretar Scores

- **90-100**: Bom (verde)
- **50-89**: Precisa melhorar (laranja)
- **0-49**: Ruim (vermelho)

## Métricas Chave (Core Web Vitals)

| Métrica | Bom | Precisa Melhorar | Ruim |
|---------|-----|-------------------|------|
| LCP (Largest Contentful Paint) | < 2.5s | 2.5s - 4s | > 4s |
| FID (First Input Delay) | < 100ms | 100ms - 300ms | > 300ms |
| CLS (Cumulative Layout Shift) | < 0.1 | 0.1 - 0.25 | > 0.25 |

## Comparar Resultados

```bash
# Rodar múltiplas vezes e comparar
for i in 1 2 3; do
  npx lighthouse <url> --output=json --chrome-flags="--headless --no-sandbox" 2>/dev/null | \
    jq -r '"Run '$i': perf=\(.categories.performance.score) a11y=\(.categories.accessibility.score)"'
done
```

## Tips

- Rode 3+ vezes e tire a média (resultados variam)
- Use `--chrome-flags="--headless --no-sandbox"` para servers
- Scores são multiplicados por 100 (0.95 = 95 pontos)
- Use `--preset=desktop` para simular desktop (padrão é mobile)
- Instale Chrome/Chromium no server se não tiver
