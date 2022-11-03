---
title: "git: Cherry-pick branch"
date: 2022-11-03T14:40:44+01:00
draft: false
tags: ["linux", "git"]
---

Very useful command to cherry pick the whole branch
```bash
git cherry-pick 751a77^..8b62f1
```
`^` - means the the first commit will be included too
