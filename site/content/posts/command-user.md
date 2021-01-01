---
title: "Run command as a different user"
date: 2017-10-01T23:23:47+01:00
draft: false
tags: ["linux", "bash"]
---

Command runuser allows to run a command as a different user. You must be root to be able to run that command:
```bash
sudo runuser -l vagrant -c "GH=1 env"
```
