---
title: "Automatically start application during the boot "
date: 2015-01-01T23:02:26+01:00
draft: false
tags: ["linux", "gnome", "fedora"]
---
To start an application automatically in fedora, you need to create a *.desktop file in the directory ~/.config/autostart, for example:
```bash
gedit ~/.config/autostart/flux.desktop
```

The content of the file can be something like:
```c
[Desktop Entry]
Type=Application
Name=flux
Comment=xflux
Exec=/home/jsn/app/xflux -l 52.3837151 -g 4.8806328
Terminal=false
```
