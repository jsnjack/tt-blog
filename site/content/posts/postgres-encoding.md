---
title: "How to fix ERROR: new encoding (UTF8) is incompatible"
date: 2016-01-01T23:11:08+01:00
draft: false
tags: ["linux", "postgresql", "encoding"]
---

Sometimes I'm unable to create a database in UTF8 encoding because the template database is in the ASCII encoding. To fix it we need to recreate template database in UTF8. Start psql console:
```bash
psql -U postgres
```

Run the following commands:
```bash
UPDATE pg_database SET datistemplate = FALSE WHERE datname = 'template1';

DROP DATABASE template1;

CREATE DATABASE template1 WITH TEMPLATE = template0 ENCODING = 'UNICODE';

UPDATE pg_database SET datistemplate = TRUE WHERE datname = 'template1';

\c template1

VACUUM FREEZE;
```
