---
title: "How to use Android phone camera on desktop computer via WiFi with IP Webcam"
date: 2023-10-15T15:46:23+02:00
draft: false
tags: ["linux", "fedora", "camera"]
---

1. Install [IP Camera](https://play.google.com/store/apps/details?id=com.pas.webcam). The application is free, with minimal adds. Compared to DroidCam, it allows you to stream FullHD video
2. Install dependencies:
```bash
sudo dnf install ffmpeg v4l2loopback akmod-v4l2loopback
```
3. Verify that v4l2loopback is loaded:
```bash
lsmod | grep v4l2loopback
```
4. If it is not loaded, run this command:
```bash
akmods --kernels $(uname-r) --rebuild
# Confirm that it was loaded
```
4. On mobile phone, start the application and select "Start server" from the menu. Select also "Fade" option to set minimal brightness of the mobile phone screen.
5. On desktop, run the following command:
```bash
ffmpeg -i http://192.168.2.26:8080/videofeed -vf format=yuv420p -f v4l2 /dev/video0
```

> If for some reason `v4l2loopback` module is not compiled correctly, you can try
> to compile it manually by following [this instruction](https://github.com/seii/fedora-green-screen/blob/master/README.md#installing-v4l2loopback)
