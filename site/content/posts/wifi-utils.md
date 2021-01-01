---
title: "Useful wifi commands"
date: 2020-01-01T23:34:43+01:00
draft: false
tags: ["linux", "wifi", "network"]
---

List used wifi frequencies
```bash
sudo iwlist wlp1s0 scan | grep Frequency | sort | uniq -c | sort -n
```

List available channels
```bash
iwlist channel
```

List wifi networks
```bash
nmcli d wifi
```
