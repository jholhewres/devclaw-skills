---
name: google-calendar
description: Access Google Calendar via Google API using the google_api tool
homepage: https://developers.google.com/calendar/api/v3/reference
metadata:
  openclaw:
    requires:
      auth_profile: google
    emoji: "📅"
---

# Google Calendar

Manage calendars and events using the `google_api` tool with Google Calendar API endpoints.

## Prerequisites

You need to configure an OAuth profile first:

```bash
auth_profile_add(
  provider="google-calendar",
  name="default",
  mode="oauth"
)
```

Or use a generic Google profile:

```bash
auth_profile_add(
  provider="google",
  name="personal",
  mode="oauth"
)
```

## Common Operations

### List Calendars

```bash
google_api(
  service="calendar",
  endpoint="/users/me/calendarList",
  method="GET"
)
```

### List Events (Primary Calendar)

```bash
google_api(
  service="calendar",
  endpoint="/calendars/primary/events",
  method="GET",
  params={
    "timeMin": "2024-01-01T00:00:00Z",
    "timeMax": "2024-01-31T23:59:59Z",
    "maxResults": "10"
  }
)
```

### List Events (Specific Calendar)

```bash
google_api(
  service="calendar",
  endpoint="/calendars/CALENDAR_ID/events",
  method="GET",
  params={
    "timeMin": "2024-01-01T00:00:00Z",
    "maxResults": "20"
  }
)
```

### Search Events

```bash
google_api(
  service="calendar",
  endpoint="/calendars/primary/events",
  method="GET",
  params={
    "q": "meeting",
    "timeMin": "2024-01-01T00:00:00Z"
  }
)
```

### Get Event Details

```bash
google_api(
  service="calendar",
  endpoint="/calendars/primary/events/EVENT_ID",
  method="GET"
)
```

### Create Event (Timed)

```bash
google_api(
  service="calendar",
  endpoint="/calendars/primary/events",
  method="POST",
  body={
    "summary": "Team Meeting",
    "description": "Weekly sync with the engineering team",
    "location": "Conference Room A",
    "start": {
      "dateTime": "2024-01-15T14:00:00",
      "timeZone": "America/New_York"
    },
    "end": {
      "dateTime": "2024-01-15T15:00:00",
      "timeZone": "America/New_York"
    },
    "attendees": [
      {"email": "colleague@company.com"},
      {"email": "manager@company.com"}
    ]
  }
)
```

### Create All-Day Event

```bash
google_api(
  service="calendar",
  endpoint="/calendars/primary/events",
  method="POST",
  body={
    "summary": "Vacation",
    "start": {
      "date": "2024-01-15"
    },
    "end": {
      "date": "2024-01-16"
    }
  }
)
```

### Update Event

```bash
google_api(
  service="calendar",
  endpoint="/calendars/primary/events/EVENT_ID",
  method="PUT",
  body={
    "summary": "Updated Meeting Title",
    "description": "Updated description",
    "start": {
      "dateTime": "2024-01-15T15:00:00",
      "timeZone": "America/New_York"
    },
    "end": {
      "dateTime": "2024-01-15T16:00:00",
      "timeZone": "America/New_York"
    }
  }
)
```

### Partial Update (Patch)

```bash
google_api(
  service="calendar",
  endpoint="/calendars/primary/events/EVENT_ID",
  method="PATCH",
  body={
    "summary": "New Title Only"
  }
)
```

### Delete Event

```bash
google_api(
  service="calendar",
  endpoint="/calendars/primary/events/EVENT_ID",
  method="DELETE"
)
```

## Using Specific Profiles

If you have multiple Google accounts configured:

```bash
google_api(
  service="calendar",
  endpoint="/calendars/primary/events",
  method="GET",
  profile="work"
)
```

## Date/Time Formats

### RFC3339 Format (for timeMin/timeMax params)

- `2024-01-15T00:00:00Z` - UTC time
- `2024-01-15T14:30:00-05:00` - With timezone offset

### Date Format (for all-day events)

- `2024-01-15` - Just the date

### Common Time Zones

- `America/New_York` - Eastern Time
- `America/Chicago` - Central Time
- `America/Denver` - Mountain Time
- `America/Los_Angeles` - Pacific Time
- `Europe/London` - GMT/BST
- `Europe/Paris` - Central European Time
- `Asia/Tokyo` - Japan Time
- `UTC` - UTC

## Query Parameters

- `timeMin` - Start of time range (RFC3339)
- `timeMax` - End of time range (RFC3339)
- `q` - Free text search
- `maxResults` - Maximum number of events (1-2500)
- `orderBy` - Sort order: "startTime" or "updated"
- `singleEvents` - "true" to expand recurring events
- `showDeleted` - "true" to include deleted/cancelled events

## Error Handling

If you see "No valid OAuth profile found":
1. Check configured profiles: `auth_profile_list()`
2. Add a profile: `auth_profile_add(provider="google-calendar", name="default", mode="oauth")`
3. Complete the OAuth flow in the WebUI

## API Reference

Full API documentation: https://developers.google.com/calendar/api/v3/reference
