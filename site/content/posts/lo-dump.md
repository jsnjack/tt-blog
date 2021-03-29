---
title: "Record all traffic on specific port"
date: 2021-02-15T14:51:35+01:00
draft: false
tags: ["linux", "debug", "wireshark", "tcpdump"]
---

Create TCP dump of all traffic on localhost on port 8888
```bash
sudo tcpdump -i lo "port 8888" -w dump
```
