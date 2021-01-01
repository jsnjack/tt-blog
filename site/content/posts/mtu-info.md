---
title: "Check and modify mtu of a connection"
date: 2020-01-01T23:36:09+01:00
draft: false
tags: ["linux", "network"]
---

To check:
```bash
[root@lxd-node3 ~]# netstat -i
Kernel Interface table
Iface             MTU    RX-OK RX-ERR RX-DRP RX-OVR    TX-OK TX-ERR TX-DRP TX-OVR Flg
ens10            1450      455      0      0 0           671      0      0      0 BMRU
eth0             1500       28      0      0 0            32      0      0      0 BMRU
lo              65536        0      0      0 0             0      0      0      0 LRU
```

Modify temporarily:
```bash
ifconfig eth0 mtu 1450
```
Permanently: edit `/etc/sysconfig/network-scripts/ifcfg-eth0` file and add:
```
MTU=1450
```
> Might be the reason why the system can establish HTTP connection but HTTPS
