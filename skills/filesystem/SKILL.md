---
name: filesystem
version: 0.1.0
author: goclaw
description: "File system operations — advanced file manipulation, search, and management"
category: system
tags: [filesystem, files, search, management, system]
requires:
  bins: [find, xargs, rsync]
---
# Filesystem

Advanced file system operations and management.

## Search Files

```bash
# Find by name
find . -name "*.txt"

# Find by extension
find . -type f -name "*.py"

# Case insensitive
find . -iname "*.TXT"

# Find directories
find . -type d -name "node_modules"

# Find by size
find . -size +100M           # Larger than 100MB
find . -size -1k             # Smaller than 1KB

# Find by time
find . -mtime -7             # Modified in last 7 days
find . -mtime +30            # Not modified in 30+ days
find . -atime -1             # Accessed in last 24h
find . -newer file.txt       # Newer than file.txt

# Find and execute
find . -name "*.log" -delete
find . -name "*.tmp" -exec rm {} \;
find . -name "*.py" -exec wc -l {} +

# Find with permissions
find . -perm 755
find . -executable -type f
```

## Search Content (ripgrep/grep)

```bash
# Search in files
grep -r "pattern" .

# With ripgrep (faster)
rg "pattern" .

# Search specific files
rg "pattern" -g "*.py"

# Show context
rg "pattern" -C 3           # 3 lines before and after
rg "pattern" -A 2           # 2 lines after
rg "pattern" -B 2           # 2 lines before

# Count matches
rg -c "pattern" .

# Files with matches only
rg -l "pattern" .

# Case insensitive
rg -i "pattern" .
```

## Copy & Sync

```bash
# Sync directories with rsync
rsync -av source/ destination/

# Sync with progress
rsync -av --progress source/ destination/

# Sync with exclude
rsync -av --exclude 'node_modules' --exclude '.git' source/ destination/

# Sync to remote
rsync -avz -e ssh source/ user@host:/path/

# Delete files not in source
rsync -av --delete source/ destination/

# Dry run
rsync -av --dry-run source/ destination/
```

## Disk Usage

```bash
# Directory sizes
du -sh *

# Sort by size
du -sh * | sort -h

# Find largest files
du -ah /path | sort -rh | head -20

# Disk free space
df -h

# Inode usage
df -i
```

## File Operations

```bash
# Create directory tree
mkdir -p path/to/deep/directory

# Copy with progress
cp -r source/ destination/ & progress -p $!

# Move with overwrite prompt
mv -i source dest

# Remove empty directories
find . -type d -empty -delete

# Remove files older than X days
find . -name "*.log" -mtime +30 -delete

# Batch rename
rename 's/.txt/.md/' *.txt

# Create tarball
tar -czvf archive.tar.gz directory/

# Extract tarball
tar -xzvf archive.tar.gz
```

## File Permissions

```bash
# Change permissions
chmod 755 script.sh
chmod -R 644 *.txt

# Change ownership
chown user:group file.txt
chown -R user:group directory/

# Make executable
chmod +x script.sh

# Find world-writable files
find . -perm -0002 -type f

# Find files without owner
find . -nouser -o -nogroup
```

## Symbolic Links

```bash
# Create symlink
ln -s /path/to/target linkname

# Hard link
ln target linkname

# Remove symlink
rm linkname

# Find broken symlinks
find . -type l -! -exec test -e {} \; -print

# View link target
readlink -f linkname
```

## Watch for Changes

```bash
# Watch directory changes
inotifywait -m -r /path/to/watch

# Watch with incron (like cron for filesystem)
incrontab -e
# /path/to/watch IN_CLOSE_WRITE /script.sh $@/$#
```

## Tips

- Use `find ... -exec ... +` for better performance
- Use `rsync -n` (dry-run) before real sync
- Use `shred` for secure file deletion
- Use `lsof` to find open files
- Use `fuser` to find processes using files

## Triggers

filesystem, find files, search files, file operations, rsync,
disk usage, file management, file search
