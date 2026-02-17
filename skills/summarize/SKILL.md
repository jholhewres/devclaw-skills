---
name: summarize
version: 0.2.0
author: devclaw
description: "Summarize URLs, articles, YouTube videos, podcasts, and files"
category: data
tags: [summarize, transcript, youtube, podcast, text-extraction]
---
# Summarize

You can summarize content from various sources: web pages, articles, YouTube videos, and files.

## Summarizing a web page

1. Fetch the page content:

```bash
# Method 1: curl + HTML strip (works everywhere)
curl -sL "URL" | sed 's/<script[^>]*>.*<\/script>//g' | sed 's/<style[^>]*>.*<\/style>//g' | sed 's/<[^>]*>//g' | sed '/^[[:space:]]*$/d' | head -500

# Method 2: readability-cli (better extraction, if installed)
readable "URL" 2>/dev/null | head -500
```

2. Then summarize the extracted text using your reasoning capabilities.

## YouTube video transcripts

```bash
# Get auto-generated subtitles (requires yt-dlp)
yt-dlp --write-auto-subs --skip-download --sub-lang pt,en -o "/tmp/%(id)s" "VIDEO_URL" 2>/dev/null
cat /tmp/*.vtt 2>/dev/null | grep -v "^[0-9]" | grep -v "^$" | grep -v "WEBVTT" | grep -v "^NOTE" | grep -v "\-\->" | sort -u

# Alternative: get video info first
yt-dlp --print title --print duration_string --print description "VIDEO_URL" 2>/dev/null
```

## PDF files

```bash
# Extract text from PDF (requires pdftotext)
pdftotext /path/to/file.pdf - | head -500

# Alternative with python
python3 -c "
import subprocess
text = subprocess.check_output(['pdftotext', '/path/to/file.pdf', '-']).decode()
print(text[:5000])
"
```

## Summarization guidelines

When summarizing, follow these best practices:

1. **Brief summary** (default): 3-5 bullet points with the key takeaways.
2. **Detailed summary**: Structured with sections, preserving important details.
3. **TL;DR**: One sentence capturing the essence.

Always:
- Preserve key facts, names, dates, and numbers.
- Maintain the original language unless the user asks otherwise.
- For technical content, keep important code snippets and terminology.
- Mention the source URL/title in your summary.
- Ask the user what level of detail they want if not specified.

## Tips

- Combine with **web-search** to find content, then summarize it.
- For very long content, break into sections and summarize each.
- Use `head -N` or `tail -N` to limit input when content is too large.
- If `yt-dlp` is not installed: `pip install yt-dlp` or `brew install yt-dlp`.

## Triggers

summarize, summary, transcribe, transcript, what does this video say,
read and summarize, tl;dr, resumir, resumo, transcrever
