---
name: linear
version: 0.1.0
author: devclaw
description: "Linear API — issues, projects, and workflow management"
category: project-management
tags: [linear, issues, tickets, agile, kanban]
requires:
  bins: [curl, jq]
  env: [LINEAR_API_KEY]
---
# Linear

Interact with Linear using their GraphQL API.

## Setup

1. Go to Settings > API in Linear
2. Create a Personal API Key
3. Set environment variable:
   ```bash
   export LINEAR_API_KEY="lin_api_xxx"
   ```

## GraphQL Endpoint

```bash
LINEAR_API="https://api.linear.app/graphql"
AUTH_HEADER="Authorization: Bearer $LINEAR_API_KEY"
```

## List Issues

```bash
# Get my issues
curl -s -X POST "$LINEAR_API" \
  -H "$AUTH_HEADER" \
  -H "Content-Type: application/json" \
  -d '{"query": "{ issues(filter: { assignee: { isMe: { eq: true } } }) { nodes { id identifier title priority state { name } } } }"}' | jq '.data.issues.nodes'

# Get issues by team
curl -s -X POST "$LINEAR_API" \
  -H "$AUTH_HEADER" \
  -H "Content-Type: application/json" \
  -d '{"query": "{ issues(filter: { team: { key: { eq: \"ENG\" } } }) { nodes { id identifier title state { name } } } }"}' | jq '.data.issues.nodes'

# Search issues
curl -s -X POST "$LINEAR_API" \
  -H "$AUTH_HEADER" \
  -H "Content-Type: application/json" \
  -d '{"query": "{ issues(filter: { search: { eq: \"bug\" } }) { nodes { identifier title } } }"}' | jq '.data.issues.nodes'
```

## Create Issue

```bash
curl -s -X POST "$LINEAR_API" \
  -H "$AUTH_HEADER" \
  -H "Content-Type: application/json" \
  -d '{
    "query": "mutation { issueCreate(input: { title: \"New bug report\", description: \"Description here\", teamId: \"TEAM_ID\", priority: 2 }) { success issue { id identifier } } }"
  }' | jq '.data.issueCreate'
```

## Update Issue

```bash
# Update status
curl -s -X POST "$LINEAR_API" \
  -H "$AUTH_HEADER" \
  -H "Content-Type: application/json" \
  -d '{
    "query": "mutation { issueUpdate(id: \"ISSUE_ID\", input: { stateId: \"STATE_ID\" }) { success } }"
  }'

# Add comment
curl -s -X POST "$LINEAR_API" \
  -H "$AUTH_HEADER" \
  -H "Content-Type: application/json" \
  -d '{
    "query": "mutation { commentCreate(input: { issueId: \"ISSUE_ID\", body: \"Progress update\" }) { success } }"
  }'
```

## Get Teams & Projects

```bash
# List teams
curl -s -X POST "$LINEAR_API" \
  -H "$AUTH_HEADER" \
  -H "Content-Type: application/json" \
  -d '{"query": "{ teams { nodes { id key name } } }"}' | jq '.data.teams.nodes'

# List projects
curl -s -X POST "$LINEAR_API" \
  -H "$AUTH_HEADER" \
  -H "Content-Type: application/json" \
  -d '{"query": "{ projects { nodes { id name status { name } } } }"}' | jq '.data.projects.nodes'

# Get workflow states
curl -s -X POST "$LINEAR_API" \
  -H "$AUTH_HEADER" \
  -H "Content-Type: application/json" \
  -d '{"query": "{ team(key: \"ENG\") { states { nodes { id name type } } } }"}' | jq '.data.team.states.nodes'
```

## Get Viewer Info

```bash
curl -s -X POST "$LINEAR_API" \
  -H "$AUTH_HEADER" \
  -H "Content-Type: application/json" \
  -d '{"query": "{ viewer { id name email } }"}' | jq '.data.viewer'
```

## Tips

- Priority: 0=No priority, 1=Urgent, 2=High, 3=Normal, 4=Low
- Use `identifier` (e.g., "ENG-123") for human-readable IDs
- GraphQL allows fetching exactly the fields you need
- Use `filter` for complex queries

## Triggers

linear, linear issue, create linear issue, linear ticket, linear api
