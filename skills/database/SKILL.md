---
name: database
version: 0.1.0
author: devclaw
description: "Database operations — PostgreSQL, MySQL, MongoDB, Redis queries"
category: data
tags: [database, postgres, mysql, mongodb, redis, sql]
requires:
  bins: [psql, mysql, mongosh, redis-cli]
---
# Database

Interact with various databases using CLI tools.

## PostgreSQL

```bash
# Connect to database
psql -h localhost -U user -d dbname

# Connection string
psql "postgresql://user:pass@host:5432/dbname"

# Execute query
psql -h localhost -U user -d dbname -c "SELECT * FROM users LIMIT 10;"

# Execute from file
psql -f queries.sql

# List databases
psql -l

# List tables
psql -d dbname -c "\dt"

# Describe table
psql -d dbname -c "\d tablename"

# Export to CSV
psql -d dbname -c "COPY (SELECT * FROM users) TO STDOUT WITH CSV HEADER" > users.csv

# Import CSV
psql -d dbname -c "\COPY users FROM 'users.csv' WITH CSV HEADER"
```

## MySQL

```bash
# Connect
mysql -h localhost -u user -p dbname

# Execute query
mysql -h localhost -u user -p dbname -e "SELECT * FROM users LIMIT 10;"

# Execute from file
mysql -h localhost -u user -p dbname < queries.sql

# Export database
mysqldump -u user -p dbname > backup.sql

# Import database
mysql -u user -p dbname < backup.sql

# Export single table
mysqldump -u user -p dbname tablename > table.sql
```

## MongoDB

```bash
# Connect
mongosh "mongodb://user:pass@host:27017/dbname"

# Or with connection string
mongosh "mongodb+srv://cluster.example.com/dbname" --username user

# Execute query (eval)
mongosh dbname --eval 'db.users.find().limit(10)'

# Export collection
mongoexport --uri="mongodb://host/dbname" --collection=users --out=users.json

# Import collection
mongoimport --uri="mongodb://host/dbname" --collection=users --file=users.json

# Dump database
mongodump --uri="mongodb://host/dbname" --out=backup/

# Restore database
mongorestore --uri="mongodb://host/dbname" backup/dbname/
```

## Redis

```bash
# Connect
redis-cli -h localhost -p 6379

# With password
redis-cli -h localhost -p 6379 -a password

# Execute command
redis-cli -h localhost GET mykey

# Set value
redis-cli -h localhost SET mykey "value"

# List all keys
redis-cli -h localhost KEYS '*'

# Get database info
redis-cli -h localhost INFO

# Monitor commands
redis-cli -h localhost MONITOR

# Export database
redis-cli -h localhost BGSAVE

# Get all key-value pairs
redis-cli -h localhost --scan | xargs -L 1 redis-cli -h localhost GET
```

## SQLite

```bash
# Connect
sqlite3 database.db

# Execute query
sqlite3 database.db "SELECT * FROM users LIMIT 10;"

# Import CSV
sqlite3 database.db -cmd ".mode csv" -cmd ".import users.csv users"

# Export to CSV
sqlite3 database.db -cmd ".mode csv" -cmd ".output users.csv" "SELECT * FROM users;"

# Dump database
sqlite3 database.db .dump > backup.sql
```

## Common Queries

```bash
# PostgreSQL - Create table
psql -d dbname -c "
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  email VARCHAR(255) UNIQUE,
  created_at TIMESTAMP DEFAULT NOW()
);"

# MySQL - Create table
mysql -u user -p dbname -e "
CREATE TABLE users (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  email VARCHAR(255) UNIQUE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);"

# MongoDB - Create index
mongosh dbname --eval 'db.users.createIndex({email: 1})'
```

## Tips

- Use connection strings for complex auth
- Use `-c` or `-e` for single queries
- Use `\d` in psql to describe tables
- Use `.help` in SQLite for commands
- Always backup before destructive operations

## Triggers

database, postgres, mysql, mongodb, redis, sql, database query,
psql, database connect
