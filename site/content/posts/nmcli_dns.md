---
title: "Change DNS nameservers using NetworkManager"
date: 2021-12-30T13:18:17+01:00
draft: false
tags: ["linux", "network", "nmcli"]
---

```bash
sudo nmcli connection modify <connection name> ipv4.dns "1.1.1.1 8.8.8.8"
```

Apply changes:
```bash
sudo nmcli dev reapply <interface>
```
