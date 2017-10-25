
### 常见命令

#### 快捷键
```
1. ctrl+c 取消本行输入
2. ctrl+r 查询历史输入
3. ctrl+d 终止输入
4. cd - （cd空格 减号）返回最近一次访问的目录
```

#### vim配置文件见  .vimrc
* set nu  显示行号
* set nonu 不显示行号
* set ignorecase 忽略大小写
* set noignorecase 不忽略大小写
* set fileencoding 显示当前文件编码
* set fileencoding=utf-8  设置当前文件编码为 utf-8
* set expandtab 设置tab自动转空格
* set noexpandtab 取消设置tab自动转空格
* set tabstop=4

#### ssh登录免密码
```
ssh 无密码登录要使用公钥与私钥。linux下可以用用ssh-keygen生成公钥/私钥对。
有机器A(192.168.1.155)，B(192.168.1.181)。现想A通过ssh免密码登录到B。

1.在A机下生成公钥/私钥对。
[root@A ~]# ssh-keygen -t rsa -P ''


2.把A机下的/root/.ssh/id_rsa.pub 复制到B机的 /root/.ssh/authorized_keys文件里，先要在B机上创建好 /root/.ssh 这个目录，或者：用scp复制或者：cat >> .ssh/authorized_keys

3.authorized_keys的权限要是600!!!
```

#### find
```
1. 命令格式：find [path]  -name [fileName]，  xargs为找到文件后执行命令
2. find . -name *.conf     寻找当前目录下所有.conf文件，包含子目录。
3. find . -name *.conf  |xargs file， 找到文件后打印文件类型
```

#### locate 快速定位某个文件地址
```
centos 安装 locate
    yum -y install mlocate
    updatedb
    
locate fileName   在整个系统中查询该文件，可以使用通配符。

-e 将排除在寻找的范围之外。
-1 如果 是 1．则启动安全模式。在安全模式下，使用者不会看到权限无法看到的档案。
	这会始速度减慢，因为 locate 必须至实际的档案系统中取得档案的权限资料。
-f 将特定的档案系统排除在外，例如我们没有到理要把 proc 档案系统中的档案放在资料库中。
-q 安静模式，不会显示任何错误讯息。
-n 至多显示 n个输出。
-r 使用正规运算式 做寻找的条件。
-o 指定资料库存的名称。
-d 指定资料库的路径
-h 显示辅助讯息
-V 显示程式的版本讯息
```

#### grep 搜索文件内容
```
grep参数：
    -i 忽略大小写
    -r 递归查询
    -n 打印结果行标
    -v 打印不包含的行
    -m 10  只找前10个匹配结果
    --color=auto 高亮显示

grep 后面的文件或者目录，可以使用通配符，如 *.log
    
样例如下：    
grep -r "search content" .   在当前目录搜索内容
```

#### ln 软连接
```
1. 创建软连接:
	ln -s /data/logs logs
2. 查看软连接:
    ll 
```

#### tar 压缩，解压
```
打包:
tar -cvf target.tar target/

解压：
tar -xvf target.tar   解压到当前目录

打包压缩：
tar -zcvf target.tar.gz target/

解压压缩包：
tar -zxvf target.tar.gz   解压到当前目录


-cvf 也可以写成  cvf
-zxvf 也可以写成  zxvf
```

#### unzip 解压zip文件
* zip -r target.zip target/   -r 递归压缩
* unzip target.zip   解压到当前目录

#### linux查看cpu核数
```
1. top ， 按1查看
2. cat /proc/cpuinfo |grep "processor"|wc -l  
3. cat /proc/cpuinfo |grep "cores"|uniq  
```

#### top，htop 查看当前系统进程使用cpu，内存等使用程度。

#### 修改权限
```
授权操作到某用户：
1. chown  -R kongzhidea:kongzhidea .    将当前目录设置为 kongzhidea:kongzhidea用户,
2. chmod -R 775 .  设置权限。

3. MAC下面：sudo chown -R kongzhihui:staff /data/

授权操作到所有用户：
1. sudo chmod -R 777 /data/
```

