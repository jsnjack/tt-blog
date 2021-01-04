---
title: "Disable touchpad in Microsoft All-in-one media keyboard"
date: 2021-01-04T21:36:59+01:00
draft: false
tags: ["linux", "keyboard", "touchpad"]
---
List all available input devices:
```bash
$ xinput list
⎡ Virtual core pointer                    	id=2	[master pointer  (3)]
⎜   ↳ Virtual core XTEST pointer              	id=4	[slave  pointer  (2)]
⎜   ↳ Logitech M185                           	id=10	[slave  pointer  (2)]
⎜   ↳ Microsoft Microsoft® Nano Transceiver v2.0 Consumer Control	id=14	[slave  pointer  (2)]
⎜   ↳ Microsoft Bluetooth Mouse               	id=20	[slave  pointer  (2)]
⎜   ↳ Microsoft Microsoft® Nano Transceiver v2.0 Consumer Control	id=13	[slave  pointer  (2)]
⎜   ↳ Microsoft Microsoft® Nano Transceiver v2.0 Mouse	id=12	[slave  pointer  (2)]
⎣ Virtual core keyboard                   	id=3	[master keyboard (2)]
    ↳ Virtual core XTEST keyboard             	id=5	[slave  keyboard (3)]
    ↳ Power Button                            	id=6	[slave  keyboard (3)]
    ↳ Video Bus                               	id=7	[slave  keyboard (3)]
    ↳ Power Button                            	id=8	[slave  keyboard (3)]
    ↳ Sleep Button                            	id=9	[slave  keyboard (3)]
    ↳ Microsoft Microsoft® Nano Transceiver v2.0	id=11	[slave  keyboard (3)]
    ↳ Microsoft Microsoft® Nano Transceiver v2.0 System Control	id=15	[slave  keyboard (3)]
    ↳ Dell WMI hotkeys                        	id=16	[slave  keyboard (3)]
    ↳ Logitech M185                           	id=17	[slave  keyboard (3)]
    ↳ Microsoft Microsoft® Nano Transceiver v2.0 Consumer Control	id=18	[slave  keyboard (3)]
    ↳ Microsoft Microsoft® Nano Transceiver v2.0 Consumer Control	id=19	[slave  keyboard (3)]
    ↳ Microsoft Bluetooth Mouse Keyboard      	id=21	[slave  keyboard (3)]

```
The touchpad device is known as `Microsoft Microsoft® Nano Transceiver v2.0` in the system. Disable it with this command:
```bash
xinput disable 12
```
And enable it again with this command:
```bash
xinput enable 12
```
Or one-liner:
```bash
xinput enable `xinput list | grep "Microsoft Microsoft® Nano Transceiver v2.0 Mouse" | grep -oP "id=[0-9]*" | grep -oP "[0-9]*"`
```
