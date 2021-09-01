---
title: "Kernel panic after update"
date: 2021-09-01T23:16:58+02:00
draft: false
tags: ["linux", "rhel", "centos", "kernel"]
---

> Traceback example: After update, kernel panic at boot with error: Unable to mount root fs on unknown-block(0,0)

##### Solution:
1. Restart your server and on the boot screen select any other boot option with older kernel
2. (optional, Leaseweb servers) Install missing dependency
```bash
sudo dnf install dracut-network
```
3. Navigate to `/boot` and verify that initramfs-file doesn't exist for the latest kernel version
4. Generate it
```bash
# 4.18.0-338.el8.x86_64 is kernel version
dracut -f /boot/initramfs-4.18.0-338.el8.x86_64.img 4.18.0-338.el8.x86_64
```
5. Regenerate grub files
```bash
grub2-mkconfig -o /boot/grub2/grub.cfg
```
6. Restart

##### Sources:
 - [https://access.redhat.com/solutions/57018](https://access.redhat.com/solutions/57018)
 - [https://access.redhat.com/solutions/1958](https://access.redhat.com/solutions/1958)
 - [https://access.redhat.com/solutions/5672791](https://access.redhat.com/solutions/5672791)
