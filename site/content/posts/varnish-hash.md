---
title: "Log Varnish hash data"
date: 2021-12-17T13:49:41+01:00
draft: false
tags: ["linux", "varnish", "debug"]
---

When starting Varnish, add `-p vsl_mask=+Hash` argument to the command. Print
varnish logs with this command:
```bash
sudo varnishlog -n /opt/varnish -q 'ReqURL ~ "my_request"'
```
