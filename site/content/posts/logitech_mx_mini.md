---
title: "Fix missing shortcuts in Logitech MX Mechanical Mini in Linux"
date: 2024-02-26T22:26:47+01:00
draft: false
tags: ["linux", "logitech"]
---

Logitech MX Mechanical Mini keyboard works mostly great in Linux. However, there are
few buttons which don't: Mic mute, Screenshots, Emoji, Voice dictation. To fix it,
install [solaar](https://github.com/pwr-Solaar/Solaar) and reboot your computer (important,
otherwise the app will hang):
```bash
sudo dnf install solaar
reboot
```

Configuring every button is 2-step process.

On the first step, the button is marked as a custom button (Diverted):
{{< img "solaar_1.png" >}}

On the second step, click `Rule Editor` in the bottom left corner and there you
can assign a new action to it:
{{< img "solaar_2.png" >}}

Solaar keeps configuration files in `~/.config/solaar/` directory. Rules are stored
in `~/config/solaar/rules.yaml` file:
```yaml
%YAML 1.3
---
- Key: [MultiPlatform Search, pressed]
- KeyPress:
  - XF86_AudioNext
  - click
...
---
- Key: [Voice Dictation, pressed]
- Execute: [firefox, 'https://gemini.google.com']
...
---
- Key: [Snipping Tool, pressed]
- KeyPress:
  - Print
  - click
...
---
- Key: [Open Emoji Panel, pressed]
- Execute: gnome-characters
...
---
- Key: [Mute Microphone, pressed]
- KeyPress:
  - XF86_AudioMicMute
  - click
...
```
The configuration file is located in `~/.config/solaar/config.yaml` and is unique
for every keyboard. It contains the keyboard's name, serial number and etc. Our
modified configuration should look like this:
```yaml
- 1.1.10
- _NAME: MX Mechanical Mini
  _absent: [hi-res-scroll, lowres-scroll-mode, hires-smooth-invert, hires-smooth-resolution, hires-scroll-mode, scroll-ratchet, smart-shift, thumb-scroll-invert,
    thumb-scroll-mode, onboard_profiles, report_rate, pointer_speed, dpi, speed-change, backlight-timed, reprogrammable-keys, persistent-remappable-keys,
    crown-smooth, divert-crown, divert-gkeys, m-key-leds, mr-key-led, gesture2-gestures, gesture2-divert, gesture2-params, sidetone, equalizer, adc_power_management]
  # <redacted>
  _sensitive: {divert-keys: true, multiplatform: false}
  # <redacted>
  backlight: true
  change-host: null
  disable-keyboard-keys: {1: false, 4: false, 8: false, 16: false}
  divert-keys: {212: 1, 226: 0, 227: 0, 231: 0, 232: 0, 233: 0, 259: 1, 264: 1, 266: 1, 267: 0, 268: 0, 269: 0, 270: 0, 271: 0, 272: 0, 273: 0, 274: 0,
    277: 0, 279: 0, 280: 0, 281: 0, 282: 0, 283: 0, 284: 1, 286: 0, 316: 0, 321: 0}
  fn-swap: false
  multiplatform: 0
```
