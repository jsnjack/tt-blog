---
title: "top-like application for intel GPU"
date: 2021-01-18T22:39:23+01:00
draft: false
tags: ["linux", "gpu", "youtube", "intel"]
---
Top-like application to monitor performance of intel GPU:
```bash
sudo dnf install igt-gpu-tools
sudo intel_gpu_top
```

Playing 4k [video sample](https://www.youtube.com/watch?v=A_hF37w6Uao&t=293s) on youtube (intel i3-8100T):
```
intel-gpu-top: Intel Coffeelake (Gen9) @ /dev/dri/card0 - 1063/1083 MHz;    1% RC6;  7.35/19.18 W;      467 irqs/s

      IMC reads:     6889 MiB/s
     IMC writes:     3901 MiB/s

          ENGINE      BUSY                                      MI_SEMA MI_WAIT
     Render/3D/0   98.29% |██████████████████████████████████▍|      0%      0%
       Blitter/0    0.00% |                                   |      0%      0%
         Video/0    0.00% |                                   |      0%      0%
  VideoEnhance/0    0.00% |                                   |      0%      0%
```

Playing the same video in 1080p (intel i3-8100T)::
```
intel-gpu-top: Intel Coffeelake (Gen9) @ /dev/dri/card0 -  591/ 591 MHz;   30% RC6;  3.06/ 6.47 W;      474 irqs/s

      IMC reads:     4058 MiB/s
     IMC writes:     2654 MiB/s

          ENGINE      BUSY                                      MI_SEMA MI_WAIT
     Render/3D/0   55.64% |███████████████████▍               |      0%      0%
       Blitter/0    0.00% |                                   |      0%      0%
         Video/0    0.00% |                                   |      0%      0%
  VideoEnhance/0    0.00% |                                   |      0%      0%
```
