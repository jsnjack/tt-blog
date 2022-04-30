---
title: "Remove unused runtimes in flatpak"
date: 2022-04-30T13:25:20+02:00
draft: false
tags: ["linux", "flatpak"]
---

```bash
flatpak uninstall --unused
```

Maintenance commands for flatpak:
```bash
flatpak update && flatpak uninstall --unused
```
