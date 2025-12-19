---
title: "Tilix configuration"
date: 2025-12-19T23:26:25+01:00
draft: false
tags: ["linux", "tilix", "terminal", "fresh system"]
---

Tilix configuration:
```
[/]
accelerators-enabled=true
enable-wide-handle=true
prompt-on-close=true
prompt-on-close-process=false
quake-specific-monitor=0
terminal-title-style='none'
theme-variant='dark'
unsafe-paste-alert=false
window-style='normal'

[keybindings]
app-new-session='F2'
app-new-window='<Shift><Ctrl>N'
app-preferences='disabled'
app-shortcuts='disabled'
nautilus-open='<Ctrl><Alt>t'
session-add-auto='F1'
session-add-down='<Shift>F2'
session-add-right='<Primary>F2'
win-reorder-next-session='<Primary><Alt>Page_Up'
win-reorder-previous-session='<Primary><Alt>Page_Down'
win-switch-to-next-session='<Primary>Right'
win-switch-to-previous-session='<Primary>Left'

[profiles/2b7c4080-0ddd-46c5-8f23-563fd3ba789d]
background-color='#1C1C1F'
badge-color-set=false
bold-color-set=false
cursor-colors-set=false
foreground-color='#FFFFFF'
highlight-colors-set=false
palette=['#241F31', '#C01C28', '#2EC27E', '#F5C211', '#1E78E4', '#9841BB', '#0AB9DC', '#C0BFBC', '#5E5C64', '#ED333B', '#57E389', '#F8E45C', '#51A1FF', '#C061CB', '#4FD2FD', '#F6F5F4']
use-theme-colors=false
visible-name='Default'
```

To apply this configuration, save it to a file and import like so:
```
dconf load /com/gexperts/Tilix/ < tilix-settings.conf
```
