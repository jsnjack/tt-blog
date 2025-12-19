---
title: "Install all video codecs for AMD GPU"
date: 2025-12-19T22:17:11+01:00
draft: false
tags: ["linux", "amd", "fedora", "fresh system"]
---

Install all good (the ones that work) video codecs for AMD GPU:
```bash
# Enable RPM Fusion repositories
sudo dnf install https://mirrors.rpmfusion.org/free/fedora/rpmfusion-free-release-$(rpm -E %fedora).noarch.rpm https://mirrors.rpmfusion.org/nonfree/fedora/rpmfusion-nonfree-release-$(rpm -E %fedora).noarch.rpm

# Install hardware acceleration libraries for AMD
sudo dnf swap mesa-va-drivers mesa-va-drivers-freeworld
sudo dnf swap mesa-vdpau-drivers mesa-vdpau-drivers-freeworld

# Switch to full ffmpeg package
sudo dnf swap ffmpeg-free ffmpeg --allowerasing

# Install additional codecs
sudo dnf group update multimedia --setopt="install_weak_deps=False" --exclude=PackageKit-gstreamer-plugin
sudo dnf group update sound-and-video
sudo dnf install gstreamer1-plugin-openh264 mozilla-openh264
sudo dnf config-manager setopt fedora-cisco-openh264.enabled=1
```
