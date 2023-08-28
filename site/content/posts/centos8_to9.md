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

3. Rebuild RPM database (this will chnage the backend to sqlite):
```bash
sudo rpm --rebuilddb
```

4. Disable 8-specific modules (they are blocking kernel updates)
```bash
sudo dnf module disable python36 virt
```

5. Disable subscription manager. Open file `/etc/yum/pluginconf.d/subscription-manager.conf` and set enabled to `0`

6. Reboot and verify
```bash
cat /etc/redhat-release
```

7. Verify that the latest kernel is used
```bash
uname -a
```

8. If not, use grubby to set the latest kernel as the default one, reboot the system and remove old kernels from CentOS 8:
```bash
# List all boot options
grubby --info=ALL

# Reflect the desired kernel in configuration
grubby --set-default vmlinuz-<version>.<arch>

# Make sure the index is also set in `/etc/default/grub` file

# Regenerate boot configuration
sudo grub2-mkconfig -o /boot/grub2/grub.cfg

sudo reboot

# Remove old kernels...
```

> The instructions are inspired by CentOS 8 to CentOS Stream 8 migration guide and Fedora upgrade procedure
