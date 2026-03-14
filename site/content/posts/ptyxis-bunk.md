---
title: "Ptyxis terminal setup"
date: 2026-03-14T12:00:00+01:00
draft: false
tags: ["linux", "fedora", "terminal", "fresh system"]
---

## Ptyxis

```bash
gsettings set org.gnome.Ptyxis cursor-blink-mode 'off'
dconf write /org/gnome/Ptyxis/Shortcuts/new-tab "'F2'"
dconf write /org/gnome/Ptyxis/Shortcuts/move-next-tab "'<Control>Right'"
dconf write /org/gnome/Ptyxis/Shortcuts/move-previous-tab "'<Control>Left'"
```
