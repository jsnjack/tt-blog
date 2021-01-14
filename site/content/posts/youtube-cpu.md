---
title: "High CPU usage on youtube in fedora"
date: 2021-01-14T22:41:33+01:00
draft: false
tags: ["linux", "cpu", "fedora", "firefox", "youtube"]
---
I've started to experience quite high CPU usage when watching youtube videos on 1440p screen ([video sample](https://www.youtube.com/watch?v=A_hF37w6Uao&t=293s)). By default, youtube loads videos in VP9 codec. Using this codec results in 80% CPU usage (i3-8100T). By installing this extension [enhanced-h264ify](https://addons.mozilla.org/en-US/firefox/addon/enhanced-h264ify/) and by blocking VP8, VP9 and AV1, youtube will load videos in avc1(h.264) codec. The CPU usage is down to about 30%.

Unfortunately, this will disable all 4k videos - they are available only in VP9.
