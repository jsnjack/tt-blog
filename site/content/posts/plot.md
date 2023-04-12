---
title: "Plot real time ping data in terminal"
date: 2023-04-12T11:28:22+02:00
draft: false
tags: ["linux", "plot", "debug"]
---

Install `ttyplot`:
```bash
grm install tenox7/ttyplot -n ttyplot
```

Ping plot:
```bash
ping 8.8.8.8 | sed -u 's/^.*time=//g; s/ ms//g' | ttyplot -t "ping to 8.8.8.8" -u ms
```

[More examples](https://github.com/tenox7/ttyplot)
