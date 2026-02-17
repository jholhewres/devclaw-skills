---
name: redis-cli
description: "Interagir com Redis: keys, get, set, info e monitor"
metadata: {"openclaw":{"always":false,"emoji":"🔴"}}
---
# Redis CLI

Interface com Redis para cache, filas e dados em memória.

## Conexão

```bash
# Local
redis-cli

# Remoto
redis-cli -h <host> -p <port>
redis-cli -h <host> -p <port> -a <password>

# Com URL
redis-cli -u redis://:<password>@<host>:<port>/<db>

# Ping
redis-cli ping
```

## Keys

```bash
# Listar keys (cuidado em produção!)
redis-cli KEYS "prefix:*"        # pattern match
redis-cli SCAN 0 MATCH "user:*" COUNT 100   # seguro para produção

# Info de uma key
redis-cli TYPE <key>
redis-cli TTL <key>              # tempo restante (-1 = sem expiração)
redis-cli OBJECT ENCODING <key>

# Deletar
redis-cli DEL <key>
redis-cli DEL key1 key2 key3

# Expiração
redis-cli EXPIRE <key> <seconds>
redis-cli PERSIST <key>          # remover expiração
```

## Strings

```bash
redis-cli SET <key> <value>
redis-cli SET <key> <value> EX 3600   # expira em 1h
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
redis-cli LRANGE <key> 0 -1     # todos elementos
redis-cli LLEN <key>
```

## Sets

```bash
redis-cli SADD <key> <member>
redis-cli SMEMBERS <key>
redis-cli SCARD <key>            # count
redis-cli SISMEMBER <key> <member>
```

## Info e Monitor

```bash
# Info geral
redis-cli INFO
redis-cli INFO memory
redis-cli INFO stats
redis-cli INFO keyspace

# Clientes conectados
redis-cli CLIENT LIST

# Monitor (ver comandos em tempo real - usar com cuidado)
redis-cli MONITOR

# Tamanho do banco
redis-cli DBSIZE

# Flush (CUIDADO!)
redis-cli FLUSHDB       # banco atual
redis-cli FLUSHALL      # todos bancos
```

## Tips

- Use `SCAN` em vez de `KEYS` em produção (não bloqueia)
- Use `--pipe` para import em massa
- Use `--csv` para output em CSV
- Use `--bigkeys` para encontrar keys grandes
- Use `--latency` para medir latência