#### 设置别名
```
alias 或者  alias -p  列出所有设置的别名。

设置别名，可以放到 ~/.bash_rc文件中：

alias grep='grep --color=auto'
alias l='ls -l'
alias la='ls -a'
alias ll='ls -la'
```


#### crontab 定时任务
```
30 6 * * *  cd /data && bash run.sh   #  凌晨6点执行，run.sh中需要有  source /home/zhihui.kzh/.bash_rc， 命令最好用绝对地址。

0 */1 * * *  每隔一小时执行
*/5 * * * *  每隔5分钟执行
0 9 * * 5    周五上午9点执行
```

#### data 日期操作
```
date +%Y-%m-%d     # 今天日期，精确到天，结果如：2017-09-01
date +"%Y-%m-%d" -d "-1day"   #昨天日期，精确到天
date +"%Y-%m-%d %H:%M:%S"   # 2014-04-23 12:14:35
```

#### 远程上传/下载文件
```
linux下rzsz的安装：sudo apt-get install lrzsz

rz -be 是使用二进制来上传。
rz 使用普通方式来上传

sz 下载文件
```

#### wget，下载网络文件
```
wget http://download.url

wget http://download.url  --no-check-certificate
```

#### shell 获取字符串长度
```
echo "abcd" |awk '{print length($0)}'
```

#### watch 监控某命令运行结果
```
watch可以帮你监测一个命令的运行结果，watch是周期性的执行下个程序，并全屏显示执行结果。
	例如 监测某个文件的大小变化！

1．命令格式：
watch[参数][命令]

2．命令功能：
可以将命令的输出结果输出到标准输出设备，多用于周期性执行命令/定时执行命令

3．命令参数：
-n或--interval  watch缺省每2秒运行一下程序，可以用-n或-interval来指定间隔的时间。
-d或--differences  用-d或--differences 选项watch 会高亮显示变化的区域。
	 而-d=cumulative选项会把变动过的地方(不管最近的那次有没有变动)都高亮显示出来。
-t 或-no-title  会关闭watch命令在顶部的时间间隔,命令，当前时间的输出。
-h, --help 查看帮助文档

4．使用实例：
实例1：
	命令：每隔一秒高亮显示网络链接数的变化情况
	watch -n 1 -d netstat -ant

实例2：
	watch 'du ~/logs/kk.log'
	
ctrl+c 退出watch命令。
```

#### wc 统计字符 
```
参数如下：
	-c 统计字节数。
	-l 统计行数。
	-m 统计字符数。这个标志不能与 -c 标志一起使用。
	-w 统计字数。一个字被定义为由空白、跳格或换行字符分隔的字符串。
	-L 打印最长行的长度。
	-help 显示帮助信息
	--version 显示版本信息

实例1：查看文件的字节数、字数、行数
	wc test.txt
	
实例2：用wc命令怎么做到只打印统计数字不打印文件名
	[root@localhost test]#wc -l test.txt
		7 test.txt
	[root@localhost test]#cat test.txt |wc -l
		7

实例3：用来统计当前目录下的文件数，包含目录
	ls -l | wc -l		
```

#### nl 显示行号
```
nl file
```

#### free 查询内存
```
free -m   输出单位为M
free -g   输出单位为G
```

#### du查询磁盘
```
du -sh .  查询当前目录的大小
du -sh /data/* 查询data目录下所有文件大小
```

#### nohup 后台执行某命令，默认输出到 nohup.out，可以指定输出日志文件。

#### ps 显示进程
```
查询所有进程
ps -ef 
ps -axu  
    有的机器上可以用  pa -axu，有的机器上使用  ps -ef。

查询某个进程
ps -axu   |grep java


%CPU %MEM  分别代表cpu和内存使用情况。

找出当前系统内存使用量较高的进程：ps -aux | sort -rnk 4 | head -20
找出当前系统CPU使用量较高的进程：ps -aux | sort -rnk 3 | head -20
```

#### 持续ping并将结果记录到日志
```

ping api.jpush.cn | awk '{ print $0"    " strftime("%Y-%m-%d %H:%M:%S",systime()) } '  >> /tmp/push.log &
```
