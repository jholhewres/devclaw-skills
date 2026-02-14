---
name: health
version: 0.1.0
author: goclaw
description: "Health and system monitoring — check services, uptime, and alerts"
category: monitoring
tags: [health, monitoring, uptime, alerts, status]
requires:
  bins: [curl, jq]
---
# Health

Monitor system health, services, and uptime.

## HTTP Health Checks

```bash
# Basic health check
curl -s -o /dev/null -w "%{http_code}" https://example.com

# With response time
curl -s -o /dev/null -w "HTTP: %{http_code}, Time: %{time_total}s\n" https://example.com

# Check multiple endpoints
for url in "https://api.example.com/health" "https://app.example.com"; do
  status=$(curl -s -o /dev/null -w "%{http_code}" "$url")
  echo "$url: $status"
done

# Detailed check
curl -s -w "\n\nDNS: %{time_namelookup}s\nConnect: %{time_connect}s\nTTFB: %{time_starttransfer}s\nTotal: %{time_total}s\nHTTP: %{http_code}\n" -o /dev/null https://example.com
```

## Service Monitoring

```bash
# Check if port is open
nc -zv example.com 443

# Check with timeout
timeout 5 bash -c "echo > /dev/tcp/example.com/443" && echo "UP" || echo "DOWN"

# Check systemd service
systemctl is-active nginx

# Check service status
systemctl status nginx --no-pager

# Check Docker container
docker inspect -f '{{.State.Status}}' container_name
```

## System Health

```bash
# CPU usage
top -bn1 | grep "Cpu(s)" | awk '{print $2}' | cut -d'%' -f1

# Memory usage
free -h

# Disk usage
df -h

# Load average
uptime

# Check running processes
ps aux --sort=-%mem | head -10

# Check open files
lsof | wc -l

# Network connections
netstat -an | grep ESTABLISHED | wc -l
```

## SSL Certificate Check

```bash
# Check expiration
echo | openssl s_client -connect example.com:443 2>/dev/null | openssl x509 -noout -dates

# Days until expiration
echo | openssl s_client -connect example.com:443 2>/dev/null | openssl x509 -noout -enddate | cut -d= -f2 | xargs -I {} date -d {} +%s | xargs -I {} echo $(( ({} - $(date +%s)) / 86400 )) days

# Certificate info
echo | openssl s_client -connect example.com:443 2>/dev/null | openssl x509 -noout -text | grep -E "Issuer:|Subject:|DNS:"
```

## Uptime Monitoring

### UptimeRobot API

```bash
# Setup
export UPTIMEROBOT_KEY="xxx"

# Get monitors
curl -s -X POST "https://api.uptimerobot.com/v2/getMonitors" \
  -H "Cache-Control: no-cache" \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -d "api_key=$UPTIMEROBOT_KEY&format=json" | jq '.monitors[]'

# Create monitor
curl -s -X POST "https://api.uptimerobot.com/v2/newMonitor" \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -d "api_key=$UPTIMEROBOT_KEY&friendly_name=MySite&url=https://example.com&type=1" | jq '.'

# Get uptime
curl -s -X POST "https://api.uptimerobot.com/v2/getMonitors" \
  -d "api_key=$UPTIMEROBOT_KEY&monitors=MONITOR_ID&custom_uptime_ratios=30" | jq '.monitors[0].custom_uptime_ratio'
```

### Healthchecks.io

```bash
# Setup
export HEALTHCHECKS_UUID="xxx"

# Ping (mark as alive)
curl -s "https://hc-ping.com/$HEALTHCHECKS_UUID"

# Report failure
curl -s "https://hc-ping.com/$HEALTHCHECKS_UUID/fail"

# Report start
curl -s "https://hc-ping.com/$HEALTHCHECKS_UUID/start"

# With execution time
curl -s "https://hc-ping.com/$HEALTHCHECKS_UUID/$?rid=$(uuidgen)"
```

## Alerting

### Slack Alert

```bash
send_alert() {
  local message="$1"
  curl -s -X POST "$SLACK_WEBHOOK_URL" \
    -H "Content-Type: application/json" \
    -d "{\"text\": \"🚨 Health Alert: $message\"}"
}

# Usage
curl -s -o /dev/null -w "%{http_code}" https://example.com | grep -v "200\|301\|302" && send_alert "Site down!"
```

### Email Alert

```bash
send_email() {
  local subject="$1"
  local body="$2"
  curl -s -X POST "https://api.resend.com/emails" \
    -H "Authorization: Bearer $RESEND_API_KEY" \
    -H "Content-Type: application/json" \
    -d "{\"from\": \"alerts@example.com\", \"to\": \"admin@example.com\", \"subject\": \"$subject\", \"text\": \"$body\"}"
}
```

## Monitoring Script

```bash
#!/bin/bash
# health_check.sh

check_url() {
  local url=$1
  local status=$(curl -s -o /dev/null -w "%{http_code}" --max-time 10 "$url")
  if [[ "$status" =~ ^2[0-9][0-9]$ ]]; then
    echo "✅ $url - OK ($status)"
  else
    echo "❌ $url - FAIL ($status)"
    # Send alert here
  fi
}

check_url "https://example.com"
check_url "https://api.example.com/health"
```

## Tips

- Set appropriate timeouts for checks
- Check from multiple locations if possible
- Monitor response time, not just status
- Set up alerts for certificate expiration
- Log check results for trend analysis

## Triggers

health check, uptime, monitor, status check, service health,
ssl check, system monitoring, health monitoring
