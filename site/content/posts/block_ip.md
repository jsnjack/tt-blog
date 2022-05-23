---
title: "Block IP address in firewalld"
date: 2022-05-23T11:47:04+02:00
draft: false
tags: ["linux", "firewalld"]
---

```bash
sudo firewall-cmd --add-rich-rule="rule family='ipv4' source address='89.20.160.77' reject" --timeout=1h
```

Valid values for timeout  - numbers followed by `s`, `m` or `h`

[Documentation](https://firewalld.org/documentation/man-pages/firewall-cmd.html)
