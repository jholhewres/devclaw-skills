---
name: google-sheets
description: Access Google Sheets via Google API using the google_api tool
homepage: https://developers.google.com/sheets/api/reference/rest
metadata:
  openclaw:
    requires:
      auth_profile: google
    emoji: "📊"
---

# Google Sheets

Read and write spreadsheet data using the `google_api` tool with Google Sheets API endpoints.

## Prerequisites

You need to configure an OAuth profile first:

```bash
auth_profile_add(
  provider="google-sheets",
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

## Spreadsheet ID

The spreadsheet ID is found in the URL:
- URL: `https://docs.google.com/spreadsheets/d/SPREADSHEET_ID/edit`
- ID: The long string between `/d/` and `/edit`

## Common Operations

### Get Spreadsheet Metadata

```bash
google_api(
  service="sheets",
  endpoint="/spreadsheets/SPREADSHEET_ID",
  method="GET"
)
```

### Read Values (Single Range)

```bash
google_api(
  service="sheets",
  endpoint="/spreadsheets/SPREADSHEET_ID/values/Sheet1!A1:D10",
  method="GET"
)
```

### Read Values (Multiple Ranges)

```bash
google_api(
  service="sheets",
  endpoint="/spreadsheets/SPREADSHEET_ID/values:batchGet",
  method="GET",
  params={
    "ranges": "Sheet1!A1:D10,Sheet2!A1:B5",
    "valueRenderOption": "FORMATTED_VALUE"
  }
)
```

### Write Values (Single Range)

```bash
google_api(
  service="sheets",
  endpoint="/spreadsheets/SPREADSHEET_ID/values/Sheet1!A1",
  method="PUT",
  body={
    "values": [
      ["Name", "Age", "City"],
      ["Alice", "30", "NYC"],
      ["Bob", "25", "LA"]
    ],
    "range": "Sheet1!A1:C3"
  }
)
```

### Append Values

```bash
google_api(
  service="sheets",
  endpoint="/spreadsheets/SPREADSHEET_ID/values/Sheet1!A1:append",
  method="POST",
  body={
    "values": [
      ["Charlie", "35", "Chicago"]
    ]
  }
)
```

### Clear Values

```bash
google_api(
  service="sheets",
  endpoint="/spreadsheets/SPREADSHEET_ID/values/Sheet1!A1:D10:clear",
  method="POST"
)
```

### Create New Spreadsheet

```bash
google_api(
  service="sheets",
  endpoint="/spreadsheets",
  method="POST",
  body={
    "properties": {
      "title": "My New Spreadsheet"
    },
    "sheets": [
      {
        "properties": {
          "title": "Sheet1"
        }
      }
    ]
  }
)
```

### Add New Sheet

```bash
google_api(
  service="sheets",
  endpoint="/spreadsheets/SPREADSHEET_ID:batchUpdate",
  method="POST",
  body={
    "requests": [
      {
        "addSheet": {
          "properties": {
            "title": "New Sheet"
          }
        }
      }
    ]
  }
)
```

## A1 Notation

Google Sheets uses A1 notation for cell references:

- `Sheet1!A1` - Single cell
- `Sheet1!A1:B10` - Range
- `Sheet1!A:A` - Entire column
- `Sheet1!1:1` - Entire row
- `Sheet1!A1:Z` - Range to end of data
- `'Sheet Name'!A1` - Sheet names with spaces need quotes

## Value Render Options

- `FORMATTED_VALUE` - Formatted value (default)
- `UNFORMATTED_VALUE` - Raw value
- `FORMULA` - Show formulas instead of values

## Using Specific Profiles

```bash
google_api(
  service="sheets",
  endpoint="/spreadsheets/SPREADSHEET_ID",
  method="GET",
  profile="work"
)
```

## Error Handling

If you see "No valid OAuth profile found":
1. Check configured profiles: `auth_profile_list()`
2. Add a profile: `auth_profile_add(provider="google-sheets", name="default", mode="oauth")`
3. Complete the OAuth flow in the WebUI

## API Reference

Full API documentation: https://developers.google.com/sheets/api/reference/rest
