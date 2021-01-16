---
title: "Monitor thread count"
date: 2021-01-15T15:01:54+01:00
draft: false
tags: ["linux", "debug", "monitor"]
---

```bash
watch -n 0.5 "ps -eLf | grep traffic_server | wc -l"
```

