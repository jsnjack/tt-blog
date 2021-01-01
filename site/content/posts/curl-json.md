---
title: "Prettify json output in curl"
date: 2020-01-01T23:14:49+01:00
draft: false
tags: ["linux", "curl"]
---

Install jq package:
```bash
sudo dnf install jq
```

Usage example:
```bash
curl -u key:x "https://api.com/users.json" | jq
```
