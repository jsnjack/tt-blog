---
title: "Dell 16 Max Kernel Tweaks"
date: 2026-03-13T21:29:33+01:00
draft: false
tags: ["linux", "dell", "amd", "kernel", "fresh system"]
---

My Dell Pro Max 16 MC16255 with AMD Radeon 890M GPU requires some kernel parameters to work reliably on Linux. Without these, you'll get screen artifacts and suspend issues with audio and external USB-C screens.

See also: [AMD GPU artifacts](../amd_gpu), [Fingerprint setup](../dell-fingerptint), [Video codecs](../codecs-amd)

Apply all parameters:
```bash
sudo grubby --update-kernel=ALL --args="amdgpu.dcdebugmask=0x410 amdgpu.dcfeaturemask=0x2 usbcore.autosuspend=-1 snd_hda_intel.power_save=0 amdgpu.vpe_enabled=0 pcie_aspm=off"
sudo reboot
```

What each parameter does:

| Parameter | Description |
|-----------|-------------|
| `amdgpu.dcdebugmask=0x410` | Disables AMD Display Core features that cause screen artifacts and flickering |
| `amdgpu.dcfeaturemask=0x2` | Enables only basic display functionality, disables experimental features |
| `usbcore.autosuspend=-1` | Disables USB autosuspend to fix external USB-C screen issues after suspend/resume |
| `snd_hda_intel.power_save=0` | Disables audio power management to fix audio not working after suspend |
| `amdgpu.vpe_enabled=0` | Disables Video Processing Engine which can cause video playback issues |
| `pcie_aspm=off` | Disables PCIe power management to fix suspend/resume issues with PCIe devices |

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

