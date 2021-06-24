---
title: "Tune kernel parameters with sysctl"
date: 2021-06-24T22:40:30+02:00
draft: false
tags: ["linux", "podman", "sysctl"]
---

Lets say there is this error when using rootless containers with podman:
```
Error: rootlessport cannot expose privileged port 80, you can add 'net.ipv4.ip_unprivileged_port_start=80' to /etc/sysctl.conf (currently 1024), or choose a larger port number (>= 1024): listen tcp 0.0.0.0:80: bind: permission denied
```

To solve this problem, the value of `net.ipv4.ip_unprivileged_port_start` needs to be changed

Print current value:
```bash
sudo sysctl net.ipv4.ip_unprivileged_port_start
```

Print all configuration:
```bash
sudo sysctl -a
```

Temporarily change the value:
```bash
sudo sysctl -w net.ipv4.ip_unprivileged_port_start=80
```

To permanently modify the value, create a new file in `/etc/sysctl.d`. To apply changes, either reboot or execute `sudo sysctl -p /etc/sysctl.d/99-custom.conf`
