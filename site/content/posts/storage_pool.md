---
title: "lxd: Manually remove lvm storage pool"
date: 2022-04-22T18:01:38+02:00
draft: false
tags: ["linux", "lxd"]
---

```bash
lvremove /dev/default/LXDThinPool
```

To list all lvm pools:
```bash
lvdisplay
```
