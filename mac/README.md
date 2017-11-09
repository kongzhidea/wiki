

## 1. bash_profile 
* 见.bash_profile

## 2.安装常见软件--brew
```
安装好了后，执行  sudo brew update
将 /usr/local 设置权限到当前用户

cd /usr/local
sudo chown -R kongzhihui:staff .

安装软件，不要用root用户：
brew install wget

卸载软件：
brew uninstall wget
```

## 3.安装常见软件
* 视频播放  MplayerX
* [macdown](http://www.pc6.com/mac/137960.html)
* item2    mac终端
* ihosts 绑定环境。
* 连接数据库软件--sequel pro，下载对应的dmg文件即可。

## 4.显示/隐藏  隐藏文件夹
* 显示：
	* defaults write com.apple.finder AppleShowAllFiles -boolean true ; killall Finder
* 隐藏：
	* defaults write com.apple.finder AppleShowAllFiles -boolean false ; killall Finder

### 4.1 文本编辑
* 点击文本编辑，弹出窗口而不是新建文件：
    * 设置-->iCloud -->取消 iCloud Drive勾选
* OS X 文本编辑等应用为何会自动转换引号
    * 系统偏好设置-->键盘-->文本-->取消 使用智能引号和破折号  
    * 文本编辑-->偏好设置-->取消 智能引号，智能破折号，智能链接
    
## 5.在finder中打开终端
```
系统偏好设置 -> 键盘 -> 快捷键 -> 服务，勾选「新建位于文件夹位置的终端窗口」
（后面的键盘快捷键可以不选），然后在 Finder 里面选中文件夹右键菜单的「服务」，
下面就会有「新建位于文件夹位置的终端窗口」这一子菜单了。

缺点：只能对选中的文件夹操作。



在终端中打开 finder， 使用open 命令，后面接 打开的目录参数。
```

## 6.安装nginx
```
brew install nginx

nginx.conf在：/usr/local/etc/nginx

默认html 在 ：/usr/local/var/www
启动nginx ：sudo nginx -c /usr/local/etc/nginx/nginx.conf
```

## 7.安装java
```
下载：jdk-8u111-macosx-x64_8.0.1110.14.dmg
默认java_home : /Library/Java/JavaVirtualMachines/jdk1.8.0_111.jdk/Contents/Home

export JAVA_HOME="/Library/Java/JavaVirtualMachines/jdk1.8.0_111.jdk/Contents/Home"
```

## 8.安装idea，eclipse，pycharm，mysql 下载对应的dmg文件即可。
* [mysql安装](https://segmentfault.com/a/1190000004061246)
* mysql启动时候，保证/usr/local/mysql是mysql用户，执行命令：sudo chown -R mysql:mysql /usr/local/mysql-5.7.16-osx10.11-x86_64/

## 9.tree
* brew install tree

## mdfind
* mdfind test
* 类似linux的locate功能

## 切换root用户
* sudo -i

## 查找进程
* ps -ef |grep pid

## say
* say "hello"
* 文本转语音，朗读给定字符串

## java InetAddress.getLocalHost() 方法特别慢
```
处理方式：
hostname   获取主机名：kongzhihuideMacBook-Pro.local

/etc/hosts中增加
127.0.0.1		   kongzhihuideMacBook-Pro.local
::1             kongzhihuideMacBook-Pro.local
```

## 回车 换行符
```
CRLF:
	CR   \r
	LF   \n
	CRLF  \r\n

windows 回车为：CRLF：\r\n
Linux和MAC 为LF：\n


这样在vim中打开windows中记事本文件时候，出现的^M即为\r符：
VIM中删除^M:
	“1,$s/\r//g”
	“1,$s/^M//g”
Linix下打出 ^M：ctrl+V，ctrl+M

git diff时候忽略 \r：
	git config --global core.whitespace cr-at-eol

```

## vim配置
```
见 .vimrc
syntax on   彩色vim
set ignorecase 忽略大小写
set expandtab tab转换为空格
set noexpandtab tab不转换为空格
set tabstop=4 tab转换空格数量
```

## 快捷键
```
Dock栏中保留应用--邮件-选项-在Dock中保留
双显示器--鼠标移动到最底部，然后往下滑动，可以将Dock栏切换至此屏幕，恢复时候操作一样，也可以关闭显示器，合盖后自动恢复。

设置--触摸板，设置轻点来点按：系统便好设置－－触摸板－－轻点来点按
设置--键盘--文本， 取消 自动大些单词首字母，连续两个空格插入句号。

chrome   调试工具:  command+alt+i
chrome  刷新页面 :  command+r
chrome 下载内容： command+shift+j

关闭窗口   command+w

终端 打开新标签: command+t
从桌面打开终端：command+t
终端 窗口切换：  command+{， 按的时候：command+shift+[
清除屏幕内容：  ctrl+l，或 Command + K 
保存终端输出    Command + S 

终端打开 finder：  open 命令，后面加路径，例如： open .
finder打开终端：https://my.oschina.net/ioslighter/blog/359258?p={{totalPage}}

F3（不按fn） 打开当前应用选择界面，对应的手势   三个手指上滑，快捷键：ctrl+up。
        打开当前应用选择界面，但是只展示相同应用窗口，对应手势  三个手指下滑，快捷键：ctrl+up，然后可以三个手指向左，右滑动。
F4（不按fn） 打开全部应用选择界面，对应的手势   5个手指向中间滑。
回到桌面手势： 5个手指从中间向四周划。
command+h   隐藏当前界面。


隐藏底部菜单栏：alt+command+d

截屏快捷键：ctrl+command+shift+4 自定义截屏，ctrl+command+shift+4 截全屏，推荐使用钉钉的截屏软件。

右键： 两根手指 一起点击，或者ctrl+单击

光标移动： 移动到最左侧：ctrl+a， 移动到最右侧：ctrl+e
清除当前输入命令  ctrl+u， 
取消当前命令  ctrl+c
删除当前输入命令最后一个单词 ctrl + w

finder中  跳转到上一级：command+ 上

command + C, command + option + V	剪切文件或者文件夹
fn + delete	向后删除（适用所有编辑情况）
command + option + shift + V	实现去除格式的粘贴文本
command + option + shift + esc	强制退出当前窗口（当应用停止响应的时候，非常好用）

文本处理相关的快捷键操作


快捷键	对应功能
command + 向左（右）箭头	将光标快速移动到行首（末）
fn + 向上（下）箭头	向上（下）滚动一页（Page Up, Page Down）
fn + 向左（右）箭头	快速移动到文稿的开头（末尾）（Home，End）
option + 向左（右）箭头	将光标移动到当前单词的开头（末尾）
finder中的一些常用快捷键
快捷键	对应功能
command + shift + G	快速前往指定路径的文件夹
command + 1, 2, 3, 4	切换图标，列表，分栏，cover flow视图
向上和向下的箭头	选择指定文件
向左和向右的箭头	上级和下级目录
command + 向上箭头	上层文件夹
command + 【	向前
command + 】	向后
enter	重命名
space	快速预览
command + W	停止快速预览
command + I	快速查看文件或者文件夹的详细信息（大小，时间等）
command + C, command + option + V	复制文件或者文件夹到指定路径
command + shift + N	当前目录下新建一个文件夹
command + option + C	复制当前选中文件或者文件夹的路径
command + option + P 显示（隐藏）路径，这个属于一次性设置操作，显示路径后会给你更加清晰的导航效果。



终端快捷键

终端通用的快捷键
快捷键	对应功能
control + A	跳到行首
control + E	跳到行末
control + U	清除光标到行首的内容
control + K	清除光标到行末的内容
control + W	删除光标所在位置的单词
control + L	清除屏幕内容
tab	补全或者提示内容
iTerm作为常用的终端替代者，我们也介绍一下它的快捷键
快捷键	对应功能
command + T	新建标签页（很多通用）
command + W	关闭标签页（很多通用）
command + 左右箭头/ command + 数字	切换标签
command + enter	切换全屏
command + D	垂直分屏
command + shift + D	水平分屏
command + option + 箭头	切换分屏
command + control + 箭头	调节分屏的大小
command + shift + h	查看剪切板历史
control + r	搜索历史命令（这个非常常用）
关于chrome浏览器的快捷键

标签页和窗口快捷键
快捷键	对应功能
command + N	打开新窗口
command + T	打开新标签页
command + shift + N	隐身模式打开新窗口
command + O	然后选择文件，使用chrome打开指定文件
command + W	关闭当前标签页
command + shift + W	关闭当前窗口
command + shift + T	恢复误关闭的标签
command + R	刷新当前页面
command + option + 向左（右）的箭头	向左（右）切换标签页
功能快捷键
快捷键	对应功能
command + shift + B	打开（关闭）书签栏
command + Y	打开历史记录
command + shift + delete	打开清除缓存页
command + shift + J	打开下载内容页
地址栏快捷键
快捷键	对应功能
command + L	快速锁定地址栏
输入网址 + command + enter	通过新的标签页打开键入的网址
网页快捷键
快捷键	对应功能
command + option + I	打开开发者工具
command + option + J	打开开发者工具中的控制台
command + option + U	通过新标签页打开当前页面的源代码
command + D	将当前页面添加收藏
space	向下滚动页面
command + 向上的箭头	快速滚到页面顶部
command + 向下的箭头	快速滚到页面底部
fn + 向上的箭头	向上翻页
fn + 向下的箭头	向下翻页


idea中  调出设置页面：command+,
    设置快捷键为eclipse模式，
    设置 跟进类文件， 快捷键（keymap）  declaration，添加鼠标快捷键，设置为command+单击
    查看方法被调用情况：快捷键（keymap）－－method hierarchy，添加一个快捷键:command+shift+h
    设置 copy,cut,past,undo,redo,save快捷键 ，加一个  command+c等。
    设置 replace快捷键  加一个  command+f
    生成getter和setter，找到keymap设置，然后搜索generate，设置快捷键为：command+i
    run 执行main方法，增加快捷键 command+r
    注释快捷键：keymap--commont--，设置快捷键未  command+/
    eclipse快捷键：
    自动import类，快捷键：alt＋enter
    ctrl+o，搜索类中方法
    ctrl+shift+t，搜索类
    ctrl+shift+r，搜索文件
    ctrl+shift+f，格式化代码
    ctrl+e 查询最近文件
    psvm，快捷生成main方法，
    sout，快捷生成System.out.println代码
    shift+enter，快速切换到下一行，并新启一行
    ctrl+alt+下， 复制一行到下一行，并新启一行
    ctrl+d, 删除一行
    复制  ctrl+c，剪切  ctrl+v，撤销  ctrl+z
    ctrl+/ ，注释一行
```


	
