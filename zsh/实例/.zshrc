# nvm配置
export NVM_DIR="$HOME/.nvm"
nvm_sh="/opt/homebrew/opt/nvm/nvm.sh"
[ -s "$nvm_sh" ] && source "$nvm_sh"
nvm_completion="/opt/homebrew/opt/nvm/etc/bash_completion.d/nvm"
[ -s "$nvm_completion" ] && source "$nvm_completion"

# 别名
alias ll="ls -l -a"
alias rc="source ~/.zshrc"
alias mp="cd ~/my-projects"
alias yj="cd ~/yh-projects/ycjob"
alias dotnet9="/opt/homebrew/Cellar/dotnet@9/9.0.112/libexec/dotnet"

# cd后执行
chpwd() {
    local need16_prefix='/Users/au/yh-projects'
    local node_version_in_tree=16
    local node_version_default=24
    local in_now=$PWD
    local in_old=$OLDPWD

    # 当前是否位于 need-node-16 目录树内
    local now_in_tree=false old_in_tree=false
    [[ $in_now == $need16_prefix || $in_now == $need16_prefix/* ]] && now_in_tree=true
    [[ $in_old == $need16_prefix || $in_old == $need16_prefix/* ]] && old_in_tree=true

    if [[ $1 == 'is_init' ]]; then
        if [[ $now_in_tree == true ]]; then
            nvm use $node_version_in_tree
        else
            if [[ $(node -v) != v${node_version_default}* ]]; then
                nvm use $node_version_default
            fi
        fi
    else
        case "$old_in_tree,$now_in_tree" in
            false,true)  nvm use $node_version_in_tree ;;   # 进入树
            true,false)  nvm use $node_version_default ;;   # 离开树
        esac
    fi
    
    # my项目中检测git用户名是否正确
    my_name=Au
    my_email=18971501210@189.cn
    if [[ $PWD == /Users/au/my-projects* && -d .git ]]; then
        local name=$(git config --get user.name 2>/dev/null)
        if [[ $name != $my_name ]]; then
            print -P "%F{red}⚠️  Git user.name 不是 $my_name（当前：${name:-空}）已经重设%f"
            git config user.name "$my_name"
        fi
        local email=$(git config --get user.email 2>/dev/null)
        if [[ $email != $my_email ]]; then
            print -P "%F{red}⚠️  Git user.email 不是 $my_email（当前：${email:-空}）已经重设%f"
            git config user.email "$my_email"
        fi
    fi
}
# zshrc时触发一次
chpwd 'is_init'

# 末尾是脚本管理的开关
