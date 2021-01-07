### 在mac终端上 git没有自动补全功能，需要手动加上此功能


## bash环境：
#### 1. 下载Git-completion.bash
* github地址：https://github.com/markgandolfo/git-bash-completion.git

#### 2. copy到用户根目录~/
* cp git-completion.bash ~/.git-completion.bash

#### 3. 使之生效
* source ~/.git-completion.bash

#### 4. 修改~/.profile
* 增加 source ~/.git-completion.bash


## zsh环境：
* https://blog.csdn.net/puliscode/article/details/102495868
```
macOS zsh 下git的自动完成功能
更新catalina后mac的默认shell提示变更为zsh，zsh理论上完美兼容bashrc，但是git 的自动提示功能不能用了，需要重新设置，好在有人已经写好了脚本，只需要下载配置就好了：

安装hub
brew install hub

mkdir ~/.zsh
mkdir ~/.zsh/completions

复制脚本到用户目录
cp /usr/local/etc/bash_completion.d/hub.bash_completion.sh ~/.zsh/completions/_hub

编辑.zshrc文件
vim .zshrc

fpath=(~/.zsh/completions $fpath)
autoload -U compinit && compinit

执行配置文件
source .zshrc
```
