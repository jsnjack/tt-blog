---
title: "How to reload the connection with NetworkManager"
date: 2021-12-30T13:14:17+01:00
draft: false
tags: ["linux", "network", "nmcli"]
---

Reload interface configuration:
```bash
sudo nmcli connection reload
```

This command will ONLY make NetworkManager aware of new configuration changes. To apply the new changes, run the following command:
```bash
sudo nmcli dev reapply <interface>
```
