---
name: clickup
version: 0.1.0
author: devclaw
description: "ClickUp API — tasks, lists, spaces, and productivity management"
category: project-management
tags: [clickup, tasks, productivity, project-management]
requires:
  bins: [curl, jq]
  env: [CLICKUP_API_TOKEN]
---
# ClickUp

Interact with ClickUp using their REST API.

## Setup

1. Go to Settings > Apps in ClickUp
2. Generate an API Token
3. Set environment variable:
   ```bash
   export CLICKUP_API_TOKEN="pk_xxx"
   ```

## Get Authorized User & Workspaces

```bash
# Get user info
curl -s "https://api.clickup.com/api/v2/user" \
  -H "Authorization: $CLICKUP_API_TOKEN" | jq '.user'

# Get workspaces (teams)
curl -s "https://api.clickup.com/api/v2/team" \
  -H "Authorization: $CLICKUP_API_TOKEN" | jq '.teams[]'

# Get spaces in workspace
curl -s "https://api.clickup.com/api/v2/team/TEAM_ID/space" \
  -H "Authorization: $CLICKUP_API_TOKEN" | jq '.spaces[]'

# Get lists in space
curl -s "https://api.clickup.com/api/v2/space/SPACE_ID/list" \
  -H "Authorization: $CLICKUP_API_TOKEN" | jq '.lists[]'
```

## Tasks

```bash
# Get tasks from list
curl -s "https://api.clickup.com/api/v2/list/LIST_ID/task" \
  -H "Authorization: $CLICKUP_API_TOKEN" | jq '.tasks[]'

# Get task details
curl -s "https://api.clickup.com/api/v2/task/TASK_ID" \
  -H "Authorization: $CLICKUP_API_TOKEN" | jq '.'

# Get tasks with filters
curl -s "https://api.clickup.com/api/v2/list/LIST_ID/task?status[]=Open&due_date_lt=1735689600000" \
  -H "Authorization: $CLICKUP_API_TOKEN" | jq '.tasks[]'
```

## Create Task

```bash
curl -s -X POST "https://api.clickup.com/api/v2/list/LIST_ID/task" \
  -H "Authorization: $CLICKUP_API_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Implement authentication",
    "description": "Add OAuth2 support",
    "priority": 2,
    "due_date": 1735689600000,
    "assignees": [USER_ID]
  }' | jq '.'
```

## Update Task

```bash
# Update task
curl -s -X PUT "https://api.clickup.com/api/v2/task/TASK_ID" \
  -H "Authorization: $CLICKUP_API_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"name": "Updated task name", "status": "in progress"}'

# Close/complete task
curl -s -X PUT "https://api.clickup.com/api/v2/task/TASK_ID" \
  -H "Authorization: $CLICKUP_API_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"status": "complete"}'
```

## Comments

```bash
# Add comment
curl -s -X POST "https://api.clickup.com/api/v2/task/TASK_ID/comment" \
  -H "Authorization: $CLICKUP_API_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"comment_text": "Great progress!"}'

# Get comments
curl -s "https://api.clickup.com/api/v2/task/TASK_ID/comment" \
  -H "Authorization: $CLICKUP_API_TOKEN" | jq '.comments[]'
```

## Time Tracking

```bash
# Start time entry
curl -s -X POST "https://api.clickup.com/api/v2/team/TEAM_ID/time_entries/start" \
  -H "Authorization: $CLICKUP_API_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"tid": "TASK_ID"}'

# Stop time entry
curl -s -X POST "https://api.clickup.com/api/v2/team/TEAM_ID/time_entries/stop" \
  -H "Authorization: $CLICKUP_API_TOKEN"

# Get time entries
curl -s "https://api.clickup.com/api/v2/team/TEAM_ID/time_entries?task_id=TASK_ID" \
  -H "Authorization: $CLICKUP_API_TOKEN" | jq '.data[]'
```

## Tips

- Priority: 1=Urgent, 2=High, 3=Normal, 4=Low
- Dates are Unix timestamps in milliseconds
- Status names depend on your ClickUp workspace configuration
- Use `include_closed=true` to get completed tasks

## Triggers

clickup, clickup task, create clickup task, clickup api, clickup list
