---
name: docker
version: 0.1.0
author: devclaw
description: "Docker management — containers, images, volumes, and networks"
category: devops
tags: [docker, containers, devops, deployment]
requires:
  bins: [docker]
---
# Docker

Manage Docker containers, images, volumes, and networks.

## Containers

```bash
# List running containers
docker ps

# List all containers (including stopped)
docker ps -a

# Run container
docker run -d --name myapp -p 8080:80 nginx

# Run with environment variables
docker run -d -e DATABASE_URL=postgres://... -e API_KEY=xxx myimage

# Run with volume
docker run -d -v /host/path:/container/path myimage

# Run interactively
docker run -it ubuntu bash

# Stop container
docker stop myapp

# Start container
docker start myapp

# Remove container
docker rm myapp

# Force remove running container
docker rm -f myapp
```

## Execute in Container

```bash
# Execute command
docker exec myapp ls /app

# Interactive shell
docker exec -it myapp bash

# Run as specific user
docker exec -u root myapp command

# Copy files
docker cp file.txt myapp:/app/
docker cp myapp:/app/logs/ ./logs/
```

## Logs & Monitoring

```bash
# View logs
docker logs myapp

# Follow logs
docker logs -f myapp

# Last 100 lines
docker logs --tail 100 myapp

# With timestamps
docker logs -t myapp

# Stats (CPU, memory, network)
docker stats myapp

# Inspect container
docker inspect myapp
```

## Images

```bash
# List images
docker images

# Pull image
docker pull nginx:latest
docker pull node:18-alpine

# Build image
docker build -t myapp:latest .

# Build with no cache
docker build --no-cache -t myapp:latest .

# Remove image
docker rmi nginx:latest

# Remove dangling images
docker image prune

# Save/load image
docker save myapp:latest | gzip > myapp.tar.gz
docker load < myapp.tar.gz
```

## Docker Compose

```bash
# Start services
docker compose up -d

# Build and start
docker compose up -d --build

# Stop services
docker compose down

# View logs
docker compose logs -f

# Scale service
docker compose up -d --scale web=3

# Execute in service
docker compose exec web bash
```

## Volumes

```bash
# List volumes
docker volume ls

# Create volume
docker volume create mydata

# Inspect volume
docker volume inspect mydata

# Remove volume
docker volume rm mydata

# Remove unused volumes
docker volume prune
```

## Networks

```bash
# List networks
docker network ls

# Create network
docker network create mynet

# Connect container to network
docker network connect mynet myapp

# Disconnect
docker network disconnect mynet myapp

# Remove network
docker network rm mynet
```

## Cleanup

```bash
# Remove all stopped containers
docker container prune

# Remove all unused images
docker image prune -a

# Remove everything unused
docker system prune -a

# Show disk usage
docker system df
```

## Tips

- Use `--rm` to auto-remove containers after exit
- Use `-d` for detached (background) mode
- Use `-p host:container` for port mapping
- Use `--restart unless-stopped` for auto-restart
- Use `docker compose` for multi-container apps

## Triggers

docker, container, docker container, docker image, docker compose,
run container, docker logs, docker build
