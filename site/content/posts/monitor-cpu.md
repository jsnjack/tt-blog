---
title: "Monitor cpu and memory usage of the process"
date: 2020-01-01T23:24:41+01:00
draft: false
tags: ["linux", "debug", "monitor", "cpu"]
---

Print every second cpu and memory usage of the process:
```bash
watch -n 1 "ps -eo pid,pcpu,pmem,args | kazy -i xnotitle -e kazy"
```
