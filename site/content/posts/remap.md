---
title: "Remap the key between left shift and Z in Logitech MX Mechanical Mini in Linux"
date: 2024-05-20T11:45:44+02:00
draft: false
tags: ["linux", "logitech"]
---

In Europe, Logitech only sells keyboards with ISO layout. This means that there
is an additional key between the left shift and Z. This key is often in the way
if you are used to the US layout (ANSI layout). [Solaar](./logitech_mx_mini.md)
does not allow you to remap this key.

1. Install evtest:
```bash
sudo dnf install evtest
```

2. Start it and find the device:
```bash
sudo evtest
```

`evtest` will output all available devices. For example:
```
No device specified, trying to scan all of /dev/input/event*
Available devices:
/dev/input/event0:	Sleep Button
/dev/input/event1:	Power Button
...
/dev/input/event27:	solaar-keyboard
/dev/input/event3:	Logitech USB Receiver
/dev/input/event4:	Logitech USB Receiver Mouse
/dev/input/event5:	Logitech USB Receiver Consumer Control
/dev/input/event6:	Logitech USB Receiver System Control
```

The keyboard is `solaar-keyboard`. However, it is connected via the USB receiver.
This means that the device is `/dev/input/event3`.

3. Record the key code by pressing the key:
```bash
Event: time 1692951945.797570, type 4 (EV_MSC), code 4 (MSC_SCAN), value 70064
Event: time 1692951945.797570, type 1 (EV_KEY), code 86 (KEY_102ND), value 1
```

> If you selected `solaar-keyboard` as the device, `evtest` won't record any key

4. Create a new rule:
```bash
sudo vim /etc/udev/hwdb.d/mx-keys.hwdb
```

Add the following content:
```bash
evdev:name:Logitech USB Receiver:*
  KEYBOARD_KEY_70064=leftshift
```

5. Update the rules to test the new configuration:
```bash
sudo systemd-hwdb update
sudo udevadm trigger
```
