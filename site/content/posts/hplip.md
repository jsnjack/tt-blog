---
title: "Configure HP printer in fedora"
date: 2021-01-29T20:14:29+01:00
draft: false
tags: ["linux", "fedora", "hp", "printer"]
---

Default `Printer` application in GNOME (fedora 33) is not able to successfully configure wireless printer, it needs to be done manually:
```bash
sudo dnf install hplip hplip-gui
sudo hp-setup
```
Select `Network/Ethernet/Wireless network` and your printer will appear in the list. Follow the instructions and add the printer.
