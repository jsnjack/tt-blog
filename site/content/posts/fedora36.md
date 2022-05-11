---
title: "Upgrade to fedora 36"
date: 2022-05-11T13:41:05+02:00
draft: false
tags: ["linux", "fedora", "fedora36"]
---

Make sure your current system is up to date:
```bash
sudo dnf upgrade --refresh
```

Reboot:
```bash
reboot
```

Install dnf plugin which updates OS:
```bash
sudo dnf install dnf-plugin-system-upgrade
```

Download and prepare packages for upgrade:
```bash
sudo dnf system-upgrade download --releasever=36
```

Reboot and start upgrade process
```bash
sudo dnf system-upgrade reboot
```

### Fix known issues
#### Missing icons in openweather extension
```bash
sudo dnf install gnome-icon-theme
```

Restart gnome-shell to apply changes: `Alt + F2` and type `r`

#### Blurry font in GTK4 applications
Create `~/.config/gtk-4.0/settings.ini` file with the following content:
```
[Settings]
gtk-hint-font-metrics=1
```

#### Replace gedit with new text editor
```bash
sudo dnf remove gedit
sudo dnf install gnome-text-editor
```
