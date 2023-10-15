---
title: "How to use Android phone camera on desktop computer via WiFi with DroidCam"
date: 2023-08-02T09:39:57+02:00
draft: false
tags: ["linux", "fedora", "camera"]
---

1. Install [DroidCam](https://play.google.com/store/apps/details?id=com.dev47apps.droidcam&hl=en&gl=US) on your phone. There are free and paid versions. Paid version enables HD quality
2. Install DroidCam in your desktop:
```bash
sudo dnf copr enable meeuw/droidcam
sudo dnf install droidcam
```
3. Launch both mobile and desktop applications. Connect desktop client to your phone's camera (IP address and port are displayed on your phone). Desktop client will connect to your camera via WiFi and create a virtual camera device.

If droidcam doesn't start with the following error:
```bash
$ droidcam
Fatal: droidcam video device reported pixel format 34524742 (BGR4), expected 32315559 (YU12/I420)
Try 'v4l2loopback-ctl set-caps "video/x-raw, format=I420, width=640, height=480" /dev/video<N>'
```

Run this command:
```bash
v4l2loopback-ctl set-caps /dev/video "YU12:640x480"
```
