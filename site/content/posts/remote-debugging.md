---
title: "Remote debugging with headless Chrome"
date: 2021-11-23T15:11:54+01:00
draft: false
tags: ["linux", "chrome", "selenium"]
---

Add the following arguments to chrome options:
```python
options = webdriver.ChromeOptions()
options.add_argument("headless")
options.add_argument("remote-debugging-address=0.0.0.0")
options.add_argument("remote-debugging-port=9222")
wd = webdriver.Chrome(chrome_options=options)
```

 - start Chrome and navigate to `chrome://inspect` page
 - make sure `Discover network targets` is checked
 - click `Configure` button next to the previous option
 - add a new target, for example (must be an IP address): 10.214.36.58:9222
 - the remote browser should be automatically discovered and appear in `Remote target` list
 - click `Inspect`
