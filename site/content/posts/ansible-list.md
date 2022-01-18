---
title: "Ansible: List all vars of the host"
date: 2022-01-18T14:57:39+01:00
draft: false
tags: ["linux", "ansible"]
---

```bash
ansible stag-s4 -m debug -a "var=hostvars[inventory_hostname]"
```
