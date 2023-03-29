---
title: "git: How to remove submodule pointer update in pull request"
date: 2023-03-29T11:13:27+02:00
draft: false
tags: ["linux", "git"]
---

Case 1: it is done in a separate commit:
 - rebase on a base branch `git rebase -i <base_branch>`
 - in the editor, find the commit which updates the pointer and remove it from history (by replacing `pick` with `d`)
 - force push it when the rebase is done
 
Case 2: it is done in the commit which contains other changes
 - rebase on a base branch `git rebase -i <base_branch>`
 - in the editor, find the commit which includes the pointer update and mark it for editing (by replacing `pick` with `e`)
 - git will start rebasing the code and will stop on selected commit for editing
 - unstage the change `git restore --staged <path_to_submodule>`
 - continue the rebase `git rebase --continue`
 - force push it when the rebase is done
