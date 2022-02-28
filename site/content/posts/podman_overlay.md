---
title: "Unable to start Podman containers in LXD"
date: 2022-02-05T21:35:51+01:00
draft: false
tags: ["linux", "lxd", "podman"]
---

```bash
[root@fe ~]# podman run hello-world
ERRO[0000] 'overlay' is not supported over extfs at "/var/lib/containers/storage/overlay" 
Error: kernel does not support overlay fs: 'overlay' is not supported over extfs at "/var/lib/containers/storage/overlay": backing file system is unsupported for this graph driver
```

To fix this error, you'll need to edit `/etc/containers/storage.conf` file and make sure that it contains the following options:
```
[storage]
driver = "overlay"
[storage.options]
mount_program = "/usr/bin/fuse-overlayfs"
```

Run the following command to verify the result:
```bash
podman --log-level=debug ps -a
```
