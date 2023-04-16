---
title: "Monitor number of open sockets by a process"
date: 2023-04-16T14:22:22+02:00
draft: false
tags: ["linux", "debug"]
---

Example with `watch` command:
```bash
watch -n 1 'ps aux | kazy -i cmproxy -e kazy | kazy -r "[\d]+" -x | xargs -I % ls -l /proc/%/fd/ | kazy -i socket | wc -l'
```

Example with `ttyplot` command:
```bash
{ while true; do ps aux | kazy -i cmproxy -e kazy | kazy -r "[\d]+" -x | xargs -I % ls -l /proc/%/fd/ | kazy -i socket | wc -l; sleep 1; done } | ttyplot
```
