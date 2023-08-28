---
title: "Upgrade CentOS Stream 8 to CentOS Stream 9"
date: 2023-08-28T15:39:12+02:00
draft: false
tags: ["linux", "centos"]
---

1. Install CentOS 9 repositories. All CentOS 9 packages are listed [here](https://mirror.stream.centos.org/9-stream/BaseOS/x86_64/os/Packages/)
```bash
sudo dnf install https://mirror.stream.centos.org/9-stream/BaseOS/x86_64/os/Packages/centos-stream-release-9.0-22.el9.noarch.rpm https://mirror.stream.centos.org/9-stream/BaseOS/x86_64/os/Packages/centos-gpg-keys-9.0-22.el9.noarch.rpm https://mirror.stream.centos.org/9-stream/BaseOS/x86_64/os/Packages/centos-stream-repos-9.0-22.el9.noarch.rpm
```

2. Run command to switch packages:
```bash
sudo dnf --releasever=9 --allowerasing --setopt=deltarpm=false distro-sync -y
```

3. Reboot and verify
```bash
cat /etc/redhat-release
```

> The instructions are inspired by CentOS 8 to CentOS Stream 8 migration guide and Fedora upgrade procedure
