---
name: docker-compose
description: "Advanced Docker Compose: multi-service, scaling, profiles"
---
# Docker Compose Advanced

Use the **bash** tool for advanced Compose workflows.

## Profiles & Environment
```bash
docker compose --profile debug up -d
docker compose -f docker-compose.yml -f docker-compose.prod.yml up -d
docker compose --env-file .env.staging up -d
```

## Scaling & Health
```bash
docker compose up -d --scale worker=3
docker compose ps --format json | jq '.[] | {Name, State, Health}'
docker compose top
```

## Networking
```bash
docker network ls
docker network inspect <network>
docker compose exec <service> ping <other-service>
```

## Tips
- Use profiles for dev/staging/prod separation
- Use multiple compose files for environment overrides
- Always check health status after scaling
