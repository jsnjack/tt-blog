---
title: "Setup lxd server as remote server"
date: 2021-09-03T13:34:50+02:00
draft: false
tags: ["linux", "lxd"]
---

```bash
lxc config set core.https_address "[::]"
lxc config set core.trust_password new_password
```

After all clients are connected to the remote it is recommended to unset password:
```bash
lxc config unset core.trust_password
```
