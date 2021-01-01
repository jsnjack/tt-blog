---
title: "How to format USB stick as FAT32"
date: 2016-01-01T23:08:52+01:00
draft: false
tags: ["linux", "usb"]
---

First step is to find the name of the USB stick in the system:
```bash
sudo fdisk -l
```

In my case the name of the device is /dev/sdc If the USB stick is mounted, unmount it with command:
```bash
sudo umount /dev/sdc
```

Finally, run the command to format USB stick:
```bash
mkdosfs -F 32 -I /dev/sdc
```
