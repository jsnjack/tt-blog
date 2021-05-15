---
title: "Connect dedicated servers with cloud servers in hetzner via vSwitch"
date: 2021-05-15T18:05:06+02:00
draft: false
tags: ["linux", "network", "nmcli", "hetzner", "centos"]
---

Main article is provided by Hetzner team [here](https://docs.hetzner.com/cloud/networks/connect-dedi-vswitch/). This article contains instructions for creating `vlan` interface in Centos 8 using `nmcli`.

Assumptions (same us in the main article + listed below):
 - VLAN ID is `4001`
 - parent network interface is `enp195s0`
 
Steps:
1. Create new connection:
```bash
nmcli connection add type vlan con-name vlan4001 ifname vlan4001 vlan.parent enp195s0 vlan.id 4001
```
2. Configure the connection:
```bash
nmcli connection modify vlan4001 802-3-ethernet.mtu 1400
nmcli connection modify vlan4001 ipv4.addresses '10.0.1.2/24'
nmcli connection modify vlan4001 ipv4.gateway '10.0.1.1'
nmcli connection modify vlan4001 ipv4.dns '10.0.0.4'  # (optional)
nmcli connection modify vlan4001 ipv4.method manual
nmcli connection modify vlan4001 +ipv4.routes "10.0.0.0/16 10.0.1.1"
```
3. Restart the connection
```bash
nmcli connection down vlan4001
nmcli connection up vlan4001
```
4. Verify configuration
```bash
# Prints what gateway is used to reach the ip
ip route get 10.0.0.5

# Print all connection information
nmcli connection show vlan4001

# Print routing table
ip r

# Use tui interface for NetworkManager
dnf install NetworkManager-tui
nmtui
```

> Restarting NetworkManager wasn't enough to apply custom routes. Bring interface up and down

And the link to the great [RedHat documentation](https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/8/html/configuring_and_managing_networking/configuring-vlan-tagging_configuring-and-managing-networking)

