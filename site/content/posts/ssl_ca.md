---
title: "Establish SSL connection with openssl using custom CA bundle"
date: 2022-04-21T11:23:07+02:00
draft: false
tags: ["linux", "openssl", "ssl"]
---

```bash
openssl s_client -verify_return_error -CAfile ats_certs.pem -showcerts -connect www.zilverenkruis.nl:443 -servername www.zilverenkruis.nl
```
