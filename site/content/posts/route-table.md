---
title: "Debug route table"
date: 2020-01-01T23:46:27+01:00
draft: false
tags: ["linux", "debug", "network"]
---

List all routes in system:
```bash
route -n
```
Get route for a specific IP:
```bash
ip route get 1.1.1.1
```
