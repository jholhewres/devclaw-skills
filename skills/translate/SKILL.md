---
name: translate
version: 0.1.0
author: devclaw
description: "Translate text between languages — using LLM capabilities and web APIs"
category: builtin
tags: [translate, language, i18n, multilingual]
---
# Translate

You can translate text between any languages. Use your built-in multilingual capabilities first, and external APIs for verification when needed.

## Built-in translation (preferred)

As a multilingual LLM, you can directly translate most text. Just do it inline when asked.

For formal or critical translations, verify with external APIs:

## External verification (LibreTranslate — free, no key)

```bash
# Translate text
curl -s -X POST "https://libretranslate.com/translate" \
  -H "Content-Type: application/json" \
  -d '{
    "q": "TEXT_TO_TRANSLATE",
    "source": "SOURCE_LANG_CODE",
    "target": "TARGET_LANG_CODE"
  }' | jq -r '.translatedText'

# Detect language
curl -s -X POST "https://libretranslate.com/detect" \
  -H "Content-Type: application/json" \
  -d '{"q": "TEXT_TO_DETECT"}' | jq '.[0] | {language, confidence}'

# List supported languages
curl -s "https://libretranslate.com/languages" | jq '.[] | {code, name}'
```

## Common language codes

| Language | Code |
|----------|------|
| Portuguese | pt |
| English | en |
| Spanish | es |
| French | fr |
| German | de |
| Italian | it |
| Japanese | ja |
| Chinese | zh |
| Korean | ko |
| Russian | ru |
| Arabic | ar |

## Tips

- For casual translations, use your built-in capabilities (faster, no API).
- For technical or legal text, mention that a professional review is recommended.
- When the user sends text in another language, offer to translate automatically.
- Preserve formatting (lists, headers, code blocks) during translation.
- If the source language is unclear, detect it first.
- Provide context-aware translations — don't translate proper nouns, brand names, or technical terms unless asked.
- For long texts, translate section by section to maintain quality.

## Triggers

translate, translation, what does this mean, how do you say, in english, en español,
traduzir, tradução, como se diz, em português
