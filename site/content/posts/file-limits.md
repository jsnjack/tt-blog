---
title: "Open file limits"
date: 2021-07-12T18:48:50+02:00
draft: false
tags: ["linux", "system", "systemd", "nofile", "ulimit"]
---

Print system limits:
```bash
14:53 $ ulimit -a
core file size          (blocks, -c) unlimited
data seg size           (kbytes, -d) unlimited
scheduling priority             (-e) 0
file size               (blocks, -f) unlimited
pending signals                 (-i) 514130
max locked memory       (kbytes, -l) 64
max memory size         (kbytes, -m) unlimited
open files                      (-n) 1024
pipe size            (512 bytes, -p) 8
POSIX message queues     (bytes, -q) 819200
real-time priority              (-r) 0
stack size              (kbytes, -s) 8192
cpu time               (seconds, -t) unlimited
max user processes              (-u) 514130
virtual memory          (kbytes, -v) unlimited
file locks                      (-x) unlimited
```

Print just open files limit
```bash
16:52 $ ulimit -Sn
1024
```

Get number of open files for a specific process:
```bash
# Get process id

16:56 $ ps aux | grep wshub
cobro    3200102  2.0  0.0 3676920 86488 ?       Ssl  16:12   0:54 /opt/cobro/src/wshub/wshub -port 8015 -redis 127.0.0.1:6379 -auxport 8016
client   3212070  0.0  0.0  12132  1160 pts/0    S+   16:57   0:00 grep --color=auto wshub

# Check current limits for the process

16:57 $ cat /proc/3200102/limits
Limit                     Soft Limit           Hard Limit           Units     
Max cpu time              unlimited            unlimited            seconds   
Max file size             unlimited            unlimited            bytes     
Max data size             unlimited            unlimited            bytes     
Max stack size            8388608              unlimited            bytes     
Max core file size        unlimited            unlimited            bytes     
Max resident set          unlimited            unlimited            bytes     
Max processes             514130               514130               processes 
Max open files            1024                 262144               files     
Max locked memory         65536                65536                bytes     
Max address space         unlimited            unlimited            bytes     
Max file locks            unlimited            unlimited            locks     
Max pending signals       514130               514130               signals   
Max msgqueue size         819200               819200               bytes     
Max nice priority         0                    0                    
Max realtime priority     0                    0                    
Max realtime timeout      unlimited            unlimited            us     

# Print number of open files for the process

17:03 $ sudo ls -l /proc/3200102/fd | wc -l
157
```

Systemd needs this limit specified per service (otherwise it will be 1024):
```bash

[Service]
LimitNOFILE=2048

```
