---
title: "Running Linux on Dell Pro Max 16 MC16255"
date: 2026-03-13T21:29:33+01:00
draft: false
tags: ["linux", "dell", "amd", "kernel", "fresh system"]
---

Collection of kernel and configuration tweaks to get the best experience on Linux with the Dell Pro Max 16 MC16255 laptop with AMD Radeon 890M GPU.

See also: [Fingerprint reader setup](../dell-fingerptint), [Video codecs](../codecs-amd)

Add kernel parameters:
```bash
sudo grubby --update-kernel=ALL --args="amdgpu.dcdebugmask=0x410 amdgpu.sg_display=0"
sudo reboot
```

What each parameter does:

| Parameter | Description |
|-----------|-------------|
| `amdgpu.dcdebugmask=0x410` | Fixes screen artifacts and flickering |
| `amdgpu.sg_display=0` | Fix cursor stuttering and video playback stuttering every few seconds |

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

