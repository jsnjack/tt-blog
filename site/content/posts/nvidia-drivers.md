---
title: "Install nvidia drivers"
date: 2021-02-20T22:09:42+01:00
draft: false
tags: ["linux", "fedora", "fresh system", "gpu"]
---

Verify your GPU:
```bash
/sbin/lspci | grep -e VGA
```

Make sure you:
 - disabled Wayland
 - disabled Secure Boot

Make sure you system is up to date (reboot if something was installed):
```bash
sudo dnf update -y
```

Install nvidia driver:
```bash
sudo dnf install akmod-nvidia
sudo dnf install xorg-x11-drv-nvidia-cuda
```

Wait until kernel module is compiled.

Enable video acceleration and some extras:
```bash
sudo dnf install vdpauinfo libva-vdpau-driver libva-utils xorg-x11-drv-nvidia-cuda-libs xorg-x11-drv-nvidia-cuda
```

Reboot system.

Verify that `nouveau` driver is disabled:
```
lsmod |grep nouveau
```
