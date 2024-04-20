---
title: "Monitor IO disk utilization"
date: 2022-05-10T13:56:35+02:00
draft: false
tags: ["linux", "io", "disk", "plot"]
---

Install `sysstat` package:
```bash
sudo dnf install sysstat -y
```

Install `ttyplot` to plot data and `kazy` to extract data with [grm](https://github.com/jsnjack/grm/)
```bash
grm install jsnjack/kazy-go
grm install tenox7/ttyplot==1.6.1 -n ttyplot -l
```

Check available disk partitions with `df` command. nvme0n1 is used in example:
```bash
iostat -dx 1 | kazy -i nvme0n1 | kazy -r -x "[\d.]*$" | ttyplot -s 100
```
