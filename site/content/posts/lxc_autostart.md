---
title: "Do not start LXD containers on boot"
date: 2022-02-28T16:59:24+01:00
draft: false
tags: ["linux", "lxd"]
---

For all newly created containers:
```bash
lxc profile set default boot.autostart=false
```

Or for existing container:
```bash
lxc config set <container_name> boot.autostart false
```
