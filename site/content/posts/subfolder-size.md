---
title: "Get size of the subfolders"
date: 2017-05-01T23:20:19+01:00
draft: false
tags: ["linux", "bash", "filesystem"]
---

Use this command to print size of the subfolders:
```bash
du -sh *
```
If you want to exclude Permission denied errors, for example:
```bash
du -sh * 2>/dev/null
```
