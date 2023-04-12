---
title: "Monitor DNS query time"
date: 2023-04-12T11:50:00+02:00
draft: false
tags: ["linux", "dns", "plot", "debug"]
---

```bash
{ while true; do dig google.com | kazy -i "Query time" | kazy -r -x "[\d.]*"; sleep 1; done } | ttyplot -t "dig to google.com"
```

