---
title: "How to find out which package provides a specific file "
date: 2015-01-01T22:55:10+01:00
draft: false
tags: ["linux", "dnf", "fedora"]
---

If you need to find out, for example, which package in fedora provides htpasswd command, you may do it with the command:
```bash
yum provides \*bin/htpasswd
```
