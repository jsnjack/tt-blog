---
title: "Clear file from leading new line and etc"
date: 2021-02-26T15:57:44+01:00
draft: false
tags: ["linux", "bash"]
---

```bash
cat file | tr -d [$'\t\r\n'] > new_file
```
