---
title: "Keep running ssh command after disconnecting from the server"
date: 2021-02-12T12:21:51+01:00
draft: false
tags: ["linux", "ssh"]
---

```
nohup database_migration.sh &
```
 - `nohup` will take care of running the command and forward command output to `nohup.out`
 - `&` will run the command in the background
 
 Use `tail -f nohup.out` to get the output from the command
