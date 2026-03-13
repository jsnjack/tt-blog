---
title: "Disable alien network audio devices in PipeWire"
date: 2026-03-13T21:29:33+01:00
draft: false
tags: ["linux", "pipewire", "audio"]
---

By default, PipeWire discovers and shows AirPlay audio devices from other computers on your network. This clutters your audio device list with devices you don't need.

Disable RAOP (AirPlay audio) module in PipeWire. Create `~/.config/pipewire/pipewire.conf.d/noraop.conf`:

```
context.properties = {
        module.raop = false
}
```

Restart PipeWire:
```bash
systemctl --user restart pipewire pipewire-pulse
```

Network audio devices from other computers will no longer appear in your audio settings.
