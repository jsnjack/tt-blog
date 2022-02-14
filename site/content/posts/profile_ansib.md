---
title: "Profile Ansible playbooks"
date: 2022-02-14T14:23:55+01:00
draft: false
tags: ["linux", "ansible"]
---

Install profiling module:
```
ansible-galaxy collection install ansible.posix
```

Enable profiling in `ansible.cfg` in `[defaults]` section:
```
callback_whitelist = ansible.posix.profile_tasks
```
