---
name: google-tasks
description: Access Google Tasks via Google API using the google_api tool
homepage: https://developers.google.com/tasks/reference/rest
metadata:
  openclaw:
    requires:
      auth_profile: google
    emoji: "✅"
---

# Google Tasks

Manage tasks and task lists using the `google_api` tool with Google Tasks API endpoints.

## Prerequisites

You need to configure an OAuth profile first:

```bash
auth_profile_add(
  provider="google-tasks",
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

### List Task Lists

```bash
google_api(
  service="tasks",
  endpoint="/users/@me/lists",
  method="GET"
)
```

### Get Task List

```bash
google_api(
  service="tasks",
  endpoint="/users/@me/lists/TASK_LIST_ID",
  method="GET"
)
```

### Create Task List

```bash
google_api(
  service="tasks",
  endpoint="/users/@me/lists",
  method="POST",
  body={
    "title": "My Project Tasks"
  }
)
```

### Delete Task List

```bash
google_api(
  service="tasks",
  endpoint="/users/@me/lists/TASK_LIST_ID",
  method="DELETE"
)
```

### List Tasks

```bash
google_api(
  service="tasks",
  endpoint="/lists/TASK_LIST_ID/tasks",
  method="GET"
)
```

### List Completed Tasks

```bash
google_api(
  service="tasks",
  endpoint="/lists/TASK_LIST_ID/tasks",
  method="GET",
  params={
    "showCompleted": "true",
    "showHidden": "true"
  }
)
```

### Get Task

```bash
google_api(
  service="tasks",
  endpoint="/lists/TASK_LIST_ID/tasks/TASK_ID",
  method="GET"
)
```

### Create Task

```bash
google_api(
  service="tasks",
  endpoint="/lists/TASK_LIST_ID/tasks",
  method="POST",
  body={
    "title": "Buy groceries",
    "notes": "Milk, eggs, bread",
    "due": "2024-01-15T00:00:00.000Z"
  }
)
```

### Create Task at Top

```bash
google_api(
  service="tasks",
  endpoint="/lists/TASK_LIST_ID/tasks",
  method="POST",
  body={
    "title": "Urgent task"
  },
  params={
    "previous": ""
  }
)
```

### Update Task

```bash
google_api(
  service="tasks",
  endpoint="/lists/TASK_LIST_ID/tasks/TASK_ID",
  method="PUT",
  body={
    "id": "TASK_ID",
    "title": "Buy groceries and cleaning supplies",
    "notes": "Updated list",
    "due": "2024-01-16T00:00:00.000Z",
    "status": "needsAction"
  }
)
```

### Complete Task

```bash
google_api(
  service="tasks",
  endpoint="/lists/TASK_LIST_ID/tasks/TASK_ID",
  method="PUT",
  body={
    "id": "TASK_ID",
    "title": "Task title",
    "status": "completed",
    "completed": "2024-01-15T12:00:00.000Z"
  }
)
```

### Mark Task Incomplete

```bash
google_api(
  service="tasks",
  endpoint="/lists/TASK_LIST_ID/tasks/TASK_ID",
  method="PUT",
  body={
    "id": "TASK_ID",
    "title": "Task title",
    "status": "needsAction",
    "completed": null
  }
)
```

### Move Task

```bash
google_api(
  service="tasks",
  endpoint="/lists/TASK_LIST_ID/tasks/TASK_ID/move",
  method="POST",
  params={
    "previous": "PREVIOUS_TASK_ID"
  }
)
```

### Delete Task

```bash
google_api(
  service="tasks",
  endpoint="/lists/TASK_LIST_ID/tasks/TASK_ID",
  method="DELETE"
)
```

## Task Statuses

- `needsAction` - Task is not complete
- `completed` - Task is complete

## Date Format

Due dates should be in RFC 3339 format:
- `2024-01-15T00:00:00.000Z` - Date with time
- `2024-01-15` - Date only (will have time set to midnight)

## Using Specific Profiles

```bash
google_api(
  service="tasks",
  endpoint="/users/@me/lists",
  method="GET",
  profile="work"
)
```

## Error Handling

If you see "No valid OAuth profile found":
1. Check configured profiles: `auth_profile_list()`
2. Add a profile: `auth_profile_add(provider="google-tasks", name="default", mode="oauth")`
3. Complete the OAuth flow in the WebUI

## Default Task List

The default task list ID is typically `@default`:
```bash
google_api(
  service="tasks",
  endpoint="/lists/@default/tasks",
  method="GET"
)
```

## API Reference

Full API documentation: https://developers.google.com/tasks/reference/rest
