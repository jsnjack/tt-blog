---
title: "Screen artifcacts with AMD GPU on Linux"
date: 2025-12-20T15:51:41+01:00
draft: false
tags: ["linux", "amd", "gpu", "dell", "fresh system"]
---

Apparently, it is quite common for AMD GPUs to show screen artifacts on Linux.
Most of the times it is related to power management settings. These settings
can be adjusted via `amdgpu` kernel module options. The parameters are different
for different GPUs. In my case, I have a Dell laptop with AMD Radeon 890M GPU.

To fix the screen artifacts, run the following commands:

```bash
sudo grubby --update-kernel=ALL --args="amdgpu.dcdebugmask=0x410"
```

Reboot your system.
