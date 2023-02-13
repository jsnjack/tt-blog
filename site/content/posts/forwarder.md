---
title: "Make port available as unix socket"
date: 2023-02-13T23:06:06+01:00
draft: false
tags: ["linux", "socat"]
---

```bash
socat UNIX-LISTEN:/var/run/hapee-lb.sock,mode=666,fork TCP-CONNECT:127.0.0.1:1936
```

Will create `/var/run/hapee-lb.sock` file which forwards all input to `127.0.0.1:1936`. `fork` parameter
will allow it to handle connection close and accept more than 1 request
