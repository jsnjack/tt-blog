---
title: "Upgrade to fedora 40"
date: 2024-04-23T22:41:05+02:00
draft: false
tags: ["linux", "fedora", "fedora40"]
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
sudo dnf system-upgrade download --releasever=40
```

Reboot and start upgrade process
```bash
sudo dnf system-upgrade reboot
```
