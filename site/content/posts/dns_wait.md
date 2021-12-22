---
title: "Wait until DNS propagates"
date: 2021-12-22T15:50:55+01:00
draft: false
tags: ["linux", "dns", "bash"]
---

```bash
until [[ -n $(dig new.mydomain.com A +short @1.1.1.1) ]]; do echo "waiting..." && sleep 5; done
```

`-n` flag stands for non-empty string
