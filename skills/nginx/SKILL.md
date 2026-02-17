---
name: nginx
description: "Configurar e gerenciar Nginx: sites, SSL, proxy reverso"
metadata: {"openclaw":{"always":false,"emoji":"🌍"}}
---
# Nginx

Gerenciamento de Nginx: configuração, proxy reverso, SSL.

## Comandos Básicos

```bash
# Testar configuração (SEMPRE antes de reload)
sudo nginx -t

# Reload (aplica config sem downtime)
sudo systemctl reload nginx

# Restart
sudo systemctl restart nginx

# Status
systemctl status nginx

# Versão e módulos
nginx -V
```

## Logs

```bash
# Access log
sudo tail -f /var/log/nginx/access.log
sudo tail -100 /var/log/nginx/access.log

# Error log
sudo tail -f /var/log/nginx/error.log

# Logs de um site específico (se configurado)
sudo tail -f /var/log/nginx/<site>-access.log
```

## Configuração de Sites

```bash
# Listar sites
ls /etc/nginx/sites-available/
ls /etc/nginx/sites-enabled/

# Ativar site
sudo ln -s /etc/nginx/sites-available/<site> /etc/nginx/sites-enabled/
sudo nginx -t && sudo systemctl reload nginx

# Desativar site
sudo rm /etc/nginx/sites-enabled/<site>
sudo nginx -t && sudo systemctl reload nginx
```

## Proxy Reverso (template)

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

## SSL com Certbot

```bash
# Instalar certbot
sudo apt install certbot python3-certbot-nginx

# Obter certificado (modifica nginx config automaticamente)
sudo certbot --nginx -d example.com

# Renovar
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

- **Sempre** rode `nginx -t` antes de reload/restart
- Use `sites-available` + symlink para `sites-enabled`
- Para debug: `error_log /var/log/nginx/debug.log debug;`
- Rate limiting: `limit_req_zone` na seção `http`
- Gzip: ative em `/etc/nginx/nginx.conf` para performance
