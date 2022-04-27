---
title: "lxd: Increase disk size of an existing container"
date: 2019-01-01T23:31:02+01:00
draft: false
tags: ["linux", "lxd"]
---

```bash
lxc config device override mycontainer root size=20GB
```
Restart the container to apply changes.
