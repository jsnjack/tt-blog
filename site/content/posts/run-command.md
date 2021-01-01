---
title: "Run command on suspend "
date: 2020-01-01T23:45:36+01:00
draft: false
tags: ["linux", "systemd"]
---

Create a file located in `/usr/lib/systemd/system-sleep` directory:
```bash
#!/bin/sh
if [ "${1}" == "pre" ]; then
  nordvpn d
elif [ "${1}" == "post" ]; then
  echo "Hello"
fi
```
