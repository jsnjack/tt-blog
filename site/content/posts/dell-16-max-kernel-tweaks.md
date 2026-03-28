---
title: "Running Linux on Dell Pro Max 16 MC16255"
date: 2026-03-13T21:29:33+01:00
draft: false
tags: ["linux", "dell", "amd", "kernel", "fresh system"]
---

My Dell Pro Max 16 MC16255 with AMD Radeon 890M GPU requires some kernel parameters to work reliably on Linux. Without these, you'll get screen artifacts and suspend issues with audio and external USB-C screens.

See also: [[Fingerprint reader setup](../dell-fingerptint), [Video codecs](../codecs-amd)

Apply all parameters:
```bash
sudo grubby --update-kernel=ALL --args="amdgpu.dcdebugmask=0x410 amdgpu.sg_display=0"
sudo reboot
```

What each parameter does:

| Parameter | Description |
|-----------|-------------|
| `amdgpu.dcdebugmask=0x410` | Disables AMD Display Core features that cause screen artifacts and flickering |
| `amdgpu.sg_display=0` | Disables scatter-gather display to fix cursor stuttering and video playback stuttering every few seconds |

Additionally, disable audio interface suspension in Wireplumber. Create `~/.config/wireplumber/wireplumber.conf.d/51-disable-suspension.conf`:

```
monitor.alsa.rules = [
  {
    matches = [ { node.name = "~alsa_output.*" } ],
    actions = { update-props = { session.suspend-timeout-seconds = 0 } }
  }
]
```

Restart Wireplumber:
```bash
systemctl --user restart wireplumber
```

