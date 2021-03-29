---
title: "Print repository usage"
date: 2021-03-29T23:20:46+02:00
draft: false
tags: ["linux", "dnf"]
---

Print how many packages are installed per repository:
```
sudo dnf list --installed | grep -E -o "@.*" | sort | uniq -c
```
