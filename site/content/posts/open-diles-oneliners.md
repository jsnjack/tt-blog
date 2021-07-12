---
title: "One liners to check open files limit"
date: 2021-07-12T19:12:53+02:00
draft: false
tags: ["linux", "system", "systemd", "nofile", "ulimit", "debug"]
---

Check open files limit for the process named redis:
```bash
cat /proc/$(ps aux | grep redis | head -n 1 | kazy -x -r "\d+")/limits | kazy -i "open files" -i Limit
```

Check current number of open files for the process named redis:
```bash
sudo ls -l /proc/$(ps aux | grep redis | head -n 1 | kazy -x -r "\d+")/fd | wc -l
```
