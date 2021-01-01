---
title: "Show raw SQL queries in PostgreSQL"
date: 2017-01-01T23:13:27+01:00
draft: false
tags: ["linux", "postgresql", "debug"]
---

If you want to inspect what queries exactly Django ORM sends to the PostgreSQL, you can do it with postgres logging. The first step is to enable logging. Add
```
log_statement = 'all' 
```
to the file:
```
sudo vim /etc/postgresql/9.4/main/postgresql.conf
```
Then you are able to see raw SQL queries with command:
```
sudo tail -f /var/log/postgresql/postgresql-9.4-main.log 
```
