---
name: nginx
description: "Nginx configuration, reverse proxy, and SSL"
metadata: {"openclaw":{"always":false,"emoji":"🌍"}}
---
# Nginx

Nginx management: configuration, reverse proxy, SSL.

## Setup

```bash
# Check if installed
command -v nginx

# Install — macOS
brew install nginx

# Install — Ubuntu/Debian
sudo apt install nginx
```

## Basic Commands

```bash
# Test configuration (ALWAYS before reload)
sudo nginx -t

# Reload (applies config without downtime)
sudo systemctl reload nginx

# Restart
sudo systemctl restart nginx

# Status
systemctl status nginx

# Version and modules
nginx -V
```

## Logs

```bash
# Access log
sudo tail -f /var/log/nginx/access.log
sudo tail -100 /var/log/nginx/access.log

# Error log
sudo tail -f /var/log/nginx/error.log

# Logs for a specific site (if configured)
sudo tail -f /var/log/nginx/<site>-access.log
```

## Site Configuration

```bash
# List sites
ls /etc/nginx/sites-available/
ls /etc/nginx/sites-enabled/

# Enable site
sudo ln -s /etc/nginx/sites-available/<site> /etc/nginx/sites-enabled/
sudo nginx -t && sudo systemctl reload nginx

# Disable site
sudo rm /etc/nginx/sites-enabled/<site>
sudo nginx -t && sudo systemctl reload nginx
```

## Reverse Proxy (template)

```nginx
server {
    listen 80;
    server_name example.com;

    location / {
        proxy_pass http://localhost:3000;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
```

## SSL with Certbot

```bash
# Install certbot
sudo apt install certbot python3-certbot-nginx

# Get certificate (modifies nginx config automatically)
sudo certbot --nginx -d example.com

# Renew
sudo certbot renew --dry-run
sudo certbot renew
```

## WebSocket Proxy

```nginx
location /ws {
    proxy_pass http://localhost:8080;
    proxy_http_version 1.1;
    proxy_set_header Upgrade $http_upgrade;
    proxy_set_header Connection "upgrade";
    proxy_set_header Host $host;
}
```

## Tips

- **Always** run `nginx -t` before reload/restart
- Use `sites-available` + symlink to `sites-enabled`
- For debug: `error_log /var/log/nginx/debug.log debug;`
- Rate limiting: `limit_req_zone` in the `http` section
- Gzip: enable in `/etc/nginx/nginx.conf` for performance
