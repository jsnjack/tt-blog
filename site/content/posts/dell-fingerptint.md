---
title: "Set up Dell fingerprint reader on Linux"
date: 2025-12-20T13:06:48+01:00
draft: false
tags: ["linux", "dell", "fingerprint", "fresh system"]
---

Dell laptops come with a different fingerprint readers. Some of them work out of
the box in Fedora and that was my expierince with Precision series. However,
my new Dell Pro Max 16 MC16255 came with a different unsupported fingerprint.

Find your fingerprint reader

```bash
sudo lsusb
```

You should see something like this:

```
Bus 003 Device 002: ID 0a5c:5865 Broadcom Corp. 58200
```

Where `5865` is the model of the fingerprint reader. There are some copr repositories
around that provide support for some of the models, but not for `5865`.

There is a good [article on Reddit](https://www.reddit.com/r/Fedora/comments/1ob4s51/a_solution_to_broadcom_fingerprint)
that explains how to get it working.

The solution is to manually install the released driver from Dell for Ubuntu.

Go to [Dell support site](http://dell.archive.canonical.com/updates/pool/public/libf/libfprint-2-tod1-broadcom-cv3plus/)
and download the latest `libfprint-2-tod1-broadcom-cv3plus` `orig` [package](http://dell.archive.canonical.com/updates/pool/public/libf/libfprint-2-tod1-broadcom-cv3plus/libfprint-2-tod1-broadcom-cv3plus_6.3.299-6.3.040.0.orig.tar.gz).

Extract it and move:
```
../usr/lib/x86_64-linux-gnu/libfprint-2/tod-1/libfprint-2-tod-1-broadcom-cv3plus.so
```
to
```
../usr/lib64/libfprint-2/tod-1/libfprint-2-tod-1-broadcom-cv3plus.so
```

One-liner:
```bash
mkdir -p ../usr/lib64/libfprint-2/tod-1/ && mv ../usr/lib/x86_64-linux-gnu/libfprint-2/tod-1/libfprint-2-tod-1-broadcom-cv3plus.so $_
```

Run the included install script:
```bash
sudo ./install.sh
```

Replace `libfprint` package:
```bash
sudo dnf copr enable quantt/libfprint-tod -y
sudo dnf swap libfprint libfprint-tod -y
```

Programmatically add your fingerprint:
```bash
sudo setenforce 0
sudo systemctl restart fprintd
fprintd-enroll
sudo ausearch -m avc,user_avc,selinux_err -ts today | sudo audit2allow -M fprintd_SELinux
sudo semodule -i fprintd_SELinux.pp
sudo setenforce 1
```

To enable fingerptint authentication:
```bash
sudo authselect enable-feature with-fingerprint
sudo authselect apply-changes
```

Fingerprint configuration should be also available in GNOME Settings under "Users".
