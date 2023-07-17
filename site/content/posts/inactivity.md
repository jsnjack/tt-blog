---
title: "rygel: Disable inactivity timeout when streaming video"
date: 2023-07-17T23:08:46+02:00
draft: false
tags: ["linux", "rygel", "bash"]
---

This script demonstrates how to disable system inactivity timeout in Gnome when running Rygel and
restore it to the original value after exiting:
```bash
#!/bin/bash

# Save current inactivity timeout
TIMEOUT=$(gsettings get org.gnome.settings-daemon.plugins.power sleep-inactive-ac-timeout)
echo "Current timeout: $TIMEOUT"

# Execute commands on Ctrl + C
function ctrl_c() {
  echo "Exiting rygel..."
  gsettings set org.gnome.settings-daemon.plugins.power sleep-inactive-ac-timeout $TIMEOUT
  echo "Inactivity timeout restored to $TIMEOUT. Bye!"
  sleep 2
}

# Set the trap for Ctrl+C
trap ctrl_c INT

# Disable the inactivity timeout
gsettings set org.gnome.settings-daemon.plugins.power sleep-inactive-ac-timeout 0

# Briefly start rygel to allow it to add new files
timeout 2s rygel

# Run command to fix filenames
rygel-titlefix

# Actually run rygel
rygel

```
