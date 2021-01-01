---
title: "Monitor process with top"
date: 2020-01-01T23:25:56+01:00
draft: false
tags: ["linux", "debug", "monitor"]
---

One liner to monitor a process with the name /wshub:
```bash
top -p "$(pgrep -f /wshub)"
```
