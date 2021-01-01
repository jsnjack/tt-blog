---
title: "Use mitmproxy with GNOME Web (Epiphany)"
date: 2017-07-01T23:21:34+01:00
draft: false
tags: ["linux", "mitmproxy", "epiphany"]
---

To use mitmproxy with Epiphany you need to install CA certificate from mitmproxy. To do that, start mitmproxy and navigate to mitm.it. Download *.pem certificate (available under Other link). Install mitmproxy certificate with command:
```bash
sudo trust anchor mitmproxy-ca-cert.pem 
```
With this command you can also install any custom certificate to use in Epiphany. 
