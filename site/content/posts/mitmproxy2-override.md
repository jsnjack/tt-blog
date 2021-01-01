---
title: "Modify response body with mitmproxy 2"
date: 2017-07-01T23:22:39+01:00
draft: false
tags: ["linux", "mitmproxy"]
---

Run python script with mitmproxy:
```bash
mitmproxy --ignore ws.sitesupport.net:443 -s script.py
```
Script example:
```python
from mitmproxy.script import concurrent

OLD = """var totalExpGems=0;"""
NEW = """var totalExpGems=0;debugger;"""


@concurrent
def response(flow):
    if "gem_finder" in flow.request.path:
        flow.response.headers["XX"] = "PATCHED"
        body = flow.response.content.decode("utf-8")
        if OLD in body:
            flow.response.content = body.replace(OLD, NEW).encode("utf-8")
            flow.response.headers["XXX"] = "PATCHED"
```
