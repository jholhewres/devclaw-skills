---
name: google-gmail
description: Access Gmail via Google API using the google_api tool
homepage: https://developers.google.com/gmail/api/reference/rest
metadata:
  openclaw:
    requires:
      auth_profile: google
    emoji: "📧"
---

# Google Gmail

Access Gmail messages, labels, and threads using the `google_api` tool with Gmail API endpoints.

## Prerequisites

You need to configure an OAuth profile first:

```bash
auth_profile_add(
  provider="google-gmail",
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

### List Recent Messages

```bash
google_api(
  service="gmail",
  endpoint="/users/me/messages",
  method="GET",
  params={"maxResults": "10"}
)
```

### Search Messages

```bash
google_api(
  service="gmail",
  endpoint="/users/me/messages",
  method="GET",
  params={
    "q": "from:boss@company.com newer_than:7d",
    "maxResults": "20"
  }
)
```

Search operators:
- `from:` - Sender email
- `to:` - Recipient email
- `subject:` - Subject line
- `newer_than:7d` - Messages from last 7 days
- `older_than:1m` - Messages older than 1 month
- `is:unread` - Unread messages
- `has:attachment` - Messages with attachments
- `label:important` - Messages with specific label

### Read a Message

```bash
google_api(
  service="gmail",
  endpoint="/users/me/messages/MESSAGE_ID",
  method="GET",
  params={"format": "full"}
)
```

### Mark as Read

```bash
google_api(
  service="gmail",
  endpoint="/users/me/messages/MESSAGE_ID/modify",
  method="POST",
  body={
    "removeLabelIds": ["UNREAD"]
  }
)
```

### Archive (Remove from Inbox)

```bash
google_api(
  service="gmail",
  endpoint="/users/me/messages/MESSAGE_ID/modify",
  method="POST",
  body={
    "removeLabelIds": ["INBOX"]
  }
)
```

### Move to Inbox

```bash
google_api(
  service="gmail",
  endpoint="/users/me/messages/MESSAGE_ID/modify",
  method="POST",
  body={
    "addLabelIds": ["INBOX"]
  }
)
```

### Apply Label

```bash
google_api(
  service="gmail",
  endpoint="/users/me/messages/MESSAGE_ID/modify",
  method="POST",
  body={
    "addLabelIds": ["Label_1"]
  }
)
```

### List Labels

```bash
google_api(
  service="gmail",
  endpoint="/users/me/labels",
  method="GET"
)
```

## Using Specific Profiles

If you have multiple Google accounts configured:

```bash
google_api(
  service="gmail",
  endpoint="/users/me/messages",
  method="GET",
  profile="work"
)
```

## Common Label IDs

- `INBOX` - Inbox
- `UNREAD` - Unread messages
- `STARRED` - Starred messages
- `IMPORTANT` - Important messages
- `SENT` - Sent messages
- `DRAFT` - Drafts
- `SPAM` - Spam
- `TRASH` - Trash

## Error Handling

If you see "No valid OAuth profile found":
1. Check configured profiles: `auth_profile_list()`
2. Add a profile: `auth_profile_add(provider="google-gmail", name="default", mode="oauth")`
3. Complete the OAuth flow in the WebUI

## API Reference

Full API documentation: https://developers.google.com/gmail/api/reference/rest
