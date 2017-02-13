# anyframe fzf

## How to setup

```zsh
fpath=(${0:h}/anyframe(N-/) $fpath)
source "${0:h}/anyframe-fzf/anyframe-fzf.init.zsh" # insert before autoload anyframe-init
autoload -Uz anyframe-init
anyframe-init
```

## Widgets

- anyframe-widget-checkout-git-branch-all
- anyframe-widget-git-log
- anyframe-widget-git-status
- anyframe-widget-pickup-github-features
