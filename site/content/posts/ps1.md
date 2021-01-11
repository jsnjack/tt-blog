---
title: "Add exit code to the command line"
date: 2021-01-11T22:36:02+01:00
draft: false
tags: ["linux", "bash"]
---
Some examples to add exit code to command line prompt in bash (put it into ~/.bashrc):
```bash
export PS1="\[\033[01;31m\]\$([ \$? == 0 ] || echo \"!\$? \" )\[\033[00m\]\[\033[01;32m\]\t \[\033[01;34m\]\w\[\033[00m\]\[\033[1;32m\]\n\$ \[\033[m\]"
```
Add to the default `PS1` in CentOS:
```bash
export PS1="\[\033[01;31m\]\${?##0}\[\033[00m\][\u@\h \W]\\$"
```
