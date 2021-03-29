---
title: "Dell Precision 5540 - Battery drains in sleep mode"
date: 2020-01-01T23:41:22+01:00
draft: false
tags: ["linux", "dell", "laptop", "battery"]
---

Run the following command to check the suspend mode:
```bash
cat /sys/power/mem_sleep
```
If the output looks like this:
```c
[s2idle] deep
```
It means that very inefficient suspend mode is active (drains battery from 100% to 0 in 9 hours). To enable "deep" mode, edit file `/etc/default/grub` and append `mem_sleep_default=deep` to `GRUB_CMDLINE_LINUX`:
```c
GRUB_TIMEOUT=5
GRUB_DISTRIBUTOR="$(sed 's, release .*$,,g' /etc/system-release)"
GRUB_DEFAULT=saved
GRUB_DISABLE_SUBMENU=true
GRUB_TERMINAL_OUTPUT="console"
GRUB_CMDLINE_LINUX="rd.lvm.lv=fedora/root rd.lvm.lv=fedora/swap rhgb quiet mem_sleep_default=deep"
GRUB_DISABLE_RECOVERY="true"
GRUB_ENABLE_BLSCFG=true
```
Apply changes:
```bash
sudo grub2-mkconfig -o /boot/efi/EFI/fedora/grub.cfg
```

Restart your laptop and verify changes:
```bash
$ cat /sys/power/mem_sleep
s2idle [deep]
```
