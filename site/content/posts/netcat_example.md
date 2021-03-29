---
title: "netcat example"
date: 2021-03-29T23:28:02+02:00
draft: false
tags: ["linux", "network", "debug"]
---

Listen on specified port:
```bash
nc -l -v -k 8888
```

Connect via netcat:
```bash
nc -v google.com 80
```
