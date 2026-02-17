---
name: todoist
version: 0.1.0
author: devclaw
description: "Todoist task management — create, list, complete, and organize tasks"
category: productivity
tags: [todoist, tasks, todo, productivity, project-management]
requires:
  env: [TODOIST_API_TOKEN]
---
# Todoist

Manage tasks and projects in Todoist via the REST API v2.

## Setup

1. Get your API token: Todoist Settings → Integrations → Developer → API Token
2. Store in vault: `copilot config vault-set TODOIST_API_TOKEN=your_token`

## List tasks

```bash
# All active tasks
curl -s "https://api.todoist.com/rest/v2/tasks" \
  -H "Authorization: Bearer $TODOIST_API_TOKEN" | jq '.[] | {id, content, due: .due.string, priority, project_id}'

# Filter by project
curl -s "https://api.todoist.com/rest/v2/tasks?project_id=PROJECT_ID" \
  -H "Authorization: Bearer $TODOIST_API_TOKEN" | jq '.[] | {id, content, due: .due.string}'

# Filter with Todoist filter syntax
curl -s -G "https://api.todoist.com/rest/v2/tasks" \
  --data-urlencode "filter=today | overdue" \
  -H "Authorization: Bearer $TODOIST_API_TOKEN" | jq '.[] | {id, content, due: .due.string}'
```

## Create a task

```bash
# Simple task
curl -s -X POST "https://api.todoist.com/rest/v2/tasks" \
  -H "Authorization: Bearer $TODOIST_API_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "content": "Task title",
    "due_string": "tomorrow at 10am",
    "priority": 3
  }' | jq '{id, content, due: .due.string, url}'

# Task in a specific project with labels
curl -s -X POST "https://api.todoist.com/rest/v2/tasks" \
  -H "Authorization: Bearer $TODOIST_API_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "content": "Buy groceries",
    "description": "Milk, bread, eggs",
    "project_id": "PROJECT_ID",
    "due_string": "today",
    "labels": ["shopping"],
    "priority": 2
  }' | jq '{id, content, url}'
```

## Complete a task

```bash
curl -s -X POST "https://api.todoist.com/rest/v2/tasks/TASK_ID/close" \
  -H "Authorization: Bearer $TODOIST_API_TOKEN"
```

## Update a task

```bash
curl -s -X POST "https://api.todoist.com/rest/v2/tasks/TASK_ID" \
  -H "Authorization: Bearer $TODOIST_API_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"content": "Updated title", "due_string": "next monday"}'
```

## Delete a task

```bash
curl -s -X DELETE "https://api.todoist.com/rest/v2/tasks/TASK_ID" \
  -H "Authorization: Bearer $TODOIST_API_TOKEN"
```

## Projects

```bash
# List projects
curl -s "https://api.todoist.com/rest/v2/projects" \
  -H "Authorization: Bearer $TODOIST_API_TOKEN" | jq '.[] | {id, name, color}'

# Create project
curl -s -X POST "https://api.todoist.com/rest/v2/projects" \
  -H "Authorization: Bearer $TODOIST_API_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"name": "New Project", "color": "blue"}' | jq '{id, name, url}'
```

## Priority levels

| Priority | Todoist value | Label |
|----------|--------------|-------|
| Urgent | 4 | p1 (red) |
| High | 3 | p2 (orange) |
| Medium | 2 | p3 (yellow) |
| Normal | 1 | p4 (no color) |

## Due date natural language

Todoist supports natural language via `due_string`:
- `today`, `tomorrow`, `next monday`
- `every day at 9am`, `every weekday`
- `Jan 15`, `2026-03-01`
- `in 3 days`, `in 2 hours`

## Tips

- Use `due_string` with natural language — Todoist parses it automatically.
- Priority 4 is the highest (urgent/red), 1 is the lowest (normal).
- List projects first to get project IDs before creating tasks in them.
- Use filters like `today | overdue` to see what needs attention.
- Confirm before deleting tasks.
- Use labels for cross-project categorization.

## Triggers

todoist, add task, create task, my tasks, complete task,
adicionar tarefa, minhas tarefas, marcar como feito, to-do
