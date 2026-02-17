---
name: systemd
description: "Manage Linux services with systemctl and journalctl"
---
# systemd

Use **bash** and **ssh** tools to manage Linux services.

## Services
```bash
systemctl status <service>
sudo systemctl start|stop|restart|reload <service>
sudo systemctl enable --now <service>
systemctl list-units --type=service --state=running
systemctl list-units --type=service --state=failed
```

## Logs
```bash
journalctl -u <service> --no-pager -n 50
journalctl -u <service> -f
journalctl -u <service> --since "1 hour ago"
journalctl -u <service> -p err
journalctl -b
```

## Create Service
Use **write_file** to create /etc/systemd/system/<name>.service, then:
```bash
sudo systemctl daemon-reload
sudo systemctl enable --now <name>
```

## Tips
- Always daemon-reload after modifying unit files
- Use journalctl -f for real-time log following
