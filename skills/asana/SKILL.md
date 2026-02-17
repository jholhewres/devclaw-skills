---
name: asana
version: 0.1.0
author: devclaw
description: "Asana API — tasks, projects, and team collaboration"
category: project-management
tags: [asana, tasks, projects, collaboration, team]
requires:
  bins: [curl, jq]
  env: [ASANA_ACCESS_TOKEN]
---
# Asana

Interact with Asana using their REST API.

## Setup

1. Go to https://app.asana.com/app/asana/-/account_management/apps
2. Create a Personal Access Token
3. Set environment variable:
   ```bash
   export ASANA_ACCESS_TOKEN="x_xxx"
   ```

## List Workspaces & Projects

```bash
# List workspaces
curl -s "https://app.asana.com/api/1.0/workspaces" \
  -H "Authorization: Bearer $ASANA_ACCESS_TOKEN" | jq '.data[]'

# List projects in workspace
curl -s "https://app.asana.com/api/1.0/workspaces/WORKSPACE_ID/projects" \
  -H "Authorization: Bearer $ASANA_ACCESS_TOKEN" | jq '.data[]'

# List all projects
curl -s "https://app.asana.com/api/1.0/projects" \
  -H "Authorization: Bearer $ASANA_ACCESS_TOKEN" | jq '.data[] | {gid, name}'
```

## Tasks

```bash
# List tasks in project
curl -s "https://app.asana.com/api/1.0/projects/PROJECT_ID/tasks" \
  -H "Authorization: Bearer $ASANA_ACCESS_TOKEN" | jq '.data[]'

# Get task details
curl -s "https://app.asana.com/api/1.0/tasks/TASK_ID" \
  -H "Authorization: Bearer $ASANA_ACCESS_TOKEN" | jq '.data'

# My tasks
curl -s "https://app.asana.com/api/1.0/user/me/tasks?opt_fields=name,completed,due_on" \
  -H "Authorization: Bearer $ASANA_ACCESS_TOKEN" | jq '.data[]'
```

## Create Task

```bash
curl -s -X POST "https://app.asana.com/api/1.0/tasks" \
  -H "Authorization: Bearer $ASANA_ACCESS_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "data": {
      "name": "Review pull request",
      "notes": "Check the authentication module changes",
      "projects": ["PROJECT_ID"],
      "assignee": "me",
      "due_on": "2025-01-20"
    }
  }' | jq '.data'
```

## Update Task

```bash
# Complete task
curl -s -X PUT "https://app.asana.com/api/1.0/tasks/TASK_ID" \
  -H "Authorization: Bearer $ASANA_ACCESS_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"data": {"completed": true}}'

# Update task name/notes
curl -s -X PUT "https://app.asana.com/api/1.0/tasks/TASK_ID" \
  -H "Authorization: Bearer $ASANA_ACCESS_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"data": {"name": "Updated task name"}}'

# Add comment
curl -s -X POST "https://app.asana.com/api/1.0/tasks/TASK_ID/stories" \
  -H "Authorization: Bearer $ASANA_ACCESS_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"data": {"text": "Progress update: completed review"}}'
```

## Subtasks

```bash
# Create subtask
curl -s -X POST "https://app.asana.com/api/1.0/tasks/TASK_ID/subtasks" \
  -H "Authorization: Bearer $ASANA_ACCESS_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"data": {"name": "Subtask item"}}'

# List subtasks
curl -s "https://app.asana.com/api/1.0/tasks/TASK_ID/subtasks" \
  -H "Authorization: Bearer $ASANA_ACCESS_TOKEN" | jq '.data[]'
```

## Search

```bash
# Search tasks
curl -s "https://app.asana.com/api/1.0/workspaces/WORKSPACE_ID/tasks/search?text=bug" \
  -H "Authorization: Bearer $ASANA_ACCESS_TOKEN" | jq '.data[]'
```

## Tips

- Use `opt_fields` to request specific fields: `?opt_fields=name,completed,due_on,assignee`
- Use `opt_expand` for nested resources
- GIDs are string IDs used throughout Asana API
- Use `assignee: "me"` to assign to yourself

## Triggers

asana, asana task, create asana task, asana project, asana api
