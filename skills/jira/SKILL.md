---
name: jira
version: 0.1.0
author: devclaw
description: "Jira Cloud integration — issues, transitions, worklogs, sprints"
category: project-management
tags: [jira, atlassian, issues, tickets, agile, sprints]
requires:
  bins: [curl, jq]
  env: [JIRA_URL, JIRA_EMAIL, JIRA_API_TOKEN]
---
# Jira

Manage Jira issues, transitions, and worklogs via the Jira Cloud REST API.

## Setup

1. Get API token: https://id.atlassian.com/manage-profile/security/api-tokens
2. Click "Create API Token"
3. Set environment variables:
   ```bash
   export JIRA_URL="https://your-domain.atlassian.net"
   export JIRA_EMAIL="you@example.com"
   export JIRA_API_TOKEN="your-api-token"
   ```

## Authentication

```bash
# Create auth header (use in all requests)
AUTH=$(echo -n "$JIRA_EMAIL:$JIRA_API_TOKEN" | base64)
```

## Search Issues (JQL)

```bash
# Search by text
curl -s -X POST "$JIRA_URL/rest/api/3/search" \
  -H "Authorization: Basic $AUTH" \
  -H "Content-Type: application/json" \
  -d '{"jql": "text ~ \"timeout\"", "maxResults": 20}' | jq '.issues[] | {key, summary: .fields.summary, status: .fields.status.name}'

# My open issues
curl -s "$JIRA_URL/rest/api/3/search?jql=assignee=currentUser()+AND+resolution=Unresolved&maxResults=50" \
  -H "Authorization: Basic $AUTH" | jq '.issues[] | {key, summary: .fields.summary}'

# By project
curl -s "$JIRA_URL/rest/api/3/search?jql=project=PROJ&maxResults=20" \
  -H "Authorization: Basic $AUTH" | jq '.issues[] | {key, summary}'
```

## Get Issue Details

```bash
# View issue
curl -s "$JIRA_URL/rest/api/3/issue/PROJ-123" \
  -H "Authorization: Basic $AUTH" | jq '{
    key,
    summary: .fields.summary,
    description: .fields.description,
    status: .fields.status.name,
    assignee: .fields.assignee.displayName,
    priority: .fields.priority.name
  }'
```

## Create Issue

```bash
curl -s -X POST "$JIRA_URL/rest/api/3/issue" \
  -H "Authorization: Basic $AUTH" \
  -H "Content-Type: application/json" \
  -d '{
    "fields": {
      "project": {"key": "PROJ"},
      "summary": "Bug: Login timeout",
      "description": "Users being logged out after 5 minutes",
      "issuetype": {"name": "Bug"},
      "priority": {"name": "High"}
    }
  }' | jq '.key'
```

## Update Issue

```bash
# Add comment
curl -s -X POST "$JIRA_URL/rest/api/3/issue/PROJ-123/comment" \
  -H "Authorization: Basic $AUTH" \
  -H "Content-Type: application/json" \
  -d '{"body": "Deployed to staging"}'

# Assign issue
curl -s -X PUT "$JIRA_URL/rest/api/3/issue/PROJ-123" \
  -H "Authorization: Basic $AUTH" \
  -H "Content-Type: application/json" \
  -d '{"fields": {"assignee": {"accountId": "USER_ACCOUNT_ID"}}}'
```

## Transitions (Change Status)

```bash
# List available transitions
curl -s "$JIRA_URL/rest/api/3/issue/PROJ-123/transitions" \
  -H "Authorization: Basic $AUTH" | jq '.transitions[] | {id, name}'

# Perform transition
curl -s -X POST "$JIRA_URL/rest/api/3/issue/PROJ-123/transitions" \
  -H "Authorization: Basic $AUTH" \
  -H "Content-Type: application/json" \
  -d '{"transition": {"id": "31"}}'
```

## Worklogs

```bash
# Log work
curl -s -X POST "$JIRA_URL/rest/api/3/issue/PROJ-123/worklog" \
  -H "Authorization: Basic $AUTH" \
  -H "Content-Type: application/json" \
  -d '{"timeSpent": "2h 30m", "started": "2025-01-15T09:00:00.000+0000"}'

# Get worklogs
curl -s "$JIRA_URL/rest/api/3/issue/PROJ-123/worklog" \
  -H "Authorization: Basic $AUTH" | jq '.worklogs[] | {author: .author.displayName, timeSpent, started}'
```

## Tips

- Use JQL for powerful searches: `project = PROJ AND status = "In Progress"`
- Transition IDs vary by project - always list first
- Max results default is 50, max is 100
- Use `resolution = Unresolved` for open issues

## Triggers

jira, issue, ticket, create issue, jira ticket, sprint, backlog,
transition, log work, jql
