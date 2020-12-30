---
title: "Fedora 33: font rendering"
date: 2020-12-30T22:41:58+01:00
draft: false
tags: ["linux", "fedora", "fedora33", "font"]
---

Packages from this [repository](https://copr.fedorainfracloud.org/coprs/dawid/better_fonts/) improve font rendering in fedora a lot:
```bash
sudo dnf copr enable dawid/better_fonts -y
sudo dnf install fontconfig-font-replacements fontconfig-enhanced-defaults levien-inconsolata-fonts -y
gsettings set org.gnome.desktop.interface document-font-name "Cantarell 11"
gsettings set org.gnome.desktop.interface font-name "Cantarell 11"
gsettings set org.gnome.desktop.interface monospace-font-name "Inconsolata 13"
``` 
