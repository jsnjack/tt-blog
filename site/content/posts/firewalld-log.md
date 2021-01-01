---
title: "Log denied connections in firewalld"
date: 2020-01-01T23:47:20+01:00
draft: false
tags: ["linux", "firewalld", "debug"]
---

Edit file `/etc/sysconfig/firewalld`:
```
FIREWALLD_ARGS=--debug=10
```
