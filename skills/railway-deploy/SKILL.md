---
name: railway-deploy
version: 0.1.0
author: devclaw
description: "Railway deployment — deploy apps, databases, and services"
category: devops
tags: [railway, deploy, paas, hosting, cloud]
requires:
  bins: [railway]
---
# Railway Deploy

Deploy applications and services to Railway.

## Setup

```bash
# Install Railway CLI
npm install -g @railway/cli

# Login
railway login

# Or with browser
railway login --browserless
```

## Projects

```bash
# Create new project
railway init

# Link existing project
railway link PROJECT_ID

# List projects
railway list

# View project info
railway status
```

## Deploy

```bash
# Deploy from current directory
railway up

# Deploy with name
railway up --service myapp

# Deploy from subdirectory
railway up --service api ./api

# View deployment logs
railway logs

# Follow logs
railway logs -f

# View deployment status
railway status
```

## Services

```bash
# List services
railway service list

# Create service
railway service create myservice

# Delete service
railway service delete myservice

# Connect to service
railway connect myservice
```

## Databases

```bash
# Add PostgreSQL
railway add --plugin postgresql

# Add MySQL
railway add --plugin mysql

# Add Redis
railway add --plugin redis

# Add MongoDB
railway add --plugin mongodb

# Get database URL
railway variables --service postgres

# Connect to database
railway connect postgres
```

## Variables

```bash
# List variables
railway variables

# Set variable
railway variables set KEY=value

# Set multiple
railway variables set KEY1=value1 KEY2=value2

# Get variable value
railway variables get KEY

# Delete variable
railway variables del KEY
```

## Domains

```bash
# Generate domain
railway domain

# Add custom domain
railway domain add myapp.example.com

# Remove domain
railway domain remove myapp.example.com
```

## Run Commands

```bash
# Run one-off command
railway run npm run migrate

# Run in specific service
railway run --service api npm start

# SSH into container
railway shell
```

## CLI Shortcuts

```bash
# Open Railway dashboard
railway open

# View docs
railway docs

# Logout
railway logout

# Check version
railway --version
```

## railway.toml Example

```toml
[build]
builder = "nixpacks"

[deploy]
startCommand = "npm start"
healthcheckPath = "/health"
healthcheckTimeout = 300
restartPolicyType = "on_failure"
restartPolicyMaxRetries = 3
```

## Tips

- Use `railway up` for quick deploys
- Use `railway run` for one-off commands
- Connect databases with `railway connect`
- Check logs with `railway logs -f`
- Variables are auto-injected at runtime

## Triggers

railway, deploy to railway, railway deploy, railway app,
railway database, railway cli
