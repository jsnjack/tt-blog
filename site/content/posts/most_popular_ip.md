---
title: "Extract and count IP addresses from logs"
date: 2021-03-29T23:25:40+02:00
draft: false
tags: ["linux", "journalctl", "debug", "network"]
---

```bash
journalctl -n 1000 | grep -E -o "([0-9]{1,3}[\.]){3}[0-9]{1,3}" | sort | uniq -c
```
