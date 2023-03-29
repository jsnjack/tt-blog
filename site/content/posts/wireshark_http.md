---
title: "Monitor all http traffic in wireshark on specific ports"
date: 2023-03-29T13:55:38+02:00
draft: false
tags: ["linux", "wireshark"]
---

Open wireshark GUI as root user:
```bash
sudo wireshark
```

Start capturing by clicking Capture button and add the following display filter:
```
http and (tcp.port == 8082 or tcp.port == 8000)
```
