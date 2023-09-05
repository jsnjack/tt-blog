---
title: "sumologic: filter fields"
date: 2023-09-05T22:21:27+02:00
draft: false
tags: ["sumologic"]
---

```
(sourceCategory=prod/* ss-haproxy)
| parse "{|||}" as host,forwarded, referer, ua
| where forwarded matches ","
```
