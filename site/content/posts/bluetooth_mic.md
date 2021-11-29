---
title: "Improve Bluetooth headset audio quality"
date: 2021-11-30T00:29:06+01:00
draft: false
tags: ["linux", "bluetooth", "headphones"]
---

In order to enable mic on a bluetooth headphones, you need to change device 
configuration to `Headset Head Unit` in `Settings > Sound` menu. However, if you do
so, the audio quality of the headphones will be quite bad.

Pipewire introduced support for the mSBC codec. This codec allows your headphones
to enable microphone and still have okay-ish audio quality.

mSBC codec is disabled by default. To enable it, create `~/.config/pipewire/media-session.d/bluez-monitor.conf`
file with the following content:
```
properties = {
    bluez5.msbc-support = true
}
```

Restart Pipewire:
```
systemctl --user restart pipewire.service
```

`HSP/HFP, codec mSBC` option should be available in `Settings > Sound` menu
