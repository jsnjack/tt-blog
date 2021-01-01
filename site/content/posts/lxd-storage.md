---
title: "Increase storage of a lxd container"
date: 2019-01-01T23:31:02+01:00
draft: false
tags: ["linux", "lxd"]
---

By default lxd stopped to assign root (disk) device to a container. To modify storage volume of a container add a device like this:
```bash
lxc config device add surflyc root disk pool=default path=/ size=15GB
```
Restart the container to apply changes. 
