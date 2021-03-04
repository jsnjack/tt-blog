---
title: "Remove host entry from known_hosts file"
date: 2021-03-04T13:40:21+01:00
draft: false
tags: ["linux", "ssh"]
---
To remove a host entry from `~/.ssh/known_hosts` you can do it manually by editing the file or use `ssh-keygen` application:
```bash
ssh-keygen -R 135.181.157.20
```

This should fix `Host key verification failed` error
