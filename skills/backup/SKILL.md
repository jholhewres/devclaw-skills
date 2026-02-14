---
name: backup
version: 0.1.0
author: goclaw
description: "Backup and restore — files, databases, and system backups"
category: system
tags: [backup, restore, archive, recovery, sync]
requires:
  bins: [rsync, tar]
---
# Backup

Backup and restore files, databases, and systems.

## File Backups

### rsync (Incremental)

```bash
# Basic sync
rsync -av /source/ /backup/

# With progress
rsync -av --progress /source/ /backup/

# Exclude patterns
rsync -av --exclude 'node_modules' --exclude '.git' --exclude '*.log' /source/ /backup/

# Remote backup
rsync -avz -e ssh /source/ user@remote:/backup/

# Delete files not in source (mirror)
rsync -av --delete /source/ /backup/

# Dry run
rsync -av --dry-run /source/ /backup/

# Backup with date
rsync -av --backup --backup-dir=/backup/$(date +%Y%m%d) /source/ /backup/current/
```

### tar Archives

```bash
# Create archive
tar -czvf backup.tar.gz /path/to/backup/

# With exclude
tar -czvf backup.tar.gz --exclude='*.log' --exclude='node_modules' /path/

# Split large archives
tar -czvf - /path/ | split -b 1G - backup.tar.gz.

# Verify archive
tar -tzvf backup.tar.gz

# Extract
tar -xzvf backup.tar.gz -C /destination/

# Incremental tar
tar --create --file=backup.tar --listed-incremental=backup.snar /path/
```

## Database Backups

### PostgreSQL

```bash
# Full backup
pg_dump -h localhost -U user dbname > backup.sql

# Compressed
pg_dump -h localhost -U user dbname | gzip > backup.sql.gz

# All databases
pg_dumpall -U postgres > all_databases.sql

# Custom format (better for large DBs)
pg_dump -h localhost -U user -Fc dbname > backup.dump

# Restore from SQL
psql -h localhost -U user dbname < backup.sql

# Restore from custom format
pg_restore -h localhost -U user -d dbname backup.dump
```

### MySQL

```bash
# Single database
mysqldump -u user -p dbname > backup.sql

# All databases
mysqldump -u user -p --all-databases > all_databases.sql

# With gzip
mysqldump -u user -p dbname | gzip > backup.sql.gz

# Specific tables
mysqldump -u user -p dbname table1 table2 > tables.sql

# Restore
mysql -u user -p dbname < backup.sql
```

### MongoDB

```bash
# Export database
mongodump --uri="mongodb://user:pass@host/dbname" --out=backup/

# Export specific collection
mongodump --uri="mongodb://host/dbname" --collection=users --out=backup/

# Import database
mongorestore --uri="mongodb://host/dbname" backup/dbname/

# Export to JSON
mongoexport --uri="mongodb://host/dbname" --collection=users --out=users.json
```

## Automated Backups

### Cron Job

```bash
# Edit crontab
crontab -e

# Daily backup at 2 AM
0 2 * * * rsync -av /source/ /backup/ >> /var/log/backup.log 2>&1

# Weekly full backup
0 3 * * 0 tar -czvf /backup/weekly-$(date +\%Y\%m\%d).tar.gz /source/

# Monthly database backup
0 4 1 * * pg_dump -U user dbname | gzip > /backup/db-$(date +\%Y\%m).sql.gz
```

### Backup Script

```bash
#!/bin/bash
# backup.sh

DATE=$(date +%Y%m%d_%H%M%S)
BACKUP_DIR="/backup/$DATE"
SOURCE="/data"

mkdir -p "$BACKUP_DIR"

# Files
rsync -av "$SOURCE/" "$BACKUP_DIR/files/"

# Database
pg_dump -U user mydb | gzip > "$BACKUP_DIR/database.sql.gz"

# Cleanup old backups (keep last 7)
find /backup -type d -mtime +7 -exec rm -rf {} \;

echo "Backup completed: $BACKUP_DIR"
```

## Cloud Backup

### AWS S3

```bash
# Sync to S3
aws s3 sync /local/ s3://bucket/backup/

# With glacier storage
aws s3 sync /local/ s3://bucket/backup/ --storage-class GLACIER

# Restore from S3
aws s3 sync s3://bucket/backup/ /local/
```

### rclone (Multi-cloud)

```bash
# Setup
rclone config

# Sync to Google Drive
rclone sync /local/ gdrive:backup/

# Sync to Dropbox
rclone sync /local/ dropbox:backup/

# Encrypted backup
rclone sync /local/ crypt:backup/
```

## Tips

- Use `--dry-run` before running backups
- Test restore procedures regularly
- Keep multiple backup generations
- Encrypt sensitive backups
- Store offsite or in cloud

## Triggers

backup, restore, rsync, database backup, file backup,
pg_dump, mysqldump, archive
