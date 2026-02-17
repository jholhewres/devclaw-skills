---
name: pm2
description: "Process management with PM2"
---
# PM2

Use the **bash** tool for PM2 process management.

## Manage
```bash
pm2 start <script> --name <name>
pm2 list
pm2 stop|restart|delete <name|id>
pm2 reload <name>
```

## Logs & Monitor
```bash
pm2 logs <name> --lines 100
pm2 show <name>
pm2 flush
```

## Persist
```bash
pm2 startup systemd
pm2 save
pm2 resurrect
```

## Tips
- pm2 save after any change to persist across reboots
- pm2 startup + pm2 save to survive reboots
- Works with any binary, not just Node.js
