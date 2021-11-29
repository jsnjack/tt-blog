---
title: "Merge repositories and keep history"
date: 2021-11-25T15:27:00+01:00
draft: false
tags: ["linux", "git"]
---

There is a special command `git-subtree` to merge / split repositories. It is not installed as a part of `git` package:
```bash
sudo dnf install git-subtree
```

The command to merge repository `jsnjack/X` to subfolder `myfolder/new` at branch `master`:
```bash
git subtree add -P myfolder/new git@github.com:jsnjack/X.git master
```

