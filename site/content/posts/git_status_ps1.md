---
title: "PS1 with git status without any extra dependencies"
date: 2024-11-03T15:59:51+01:00
draft: false
tags: ["linux", "bash"]
---

In this guide, we will add git status to the command line prompt without any extra dependencies.

Our new prompt has the following additional functionality:
 - If the command exits with a non-zero status, it will show the exit code in red, e.g., `✘127`.
 - The username and hostname are shown in green, e.g., `user@hostname` (should prevent the user from executing commands on the wrong host).
 - The current working directory is shown in blue.
 - The git status is shown in square brackets (only if inside a git repository):
    - The branch name is shown in yellow.
    - Untracked files are shown with a yellow question mark.
    - Staged files are shown with a green plus sign.
    - Unstaged files are shown with a red exclamation mark.
    - Ahead and behind of upstream are shown with up and down arrows, respectively.
 - The prompt is split into multiple lines for better readability.

Open your `~/.bashrc` file and add the following snippet:

```bash
function git_status_prompt() {
    local git_info=""
    if git rev-parse --is-inside-work-tree &>/dev/null; then
        # Define colors
        local branch_color="\033[0;33m"     # Yellow for branch name
        local reset_color="\033[0m"         # Reset to default color
        local untracked_color="\033[1;33m"  # Bold yellow for untracked
        local staged_color="\033[1;32m"     # Bold green for staged
        local unstaged_color="\033[1;31m"   # Bold red for unstaged

        # Show branch name if on a branch, otherwise show commit hash
        local branch_name=$(git symbolic-ref --short -q HEAD 2>/dev/null)
        if [[ -n "$branch_name" ]]; then
            git_info+="[${branch_color}${branch_name}${reset_color}"
        else
            git_info+="[${branch_color}@$(git rev-parse --short HEAD)${reset_color}"
        fi

        # Add additional status indicators with respective colors
        [[ -n $(git ls-files --others --exclude-standard) ]] && git_info+=" ${untracked_color}?${reset_color}"
        ! git diff --cached --quiet && git_info+=" ${staged_color}+${reset_color}"
        ! git diff --quiet && git_info+=" ${unstaged_color}!${reset_color}"

        # Ahead/behind of upstream
        local ahead=$(git rev-list --count @{u}..HEAD 2>/dev/null)
        local behind=$(git rev-list --count HEAD..@{u} 2>/dev/null)
        [[ $ahead -gt 0 ]] && git_info+=" ⇡$ahead"
        [[ $behind -gt 0 ]] && git_info+=" ⇣$behind"

        # Close the square bracket
        git_info+="]"
    fi
    echo -e "$git_info"
}
PS1='$(if [ $? -ne 0 ]; then echo "\[\e[0;31m\]✘$? \[\e[0m\] "; fi)\[\e[01;32m\]\u@\h\[\e[0m\] \[\e[01;34m\]\w\[\e[0m\] $(git_status_prompt)\n\$ '
```

Apply the changes by running:
```bash
source ~/.bashrc
```
