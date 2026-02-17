---
name: redis-cli
description: "Redis CLI for cache, queues, and in-memory data"
metadata: {"openclaw":{"always":false,"emoji":"🔴"}}
---
# Redis CLI

Interface for Redis cache, queues, and in-memory data.

## Setup

```bash
# Check if installed
command -v redis-cli

# Install — macOS
brew install redis

# Install — Ubuntu/Debian
sudo apt install redis-tools
```

## Connection

```bash
# Local
redis-cli

# Remote
redis-cli -h <host> -p <port>
redis-cli -h <host> -p <port> -a <password>

# With URL
redis-cli -u redis://:<password>@<host>:<port>/<db>

# Ping
redis-cli ping
```

## Keys

```bash
# List keys (WARNING: be careful in production!)
redis-cli KEYS "prefix:*"        # pattern match
redis-cli SCAN 0 MATCH "user:*" COUNT 100   # safe for production

# Info for a key
redis-cli TYPE <key>
redis-cli TTL <key>              # remaining time (-1 = no expiration)
redis-cli OBJECT ENCODING <key>

# Delete
redis-cli DEL <key>
redis-cli DEL key1 key2 key3

# Expiration
redis-cli EXPIRE <key> <seconds>
redis-cli PERSIST <key>          # remove expiration
```

## Strings

```bash
redis-cli SET <key> <value>
redis-cli SET <key> <value> EX 3600   # expires in 1h
redis-cli GET <key>
redis-cli MGET key1 key2 key3
redis-cli INCR <key>
redis-cli INCRBY <key> 10
```

## Hashes

```bash
redis-cli HSET <key> <field> <value>
redis-cli HGET <key> <field>
redis-cli HGETALL <key>
redis-cli HDEL <key> <field>
```

## Lists

```bash
redis-cli LPUSH <key> <value>
redis-cli RPUSH <key> <value>
redis-cli LRANGE <key> 0 -1     # all elements
redis-cli LLEN <key>
```

## Sets

```bash
redis-cli SADD <key> <member>
redis-cli SMEMBERS <key>
redis-cli SCARD <key>            # count
redis-cli SISMEMBER <key> <member>
```

## Info and Monitor

```bash
# General info
redis-cli INFO
redis-cli INFO memory
redis-cli INFO stats
redis-cli INFO keyspace

# Connected clients
redis-cli CLIENT LIST

# Monitor (see commands in real time — use with caution)
redis-cli MONITOR

# Database size
redis-cli DBSIZE

# Flush (WARNING!)
redis-cli FLUSHDB       # current database
redis-cli FLUSHALL      # all databases
```

## Tips

- Use `SCAN` instead of `KEYS` in production (non-blocking)
- Use `--pipe` for bulk import
- Use `--csv` for CSV output
- Use `--bigkeys` to find large keys
- Use `--latency` to measure latency
