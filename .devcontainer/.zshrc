export ZSH="$HOME/.oh-my-zsh"
ZSH_THEME="robbyrussell"
# Switch to Powerlevel10k by uncommenting the next line:
# ZSH_THEME="powerlevel10k/powerlevel10k"

plugins=(
  git
  sudo
  catimg
  docker
  git-extras
  github
  golang
  gpg-agent
  node
  nvm
  vscode
)

# NVM (installed below) init
export NVM_DIR="$HOME/.nvm"
[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"  # loads nvm
[ -s "$NVM_DIR/bash_completion" ] && \. "$NVM_DIR/bash_completion"

# Quality of life
export EDITOR=vim
setopt HIST_IGNORE_ALL_DUPS SHARE_HISTORY
HISTSIZE=50000
SAVEHIST=50000

# Fix fd name on Debian/Ubuntu (fd-find installs as fdfind)
if command -v fdfind >/dev/null 2>&1 && ! command -v fd >/dev/null 2>&1; then
  alias fd=fdfind
fi

# If running inside VS Code devcontainer, prefer login shell behavior
[[ -r /etc/profile ]] && source /etc/profile