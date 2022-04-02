---
title: "No sound in Fedora after sleep"
date: 2022-03-28T13:51:12+02:00
draft: false
tags: ["linux", "fedora", "sound"]
---

If you have no sound after the system comes back from sleep, try to run this command:
```bash
sudo systemctl start systemd-suspend
```
