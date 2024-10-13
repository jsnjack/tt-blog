---
title: "How to use your old Android phone as a webcam"
date: 2023-10-15T15:46:23+02:00
draft: false
tags: ["linux", "fedora", "camera"]
---

No matter how expensive your webcam is, it will never be as good as your old phone's camera.
In this guide, we will give a second life to your old phone by using it as a webcam.
All software used in this guide is free.

### Android phone
We will need to install the following software on your mobile phone:
1. [IP Camera](https://play.google.com/store/apps/details?id=com.pas.webcam) to stream
    video from your phone's camera to your desktop computer over WiFi.
2. (Optional) [droidVNC-NG VNC Server](https://play.google.com/store/apps/details/droidVNC_NG?id=net.christianbeier.droidvnc_ng&hl=en-US)
    to control your phone remotely from your desktop computer.

### Desktop computer
This guide is written for Fedora, but it should give you a general idea of how to
set it up on other distributions.

We will need to install `v4l2loopback` to create a virtual camera device and `ffmpeg`
to stream video from remote camera of the mobile phone to the virtual camera device.

`v4l2loopback` is a kernel module that allows you to create "virtual video devices".
It is available in rpmfusion repository.

1. Enable [rpmfusion repository](https://rpmfusion.org/Configuration):
```bash
sudo dnf install https://mirrors.rpmfusion.org/free/fedora/rpmfusion-free-release-$(rpm -E %fedora).noarch.rpm https://mirrors.rpmfusion.org/nonfree/fedora/rpmfusion-nonfree-release-$(rpm -E %fedora).noarch.rpm
```
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
akmods --kernels $(uname -r) --rebuild
# Confirm that it was loaded
lsmod | grep v4l2loopback
```

Add the following helper function to your `.bashrc` file to start streaming from your phone:
```bash
webcam () {
  ffmpeg -i http://192.168.2."$1":4747/videofeed -vf format=yuv420p -f v4l2 /dev/video0
}
```
Assuming that your phone's IP address is `192.168.2.38` and IP Camera is running on port `4747`,
you can start streaming by running `webcam 38` in your terminal. I keep my phone's IP assress dynamic as it
changes once in a while, if I don't use it for a long time.

### Hardware setup
I have an old Google Pixel 3a phone which is mounted on the top of my monitor. The
phone is attached to the mount with magnetic ring. At the time of writing, the setup
cost me around 25 EUR:
 - [Itramax magnetic phone holder](https://www.amazon.nl/dp/B0DBKT87GJ?ref=ppx_yo2ov_dt_b_fed_asin_title&th=1)
 - [Simarro magnetic rings](https://www.amazon.nl/dp/B0D6317RVG?ref=ppx_yo2ov_dt_b_fed_asin_title)

### How to use
1. On mobile phone, start `IP Camera` and select "Start server" from the menu.
In the bottom part of the screen, you will see the IP address and port number.
2. On desktop, run the following command, providing the IP address of your phone and the port number:
```bash
ffmpeg -i http://192.168.2.26:4747/videofeed -vf format=yuv420p -f v4l2 /dev/video0
## or use the helper function
webcam 26
```
> You can confgiure video resolution and other settings of your phone's camera by
> navigating to http://192.168.2.26:4747/ in your browser.

> If you install `droidVNC-NG VNC Server`, you can start `IP Camera` remotely from your desktop.
> This is useful if you want to start streaming from your phone without touching it. On the desktop computer,
> use `Connections` app (comes by default in GNOME) to connect to your phone.

### Troubleshooting
> If for some reason `v4l2loopback` module is not compiled correctly, you can
> compile it manually by following [this instruction](https://github.com/seii/fedora-green-screen/blob/master/README.md#installing-v4l2loopback).
