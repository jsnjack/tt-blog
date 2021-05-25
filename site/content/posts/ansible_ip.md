---
title: "Run Ansible playbook on remote IP address"
date: 2021-05-25T17:01:20+02:00
draft: false
tags: ["linux", "ansible"]
---

```bash
ansible-playbook my.yml -i 162.55.82.217, -e "ansible_user=root"
```

The trick is to add `,` after the IP address
