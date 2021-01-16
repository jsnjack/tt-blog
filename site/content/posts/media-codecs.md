---
title: "Install multimedia codecs in fedora"
date: 2021-01-16T21:27:31+01:00
draft: false
tags: ["linux", "fedora", "fedora33", "youtube", "codecs", "cpu", "firefox"]
---
If you are experiencing high CPU usage when watching videos on youtube, verify that the following packages are installed (requires RPMFusion repository):
```bash
# RPMFuison repositories
sudo dnf install https://mirrors.rpmfusion.org/free/fedora/rpmfusion-free-release-$(rpm -E %fedora).noarch.rpm https://mirrors.rpmfusion.org/nonfree/fedora/rpmfusion-nonfree-release-$(rpm -E %fedora).noarch.rpm

# Codecs
sudo dnf install gstreamer1-plugins-{bad-\*,good-\*,base} gstreamer1-plugin-openh264 gstreamer1-libav --exclude=gstreamer1-plugins-bad-free-devel
sudo dnf install lame\* --exclude=lame-devel
sudo dnf group upgrade --with-optional Multimedia
```
CPU usage for i3-8100T is down from 80% to 50% with this [video sample](https://www.youtube.com/watch?v=A_hF37w6Uao&t=293s)
