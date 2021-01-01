---
title: "Lxd on Fedora 31"
date: 2020-01-01T23:32:50+01:00
draft: false
tags: ["linux", "lxd", "fedora", "fedora31"]
---

Fedora 31 has started to use cgroups v2 by default. According to [Common F31 bugs](https://fedoraproject.org/wiki/Common_F31_bugs#Docker_package_no_longer_available_and_will_not_run_by_default_.28due_to_switch_to_cgroups_v2.29) it doesn't play nice with Docker, lxd and others. To make fedora switch back to the old cgroups:
```bash
sudo grubby --update-kernel=ALL --args="systemd.unified_cgroup_hierarchy=0"
```
