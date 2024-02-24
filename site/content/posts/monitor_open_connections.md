---
title: "Monitor open connections"
date: 2024-02-24T19:23:49+01:00
draft: false
tags: ["linux", "debug"]
---

This command will plot open connections:
```
{ while true; do ss -ant | grep ESTAB | wc -l; sleep 0.1; done } | ttyplot
```
