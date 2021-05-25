---
title: "Forward remote port to a local one via SSH tunnel"
date: 2021-05-25T16:56:04+02:00
draft: false
tags: ["linux", "ssh", "debug"]
---

```bash
ssh -L <LOCAL_PORT>:127.0.0.1:<REMOTE_PORT> remote_server
```
TCP connections to `<LOCAL_PORT>` will be forwarded to `127.0.0.1:<REMOTE_PORT>` on remote host `remote_server`
