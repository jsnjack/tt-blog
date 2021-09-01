---
title: "Show top 10 process with open files"
date: 2021-07-13T13:15:00+02:00
draft: false
tags: ["linux", "debug", "nofile"]
---

```bash
lsof | awk '{print $1}' | sort | uniq -c | sort -r -n | head
```
