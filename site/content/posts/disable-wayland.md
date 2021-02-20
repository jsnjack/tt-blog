---
title: "Disable Wayland and switch to X11"
date: 2021-02-20T21:53:35+01:00
draft: false
tags: ["linux", "fedora", "x11"]
---

By default fedora uses Wayland. Make sure the following section exists in `/etc/gdm/custom.conf`:
```
[daemon]
# Uncomment the line below to force the login screen to use Xorg
WaylandEnable=false
DefaultSession=gnome-xorg.desktop
```
