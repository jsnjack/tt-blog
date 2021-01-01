---
title: "fedora: Lockdown a kernel version"
date: 2020-01-01T23:39:58+01:00
draft: false
tags: ["linux", "fedora", "dnf"]
---

Fedora typically keeps 3 latest kernel versions. However, what if you want to keep an old specific kernel version which is known to work well with your system? First, we need to install a dnf plugin that allows us to keep a specific version of any package:
```bash
sudo dnf install python3-dnf-plugins-extras-versionlock
```

Lets find all available kernel versions for your system:
```bash
sudo dnf list kernel --showduplicates
```

Install that specific version of the kernel:
```bash
sudo dnf install kernel-5.3.7-301.fc31
```

And lock it:
```bash
sudo dnf versionlock add kernel-5.3.7-301.fc31
```
To remove the lock:
```bash
sudo dnf versionlock delete kernel-5.3.7-301.fc31
```
