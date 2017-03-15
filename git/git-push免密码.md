```
1.1 创建文件存储GIT用户名和密码

在%HOME%目录中，一般为C:\users\Administrator，也可以是你自己创建的系统用户名目录，
反正都在C:\users\中。文件名为.git-credentials,由于在Window中不允许直接创建以"."开头的文件，
所以需要借助git bash进行，打开git bash客户端，进行%HOME%目录，
然后用touch创建文件 .git-credentials, 用vim编辑此文件，输入内容格式：

touch .git-credentials

vim .git-credentials

https://{username}:{password}@github.com

1.2 添加Git Config 内容

进入git bash终端， 输入如下命令：

git config --global credential.helper store

执行完后查看%HOME%目录下的.gitconfig文件，会多了一项：

[credential]

    helper = store
重新开启git bash会发现git push时不用再输入用户名和密码
```