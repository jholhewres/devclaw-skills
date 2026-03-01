---
name: google-contacts
description: Access Google Contacts via Google People API using the google_api tool
homepage: https://developers.google.com/people/api/rest
metadata:
  openclaw:
    requires:
      auth_profile: google
    emoji: "👥"
---

# Google Contacts

Manage contacts using the `google_api` tool with Google People API endpoints.

## Prerequisites

You need to configure an OAuth profile first:

```bash
auth_profile_add(
  provider="google-contacts",
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

## Common Operations

### List Contacts

```bash
google_api(
  service="people",
  endpoint="/people/me/connections",
  method="GET",
  params={
    "pageSize": "20",
    "personFields": "names,emailAddresses,phoneNumbers"
  }
)
```

### Search Contacts

```bash
google_api(
  service="people",
  endpoint="/people:searchContacts",
  method="GET",
  params={
    "query": "John",
    "readMask": "names,emailAddresses,phoneNumbers"
  }
)
```

### Get Contact

```bash
google_api(
  service="people",
  endpoint":"/PEOPLE_ID",
  method="GET",
  params={
    "personFields": "names,emailAddresses,phoneNumbers,biographies"
  }
)
```

### Create Contact

```bash
google_api(
  service="people",
  endpoint="/people:createContact",
  method="POST",
  body={
    "names": [
      {
        "givenName": "John",
        "familyName": "Doe"
      }
    ],
    "emailAddresses": [
      {
        "value": "john.doe@example.com",
        "type": "work"
      }
    ],
    "phoneNumbers": [
      {
        "value": "+1 555-0123",
        "type": "mobile"
      }
    ]
  }
)
```

### Update Contact

```bash
google_api(
  service="people",
  endpoint="/PEOPLE_ID:updateContact",
  method="PATCH",
  params={
    "updatePersonFields": "emailAddresses,phoneNumbers"
  },
  body={
    "emailAddresses": [
      {
        "value": "new.email@example.com",
        "type": "work"
      }
    ],
    "phoneNumbers": [
      {
        "value": "+1 555-0199",
        "type": "mobile"
      }
    ]
  }
)
```

### Delete Contact

```bash
google_api(
  service="people",
  endpoint":"/PEOPLE_ID:deleteContact",
  method="DELETE"
)
```

### Get Self Profile

```bash
google_api(
  service="people",
  endpoint":"/people/me",
  method="GET",
  params={
    "personFields": "names,emailAddresses,photos"
  }
)
```

## Person Fields

Comma-separated list of fields to return:

- `names` - Full name
- `emailAddresses` - Email addresses
- `phoneNumbers` - Phone numbers
- `addresses` - Physical addresses
- `biographies` - Notes/biography
- `birthdays` - Birthday information
- `organizations` - Work organizations
- `urls` - Website URLs
- `photos` - Profile photos
- `memberships` - Contact group memberships

## Contact Fields

### Names
```json
{
  "givenName": "John",
  "familyName": "Doe",
  "displayName": "John Doe"
}
```

### Email Addresses
```json
{
  "value": "john@example.com",
  "type": "work",
  "formattedType": "Work"
}
```

### Phone Numbers
```json
{
  "value": "+1 555-0123",
  "type": "mobile",
  "formattedType": "Mobile"
}
```

### Addresses
```json
{
  "streetAddress": "123 Main St",
  "city": "Anytown",
  "region": "CA",
  "postalCode": "12345",
  "country": "USA",
  "type": "home"
}
```

## Using Specific Profiles

```bash
google_api(
  service="people",
  endpoint="/people/me/connections",
  method="GET",
  profile="work"
)
```

## Error Handling

If you see "No valid OAuth profile found":
1. Check configured profiles: `auth_profile_list()`
2. Add a profile: `auth_profile_add(provider="google-contacts", name="default", mode="oauth")`
3. Complete the OAuth flow in the WebUI

## Pagination

For large contact lists, use `pageToken`:
```bash
google_api(
  service="people",
  endpoint="/people/me/connections",
  method="GET",
  params={
    "pageSize": "20",
    "pageToken": "NEXT_PAGE_TOKEN",
    "personFields": "names,emailAddresses"
  }
)
```

## API Reference

Full API documentation: https://developers.google.com/people/api/rest

Note: The modern People API is preferred over the older Contacts API. Use `people` service for all contact operations.
