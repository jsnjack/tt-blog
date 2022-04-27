---
title: "Do I need to restart after dnf update?"
date: 2022-04-22T13:27:44+02:00
draft: false
tags: ["linux", "dnf", "fedora"]
---

Reboot is needed if one of the following packages was updated:
```
kernel
glibc
linux-firmware
systemd
dbus
```

There is an application to easily check it:
```bash
sudo dnf install yum-utils
needs-restarting -r
```
