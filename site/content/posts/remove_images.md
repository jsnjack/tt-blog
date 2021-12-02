---
title: "lxd: delete all images"
date: 2021-12-02T13:09:46+01:00
draft: false
tags: ["linux", "lxd"]
---

Delete all images which contain `build-cache` in their name:
```bash
lxc image list -c l --format csv | grep build-cache | xargs -I % lxc image delete %
```

