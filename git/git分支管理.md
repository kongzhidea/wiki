### 1.创建本地分支
> git branch 分支名，例如：git branch 2.0.1.20120806    
> 注：2.0.1.20120806是分支名称，可以随便定义。 
    
### 2.切换本地分支
> git checkout 分支名，例如从master切换到分支：git checkout 2.0.1.20120806    
> git checkout -b branchName  创建分支，并且切换到分支。

### 3.远程分支就是本地分支push到服务器上。比如master就是一个最典型的远程分支（默认）。
> git push origin 2.0.1.20120806    

### 4.远程分支和本地分支需要区分好，所以，在从服务器上拉取特定分支的时候，需要指定远程分支的名字。
> git checkout --track origin/2.0.1.20120806    
> 注意该命令由于带有--track参数，所以要求git1.6.4以上！这样git会自动切换到分支。

### 5.提交分支数据到远程服务器
> git push origin \<local_branch_name>:<remote_branch_name\>
> 例如：    
> git push origin 2.0.1.20120806:2.0.1.20120806    
> 一般当前如果不在该分支时，使用这种方式提交。如果当前在 2.0.1.20120806 分支下，也可以直接提交：git push

### 6.删除远程分支,origin后面加一个空格 
> git push origin :develop

### 7.删除本地分支：
> git branch -d xxxxx


### 8.本地分支与远程分支做关联：
> git branch --set-upstream-to=origin/master  master   
> 一般直接从远程checkout下来的分支，会自动关联   


### 9.查看远程分支
> git branch -a

### 10.查看本地分支
> git branch



### 11.Git 获取远程分支
> 如果安装了git客户端，直接选择fetch一下，就可以获取到了。    
> 如果用命令行，运行 git fetch，可以将远程分支信息获取到本地，再运行 git checkout -b local-branchname origin/remote_branchname  就可以将远程分支映射到本地命名为local-branchname  的一分支。 




