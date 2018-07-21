# If you come from bash you might have to change your $PATH.
# export PATH=$HOME/bin:/usr/local/bin:$PATH

# Path to your oh-my-zsh installation.
  export ZSH=/home/cb/.oh-my-zsh

# Set name of the theme to load. Optionally, if you set this to "random"
# it'll load a random theme each time that oh-my-zsh is loaded.
# See https://github.com/robbyrussell/oh-my-zsh/wiki/Themes
ZSH_THEME="steeef"

# Uncomment the following line to use case-sensitive completion.
# CASE_SENSITIVE="true"

# Uncomment the following line to use hyphen-insensitive completion. Case
# sensitive completion must be off. _ and - will be interchangeable.
# HYPHEN_INSENSITIVE="true"

# Uncomment the following line to disable bi-weekly auto-update checks.
# DISABLE_AUTO_UPDATE="true"

# Uncomment the following line to change how often to auto-update (in days).
# export UPDATE_ZSH_DAYS=13

# Uncomment the following line to disable colors in ls.
# DISABLE_LS_COLORS="true"

# Uncomment the following line to disable auto-setting terminal title.
# DISABLE_AUTO_TITLE="true"

# Uncomment the following line to enable command auto-correction.
# ENABLE_CORRECTION="true"

# Uncomment the following line to display red dots whilst waiting for completion.
# COMPLETION_WAITING_DOTS="true"

# Uncomment the following line if you want to disable marking untracked files
# under VCS as dirty. This makes repository status check for large repositories
# much, much faster.
# DISABLE_UNTRACKED_FILES_DIRTY="true"

# Uncomment the following line if you want to change the command execution time
# stamp shown in the history command output.
# The optional three formats: "mm/dd/yyyy"|"dd.mm.yyyy"|"yyyy-mm-dd"
# HIST_STAMPS="mm/dd/yyyy"

# Would you like to use another custom folder than $ZSH/custom?
# ZSH_CUSTOM=/path/to/new-custom-folder

# Which plugins would you like to load? (plugins can be found in ~/.oh-my-zsh/plugins/*)
# Custom plugins may be added to ~/.oh-my-zsh/custom/plugins/
# Example format: plugins=(rails git textmate ruby lighthouse)
# Add wisely, as too many plugins slow down shell startup.
plugins=(git)

source $ZSH/oh-my-zsh.sh

# User configuration

# export MANPATH="/usr/local/man:$MANPATH"

# You may need to manually set your language environment
# export LANG=en_US.UTF-8

# Preferred editor for local and remote sessions
# if [[ -n $SSH_CONNECTION ]]; then
#   export EDITOR='vim'
# else
#   export EDITOR='mvim'
# fi

# Compilation flags
# export ARCHFLAGS="-arch x86_64"

# ssh
# export SSH_KEY_PATH="~/.ssh/rsa_id"

# Set personal aliases, overriding those provided by oh-my-zsh libs,
# plugins, and themes. Aliases can be placed here, though oh-my-zsh
# users are encouraged to define aliases within the ZSH_CUSTOM folder.
# For a full list of active aliases, run `alias`.
#
# Example aliases
# alias zshconfig="mate ~/.zshrc"
# alias ohmyzsh="mate ~/.oh-my-zsh"

# golang 环境变量
export GOPATH=/home/cb/gocode
export GOROOT=/home/cb/go
export GOARCH=amd64
export GOOS=linux
export GOTOOLS=$GOROOT/pkg/tool/
export PATH=$PATH:$GOROOT/bin:$GOTOOLS
export GOBIN=/home/cb/gocode/bin


# Add RVM to PATH for scripting. Make sure this is the last PATH variable change.
export PATH="$PATH:$HOME/.rvm/bin"
source ~/.rvm/scripts/rvm

# 别名
alias dp="docker ps -a"
alias di="docker images"
alias cdke="cd ~/ke/kings-server"
alias cdyf="cd ~/gocode/src/yf/platform"
alias cdt1="cd ~/gocode/src/caibo/test/t1"
alias cdgo="cd ~/gocode/src"
alias cds="cd /home/cb/gocode/src/yf/server/gs"
alias cdw="cd /home/cb/gocode/src/yf/server/gw"
alias cdc="cd /home/cb/gocode/src/yf/server/gc"
alias cdt="cd /home/cb/gocode/src/yf/platform/yftools/gs2go"
alias bdh1="ruby /home/cb/ke/ke-tools/locale_tool/csv2yamls.rb /home/cb/下载/123.csv /tmp/locales"
alias bdh2="ruby /home/cb/ke/ke-tools/locale_tool/append_yamls.rb /tmp/locales /home/cb/ke/kings-server/config/locales"
alias bdh2h="ruby /home/cb/ke/ke-tools/locale_tool/append_yamls.rb /tmp/locales /home/cb/ke/kings-server/config/locales/hero"
alias bdh2a="ruby /home/cb/ke/ke-tools/locale_tool/append_yamls.rb /tmp/locales /home/cb/ke/kings-server/config/locales/activities"
alias kebs="/home/cb/ke/ke-tools/merge_branch && /home/cb/ke/ke-tools/cmd_id cn mail"
alias gg="git commit -a -m 'ke' && gp && ./stagingdeployw"
alias gg2="git commit -a -m 'ke' && gp && ./stagingdeployw two"
alias t1="ssh ubuntu@121.43.185.222"
alias t2="ssh ubuntu@121.43.185.243"
