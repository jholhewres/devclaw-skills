---
name: airtable
version: 0.1.0
author: devclaw
description: "Airtable API — databases, records, and bases management"
category: data
tags: [airtable, database, spreadsheet, records, no-code]
requires:
  bins: [curl, jq]
  env: [AIRTABLE_API_KEY, AIRTABLE_BASE_ID]
---
# Airtable

Interact with Airtable using their REST API.

## Setup

1. Get API key: https://airtable.com/account
2. Get Base ID from API docs: https://airtable.com/api
3. Set environment variables:
   ```bash
   export AIRTABLE_API_KEY="patxxx"
   export AIRTABLE_BASE_ID="appxxx"
   ```

## List Records

```bash
# Get all records from table
curl -s "https://api.airtable.com/v0/$AIRTABLE_BASE_ID/TableName" \
  -H "Authorization: Bearer $AIRTABLE_API_KEY" | jq '.records[]'

# With view filter
curl -s "https://api.airtable.com/v0/$AIRTABLE_BASE_ID/TableName?view=Grid%20view" \
  -H "Authorization: Bearer $AIRTABLE_API_KEY" | jq '.records[]'

# With field filter
curl -s "https://api.airtable.com/v0/$AIRTABLE_BASE_ID/TableName?filterByFormula={Status}='Active'" \
  -H "Authorization: Bearer $AIRTABLE_API_KEY" | jq '.records[]'

# Pagination
curl -s "https://api.airtable.com/v0/$AIRTABLE_BASE_ID/TableName?maxRecords=100&pageSize=50" \
  -H "Authorization: Bearer $AIRTABLE_API_KEY" | jq '.records[]'
```

## Get Single Record

```bash
curl -s "https://api.airtable.com/v0/$AIRTABLE_BASE_ID/TableName/recXXXXXX" \
  -H "Authorization: Bearer $AIRTABLE_API_KEY" | jq '.'
```

## Create Record

```bash
curl -s -X POST "https://api.airtable.com/v0/$AIRTABLE_BASE_ID/TableName" \
  -H "Authorization: Bearer $AIRTABLE_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "records": [{
      "fields": {
        "Name": "John Doe",
        "Email": "john@example.com",
        "Status": "Active"
      }
    }]
  }' | jq '.records[0]'
```

## Update Record

```bash
# Update specific fields
curl -s -X PATCH "https://api.airtable.com/v0/$AIRTABLE_BASE_ID/TableName" \
  -H "Authorization: Bearer $AIRTABLE_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "records": [{
      "id": "recXXXXXX",
      "fields": {"Status": "Completed"}
    }]
  }'

# Replace entire record
curl -s -X PUT "https://api.airtable.com/v0/$AIRTABLE_BASE_ID/TableName/recXXXXXX" \
  -H "Authorization: Bearer $AIRTABLE_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{"fields": {"Name": "Updated Name", "Email": "new@email.com"}}'
```

## Delete Record

```bash
curl -s -X DELETE "https://api.airtable.com/v0/$AIRTABLE_BASE_ID/TableName?records[]=recXXXXXX" \
  -H "Authorization: Bearer $AIRTABLE_API_KEY"
```

## Batch Operations

```bash
# Create multiple records (max 10)
curl -s -X POST "https://api.airtable.com/v0/$AIRTABLE_BASE_ID/TableName" \
  -H "Authorization: Bearer $AIRTABLE_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "records": [
      {"fields": {"Name": "User 1"}},
      {"fields": {"Name": "User 2"}}
    ]
  }'

# Update multiple records
curl -s -X PATCH "https://api.airtable.com/v0/$AIRTABLE_BASE_ID/TableName" \
  -H "Authorization: Bearer $AIRTABLE_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "records": [
      {"id": "rec1", "fields": {"Status": "Done"}},
      {"id": "rec2", "fields": {"Status": "Done"}}
    ]
  }'
```

## Filter Formulas

```bash
# Common formulas
# Equal: {Field}='Value'
# Contains: SEARCH('text', {Field})
# Greater than: {Number}>100
# Date: IS_AFTER({Date}, '2025-01-01')
# And: AND({A}='X', {B}='Y')
# Or: OR({Status}='Active', {Status}='Pending')

curl -s "https://api.airtable.com/v0/$AIRTABLE_BASE_ID/TableName?filterByFormula=AND({Status}='Active',{Score}>50)" \
  -H "Authorization: Bearer $AIRTABLE_API_KEY"
```

## Tips

- Personal Access Tokens (patxxx) are recommended over API keys
- Table names with spaces must be URL encoded
- Max 10 records per batch operation
- Rate limit: 5 requests/second
- Use `sort` parameter: `?sort[0][field]=Name&sort[0][direction]=desc`

## Triggers

airtable, airtable api, airtable records, database, spreadsheet
