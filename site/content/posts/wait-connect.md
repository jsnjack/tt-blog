---
title: "Wait until able to connect"
date: 2020-01-01T23:43:45+01:00
draft: false
tags: ["linux", "network", "bash"]
---

```bash
while ! nc -w5 -z 10.0.0.100 22; do echo "hello"; done
```
