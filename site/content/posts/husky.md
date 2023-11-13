---
title: "Use Github Desktop with pre-commit hooks with husky"
date: 2022-09-05T16:09:23+02:00
draft: false
tags: ["linux", "git", "husky"]
---

If your repository has an installed pre-commit hook which uses husky or anything similar and you use Github Desktop for linux installed via flatpak it won't work. The reason for this is that flatpak environment doesn't have all necessary dependencies (npx, husky and etc.). There is a way to mount parts of your host file system to the flatpak environment (with `--filesystem` flag), but it excludes all sensitive folders (including `/usr/bin/`). A possible solution would be to mount the content of `/usr/bin/` to a custom folder in flatpak environment and override `PATH` environmental variable. But the quick fix for now is to just disable it:
```
flatpak override --user --env=HUSKY=0 io.github.shiftey.Desktop
```
