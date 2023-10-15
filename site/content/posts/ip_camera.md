---
title: "How to use Android phone camera on desktop computer via WiFi with IP Webcam"
date: 2023-10-15T15:46:23+02:00
draft: false
tags: ["linux", "fedora", "camera"]
---

1. Install [IP Camera](https://play.google.com/store/apps/details?id=com.pas.webcam). The application is free, with minimal adds. Compared to DroidCam, it allows you to stream FullHD video
2. Install dependencies:
```bash
sudo dnf install ffmpeg v4l2loopback
```
3. On mobile phone, start the application and select "Start server" from the menu. Select also "Fade" option to set minimal brightness of the mobile phone screen.
4. On desktop, run the following command:
```bash
ffmpeg -i http://192.168.2.26:8080/videofeed -vf format=yuv420p -f v4l2 /dev/video0
```
