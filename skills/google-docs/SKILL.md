---
name: google-docs
description: Access Google Docs via Google API using the google_api tool
homepage: https://developers.google.com/docs/api/reference/rest
metadata:
  openclaw:
    requires:
      auth_profile: google
    emoji: "📝"
---

# Google Docs

Read and write documents using the `google_api` tool with Google Docs API endpoints.

## Prerequisites

You need to configure an OAuth profile first:

```bash
auth_profile_add(
  provider="google-docs",
  name="default",
  mode="oauth"
)
```

Or use a generic Google profile:

```bash
auth_profile_add(
  provider="google",
  name="work",
  mode="oauth"
)
```

## Document ID

The document ID is found in the URL:
- URL: `https://docs.google.com/document/d/DOCUMENT_ID/edit`
- ID: The long string between `/d/` and `/edit`

## Common Operations

### Get Document Content

```bash
google_api(
  service="docs",
  endpoint="/documents/DOCUMENT_ID",
  method="GET"
)
```

### Create New Document

```bash
google_api(
  service="docs",
  endpoint="/documents",
  method="POST",
  body={
    "title": "My New Document"
  }
)
```

### Batch Update (Insert Text)

```bash
google_api(
  service="docs",
  endpoint="/documents/DOCUMENT_ID:batchUpdate",
  method="POST",
  body={
    "requests": [
      {
        "insertText": {
          "location": {
            "index": 1
          },
          "text": "Hello, World!\n"
        }
      }
    ]
  }
)
```

### Insert Text at End

```bash
google_api(
  service="docs",
  endpoint="/documents/DOCUMENT_ID:batchUpdate",
  method="POST",
  body={
    "requests": [
      {
        "insertText": {
          "endOfSegmentLocation": {
            "segmentId": ""
          },
          "text": "Text at the end of the document."
        }
      }
    ]
  }
)
```

### Format Text (Bold)

```bash
google_api(
  service="docs",
  endpoint="/documents/DOCUMENT_ID:batchUpdate",
  method="POST",
  body={
    "requests": [
      {
        "updateTextStyle": {
          "range": {
            "startIndex": 1,
            "endIndex": 11
          },
          "textStyle": {
            "bold": true
          },
          "fields": "bold"
        }
      }
    ]
  }
)
```

### Create Heading

```bash
google_api(
  service="docs",
  endpoint="/documents/DOCUMENT_ID:batchUpdate",
  method="POST",
  body={
    "requests": [
      {
        "updateParagraphStyle": {
          "range": {
            "startIndex": 1,
            "endIndex": 20
          },
          "paragraphStyle": {
            "namedStyleType": "HEADING_1"
          },
          "fields": "namedStyleType"
        }
      }
    ]
  }
)
```

### Insert Table

```bash
google_api(
  service="docs",
  endpoint="/documents/DOCUMENT_ID:batchUpdate",
  method="POST",
  body={
    "requests": [
      {
        "insertTable": {
          "location": {
            "index": 1
          },
          "columns": 3,
          "rows": 2
        }
      }
    ]
  }
)
```

### Replace Text

```bash
google_api(
  service="docs",
  endpoint="/documents/DOCUMENT_ID:batchUpdate",
  method="POST",
  body={
    "requests": [
      {
        "replaceAllText": {
          "containsText": {
            "text": "{{NAME}}",
            "matchCase": true
          },
          "replaceText": "John Doe"
        }
      }
    ]
  }
)
```

## Document Structure

Google Docs content is structured with:
- **Body**: Main content container
- **Paragraphs**: Text blocks with styling
- **Text Runs**: Segments of text with same styling
- **Tables**: Grid-based content

## Common Index Positions

- `1` - Beginning of document
- End of document - Use `endOfSegmentLocation`

## Named Styles

- `NORMAL_TEXT` - Regular paragraph
- `HEADING_1` through `HEADING_6` - Headings
- `TITLE` - Document title
- `SUBTITLE` - Document subtitle

## Using Specific Profiles

```bash
google_api(
  service="docs",
  endpoint="/documents/DOCUMENT_ID",
  method="GET",
  profile="work"
)
```

## Error Handling

If you see "No valid OAuth profile found":
1. Check configured profiles: `auth_profile_list()`
2. Add a profile: `auth_profile_add(provider="google-docs", name="default", mode="oauth")`
3. Complete the OAuth flow in the WebUI

## API Reference

Full API documentation: https://developers.google.com/docs/api/reference/rest
