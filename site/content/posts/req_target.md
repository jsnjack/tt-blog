---
title: "curl: Request URL via proxy"
date: 2022-04-21T11:14:16+02:00
draft: false
tags: ["linux", "curl", "proxy"]
---

It is possible to request a URL in curl via proxy server using `-x` flag:
```bash
curl -x http://127.0.0.1:8080 https://google.com
```

This will send CONNECT request to your proxy server 127.0.0.1:8080. CONNECT request
will ask proxy server to tunnel TCP connection to the destination. This means that, for example,
proxy server won't verify SSL certificate.

To force curl to send regular GET request, use the following command:
```bash
curl --request-target https://google.com -x 127.0.0.1:8080 127.0.0.1:8080
```
