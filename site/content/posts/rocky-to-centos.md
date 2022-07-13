---
title: "Migrate from Rocky Linux to CentOS Stream"
date: 2022-07-13T14:37:01+02:00
draft: false
tags: ["linux", "centos", "rocky linux"]
---

There is no official guide to migrate from Rocky Linux to CentOS Stream, however, it can be achieved by using official CentOS to CentOS Stream migration guide

1. Install centos-release-stream package
```bash
dnf install https://vault.centos.org/centos/8/extras/x86_64/os/Packages/centos-release-stream-8.1-1.1911.0.7.el8.x86_64.rpm
```

2. The next step is to swap Rocky Linux repositories to CentOS Stream ones. This is done with `dnf swap centos-{linux,stream}-repos -y` command. However, this command will cause multiple conflicts in Rocky Linux. To solve them:
  - copy `/etc/pki/rpm-gpg/RPM-GPG-KEY-centosofficial` file from any Centos Stream system into Rocky Linux
  - remove conflicting packages:
    ```bash
    rpm -e --nodeps rocky-repos
    rpm -e --nodeps rocky-release
    ```
  - complete installing CentOS Stream repositories
    ```bash
    dnf swap centos-{linux,stream}-repos -y
    ```

3. Switch to CentOS Stream
```bash
dnf distro-sync -y
```

4. Verify installation and reboot the server:
```bash
cat /etc/centos-release
reboot
```
