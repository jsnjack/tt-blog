---
title: "Increase watches limit"
date: 2016-01-01T23:10:06+01:00
draft: false
tags: ["linux", "system"]
---

Open configuration file:
```bash
sudo vim /etc/sysctl.d/90-override.conf
```

Add the following line to the end of the file to increase watches limit:
```bash
fs.inotify.max_user_watches = 524288
```
