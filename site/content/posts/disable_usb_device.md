---
title: "Disable USB device"
date: 2022-03-18T13:15:42+01:00
draft: false
tags: ["linux", "udev", "dell"]
---

To disable USB device (audio device from Dell WD-19TB dock station in the example) when connecting it to your laptop, find it's ids:
```bash
sudo lsusb -v
```

For example:
```
...
Bus 005 Device 007: ID 0bda:402e Realtek Semiconductor Corp. USB Audio
Device Descriptor:
  bLength                18
  bDescriptorType         1
  bcdUSB               2.00
  bDeviceClass            0 
  bDeviceSubClass         0 
  bDeviceProtocol         0 
  bMaxPacketSize0        64
  idVendor           0x0bda Realtek Semiconductor Corp.
  idProduct          0x402e 
  bcdDevice            0.01
  iManufacturer           3 Generic
  iProduct                1 USB Audio
  iSerial                 2 200901010001
...
```
Save `idVendor` and `idProduct` values.

Create new file `/lib/udev/rules.d/90-delldock.rules` with the following content:
```bash
ACTION=="add", ATTR{idVendor}=="0bda", ATTR{idProduct}=="402e", RUN="/bin/sh -c 'echo 1 >/sys/\$devpath/remove'"
```

When the device is connected it will be automatically removed and your system won't know about it
