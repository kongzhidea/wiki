## git配置
* git config --global user.name "Your Name"设置你的仓库用户名（用于标识提交者）
* git config --global user.email "email@example.com"设置你的仓库邮箱（用于标识提交者）

## git配置文件在：~/.gitconfig

## git命令设置别名
* git config --global alias.co checkout
* git config --global alias.br branch
* git config --global alias.ci commit
* git config --global alias.st status
* git config --global alias.last 'log -1 HEAD'
    * git last

## 查看当前全局配置
* git config --global  --list

## git 安装
* windows
  * 下载Git for windows，插件推荐全部安装，包含 git bash，git history等。
* linux，看虚拟机版本选择下面命令
  *  sudo yum install git-core
  *  sudo yum install git
    * git 颜色高亮：git config --global color.ui true
  *  sudo apt-get install git
* mac
  *  brew install git
  *  [图形化管理sourcetreeapp](https://www.sourcetreeapp.com/)

## 复制远程git仓库   
​$ git clone <版本库的网址>   <br/>
​$ git clone <版本库的网址> <本地目录名>  <br/>
​该命令会在本地主机生成一个目录，与远程主机的版本库同名。如果要指定不同的目录名，可以将目录名作为git clone命令的第二个参数。
​
​
## 添加指定目录到暂存区，包括子目录
$ git add [dir]

## 添加当前目录的所有文件到暂存区
$ git add .

## 添加每个变化前，都会要求确认
## 对于同一个文件的多处变化，可以实现分次提交
$ git add -p

## 删除工作区文件，并且将这次删除放入暂存区
$ git rm [file1] [file2] ...

## 停止追踪指定文件，但该文件会保留在工作区
$ git rm --cached [file]

## 改名文件，并且将这个改名放入暂存区
$ git mv [file-original] [file-renamed]
 
## 提交暂存区到仓库区
$ git commit -m [message]

## 提交暂存区的指定文件到仓库区
$ git commit [file1] [file2] ... -m [message]

## 提交工作区自上次commit之后的变化，直接到仓库区
$ git commit -a  <br/>
$ git commit -A 　　 -A可以包含删除的文件  <br/> 

# 提交时显示所有diff信息
$ git commit -v

## 使用一次新的commit，替代上一次提交，如果代码没有任何新变化，则用来改写上一次commit的提交信息
$ git commit --amend -m [message]

## 重做上一次commit，并包括指定文件的新变化
$ git commit --amend [file1] [file2] ...

 
## 列出所有本地分支
$ git branch

## 列出所有远程分支
$ git branch -r

## 列出所有本地分支和远程分支
$ git branch -a

# 新建一个分支，但依然停留在当前分支
$ git branch [branch-name]

## 新建一个分支，并切换到该分支
$ git checkout -b [branch]

## 新建一个分支，指向指定commit
$ git branch [branch] [commit]

## 新建一个分支，与指定的远程分支建立追踪关系
$ git branch --track [branch] [remote-branch]

## 切换到指定分支，并更新工作区
$ git checkout [branch-name]

## 切换到上一个分支
$ git checkout -

## 建立追踪关系，在现有分支与指定的远程分支之间
$ git branch --set-upstream [branch] [remote-branch]

## 合并指定分支到当前分支
$ git merge [branch]

## 选择一个commit，合并进当前分支
$ git cherry-pick [commit]

## 删除分支
$ git branch -d [branch-name]

## 删除远程分支
$ git push origin --delete [branch-name] <br/>
$ git branch -dr [remote/branch]

 
## 列出所有tag
$ git tag

## 新建一个tag在当前commit
$ git tag [tag]

## 新建一个tag在指定commit
$ git tag [tag] [commit]

## 删除本地tag
$ git tag -d [tag]

# 删除远程tag
$ git push origin :refs/tags/[tagName]

## 查看tag信息
$ git show [tag]

## 提交指定tag
$ git push [remote] [tag]

## 提交所有tag
$ git push [remote] --tags

## 新建一个分支，指向某个tag
$ git checkout -b [branch] [tag]

 
## 显示有变更的文件
$ git status

## 显示当前分支的版本历史
$ git log

## 显示commit历史，以及每次commit发生变更的文件
$ git log --stat

## 搜索提交历史，根据关键词
$ git log -S [keyword]

## 显示某个commit之后的所有变动，每个commit占据一行
$ git log [tag] HEAD --pretty=format:%s

## 显示某个commit之后的所有变动，其"提交说明"必须符合搜索条件
$ git log [tag] HEAD --grep feature

## 显示某个文件的版本历史，包括文件改名
$ git log --follow [file]   <br/>
$ git whatchanged [file]

## 显示指定文件相关的每一次diff
$ git log -p [file]

# 显示过去5次提交
$ git log -5 --pretty --oneline

## 显示所有提交过的用户，按提交次数排序
$ git shortlog -sn

## 显示指定文件是什么人在什么时间修改过
$ git blame [file]

## 显示暂存区和工作区的差异
$ git diff

## 显示暂存区和上一个commit的差异
$ git diff --cached [file]

## 显示工作区与当前分支最新commit之间的差异
$ git diff HEAD

## 显示两次提交之间的差异
$ git diff [first-branch]...[second-branch]

## 显示今天你写了多少行代码
$ git diff --shortstat "@{0 day ago}"

## git diff 忽略 \r,不显示 ^M
$ git config --global core.whitespace cr-at-eol

## 显示某次提交的元数据和内容变化
$ git show [commit]

## 显示某次提交发生变化的文件
$ git show --name-only [commit]

## 显示某次提交时，某个文件的内容
$ git show [commit]:[filename]

## 显示当前分支的最近几次提交
$ git reflog

 
## 下载远程仓库的所有变动
$ git fetch [remote]

## 显示所有远程仓库，查看远程仓库git地址
$ git remote -v

## 显示某个远程仓库的信息
$ git remote show [remote]

## 增加一个新的远程仓库，并命名
$ git remote add [shortname] [url]

## 取回远程仓库的变化，并与本地分支合并
$ git pull [remote] [branch]

## 上传本地指定分支到远程仓库
$ git push [remote] [branch]

## 强行推送当前分支到远程仓库，即使有冲突
$ git push [remote] --force

## 推送所有分支到远程仓库
$ git push [remote] --all

 
## 恢复暂存区的指定文件到工作区
$ git checkout [file]

## 恢复某个commit的指定文件到暂存区和工作区
$ git checkout [commit] [file]

## 恢复暂存区的所有文件到工作区
$ git checkout .

## 重置暂存区的指定文件，与上一次commit保持一致，但工作区不变
$ git reset [file]

## 重置暂存区与工作区，与上一次commit保持一致
$ git reset --hard [commit] <br/>
$ git reset --hard HEAD^ 回到上一个版本<br/>

## 重置当前分支的指针为指定commit，同时重置暂存区，但工作区不变
$ git reset [commit]

```
git revert 撤销 某次操作，此次操作之前和之后的commit和history都会保留，并且把这次撤销
作为一次最新的提交
    * git revert HEAD                  撤销前一次 commit
    * git revert HEAD^               撤销前前一次 commit
    * git revert commit （比如：fa042ce57ebbe5bb9c8db709f719cec2c58ee7ff）撤销指定的版本，撤销也会作为一次提交进行保存。
git revert是提交一个新的版本，将需要revert的版本的内容再反向修改回去，
版本会递增，不影响之前提交的内容

--------------------------------------------------------------------------------------------------------------------------------------------

git revert 和 git reset的区别 
1. git revert是用一次新的commit来回滚之前的commit，git reset是直接删除指定的commit。 
2. 在回滚这一操作上看，效果差不多。但是在日后继续merge以前的老版本时有区别。因为git revert是用一次逆向的commit“中和”之前的提交，因此日后合并老的branch时，导致这部分改变不会再次出现，但是git reset是之间把某些commit在某个branch上删除，因而和老的branch再次merge时，
这些被回滚的commit应该还会被引入。 
3. git reset 是把HEAD向后移动了一下，而git revert是HEAD继续前进，只是新的commit的内容和要revert的内容正好相反，能够抵消要被revert的内容。
```

## 暂时将未提交的变化移除，稍后再移入
* git stash: 备份当前的工作区的内容，从最近的一次提交中读取相关内容，让工作区保证和上次提交的内容一致。同时，将当前的工作区内容保存到Git栈中。
* git stash pop: 从Git栈中读取最近一次保存的内容，恢复工作区的相关内容。由于可能存在多个Stash的内容，所以用栈来管理，pop会从最近的一个stash中读取内容并恢复。
* git stash list: 显示Git栈内的所有备份，可以利用这个列表来决定从那个地方恢复。
* git stash clear: 清空Git栈。此时使用gitg等图形化工具会发现，原来stash的哪些节点都消失了。
* 场景：当前工作区内容已被修改，但是并未完成。这时Boss来了，说前面的分支上面有一个Bug，需要立即修复。可是我又不想提交目前的修改，因为修改没有完成。但是，不提交的话，又没有办法checkout到前面的分支。此时用Git Stash就相当于备份工作区了。然后在Checkout过去修改，就能够达到保存当前工作区，并及时恢复的作用。

## [Removing sensitive data from a repository](https://help.github.com/articles/removing-sensitive-data-from-a-repository/)
