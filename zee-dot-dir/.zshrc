# fix backspace, etc
export TERM=xterm

# some bling
PS1="🐷 > "

# NVM (TODO: only if nvm in Brewfile)
if [[ "$HOMEBREW_DEPENDENCIES" == *"nvm"* ]]; then
  export NVM_DIR="$HOME/.nvm"
  . "$HOMEBREW_PREFIX/opt/nvm/nvm.sh"
  . "$HOMEBREW_PREFIX/opt/nvm/etc/bash_completion.d/nvm"

  # and automate `nvm use` on directory change
  # https://github.com/nvm-sh/nvm#zsh
  autoload -U add-zsh-hook
  load-nvmrc() {
    local node_version="$(nvm version)"
    local nvmrc_path="$(nvm_find_nvmrc)"

    if [ -n "$nvmrc_path" ]; then
      local nvmrc_node_version=$(nvm version "$(cat "${nvmrc_path}")")

      if [ "$nvmrc_node_version" = "N/A" ]; then
        nvm install
      elif [ "$nvmrc_node_version" != "$node_version" ]; then
        nvm use
      fi
    elif [ "$node_version" != "$(nvm version default)" ]; then
      echo "Reverting to nvm default version"
      nvm use default
    fi
  }
  add-zsh-hook chpwd load-nvmrc
  load-nvmrc
fi

# rbenv
# TODO: improve impl
if [[ "$HOMEBREW_DEPENDENCIES" == *"rbenv"* ]]; then
  eval "$($HOMEBREW_PREFIX/bin/rbenv init - zsh)"
fi

# goenv
# TODO: improve impl
if [[ "$HOMEBREW_DEPENDENCIES" == *"rbenv"* ]]; then
  eval "$($HOMEBREW_PREFIX/bin/goenv init -)"
fi
