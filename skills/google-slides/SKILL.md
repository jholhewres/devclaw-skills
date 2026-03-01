---
name: google-slides
description: Access Google Slides via Google API using the google_api tool
homepage: https://developers.google.com/slides/api/reference/rest
metadata:
  openclaw:
    requires:
      auth_profile: google
    emoji: "🎨"
---

# Google Slides

Read and write presentations using the `google_api` tool with Google Slides API endpoints.

## Prerequisites

You need to configure an OAuth profile first:

```bash
auth_profile_add(
  provider="google-slides",
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

## Presentation ID

The presentation ID is found in the URL:
- URL: `https://docs.google.com/presentation/d/PRESENTATION_ID/edit`
- ID: The long string between `/d/` and `/edit`

## Common Operations

### Get Presentation

```bash
google_api(
  service="slides",
  endpoint="/presentations/PRESENTATION_ID",
  method="GET"
)
```

### Get Specific Slide

```bash
google_api(
  service="slides",
  endpoint="/presentations/PRESENTATION_ID/pages/SLIDE_OBJECT_ID",
  method="GET"
)
```

### Get Slide Thumbnail

```bash
google_api(
  service="slides",
  endpoint="/presentations/PRESENTATION_ID/pages/SLIDE_OBJECT_ID/thumbnail",
  method="GET",
  params={
    "thumbnailProperties.mimeType": "PNG",
    "thumbnailProperties.thumbnailSize": "MEDIUM"
  }
)
```

### Create New Presentation

```bash
google_api(
  service="slides",
  endpoint="/presentations",
  method="POST",
  body={
    "title": "My Presentation"
  }
)
```

### Add Slide

```bash
google_api(
  service="slides",
  endpoint="/presentations/PRESENTATION_ID:batchUpdate",
  method="POST",
  body={
    "requests": [
      {
        "createSlide": {
          "slideLayoutReference": {
            "predefinedLayout": "BLANK"
          }
        }
      }
    ]
  }
)
```

### Insert Text Box

```bash
google_api(
  service="slides",
  endpoint="/presentations/PRESENTATION_ID:batchUpdate",
  method="POST",
  body={
    "requests": [
      {
        "createShape": {
          "objectId": "MyTextBox_001",
          "shapeType": "TEXT_BOX",
          "elementProperties": {
            "pageObjectId": "SLIDE_OBJECT_ID",
            "size": {
              "width": {
                "magnitude": 300,
                "unit": "PT"
              },
              "height": {
                "magnitude": 100,
                "unit": "PT"
              }
            },
            "transform": {
              "scaleX": 1,
              "scaleY": 1,
              "translateX": 50,
              "translateY": 50,
              "unit": "PT"
            }
          }
        }
      },
      {
        "insertText": {
          "objectId": "MyTextBox_001",
          "insertionIndex": 0,
          "text": "Hello, World!"
        }
      }
    ]
  }
)
```

### Replace Text

```bash
google_api(
  service="slides",
  endpoint="/presentations/PRESENTATION_ID:batchUpdate",
  method="POST",
  body={
    "requests": [
      {
        "replaceAllText": {
          "containsText": {
            "text": "{{COMPANY}}",
            "matchCase": true
          },
          "replaceText": "Acme Corp"
        }
      }
    ]
  }
)
```

### Delete Slide

```bash
google_api(
  service="slides",
  endpoint="/presentations/PRESENTATION_ID:batchUpdate",
  method="POST",
  body={
    "requests": [
      {
        "deleteObject": {
          "objectId": "SLIDE_OBJECT_ID"
        }
      }
    ]
  }
)
```

### Duplicate Slide

```bash
google_api(
  service="slides",
  endpoint="/presentations/PRESENTATION_ID:batchUpdate",
  method="POST",
  body={
    "requests": [
      {
        "duplicateObject": {
          "objectId": "SLIDE_OBJECT_ID",
          "objectIds": {
            "SLIDE_OBJECT_ID": "NEW_SLIDE_ID"
          }
        }
      }
    ]
  }
)
```

## Predefined Layouts

- `BLANK` - Blank slide
- `TITLE` - Title slide
- `TITLE_AND_BODY` - Title with body
- `TITLE_AND_TWO_COLUMNS` - Title with two columns
- `TITLE_ONLY` - Title only
- `SECTION_HEADER` - Section header
- `SECTION_TITLE_AND_DESCRIPTION` - Section with title and description
- `ONE_COLUMN_TEXT` - Single column text
- `MAIN_POINT` - Main point
- `BIG_NUMBER` - Big number

## Shape Types

- `TEXT_BOX` - Text container
- `RECTANGLE` - Rectangle shape
- `ROUND_RECTANGLE` - Rounded rectangle
- `ELLIPSE` - Ellipse/circle
- `ARROW` - Arrow
- `CALLOUT` - Callout box

## Image Upload

To insert images, first upload to Drive, then use the image URL in the API.

## Using Specific Profiles

```bash
google_api(
  service="slides",
  endpoint="/presentations/PRESENTATION_ID",
  method="GET",
  profile="work"
)
```

## Error Handling

If you see "No valid OAuth profile found":
1. Check configured profiles: `auth_profile_list()`
2. Add a profile: `auth_profile_add(provider="google-slides", name="default", mode="oauth")`
3. Complete the OAuth flow in the WebUI

## API Reference

Full API documentation: https://developers.google.com/slides/api/reference/rest
