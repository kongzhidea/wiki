

```
执行时候指定python路径
#!/usr/bin/python

设置编码
# -*- coding: utf-8 -*-

如果两个都设置，先设置python路径，再设置编码
```

```
查看python版本:
python -V
python --version

python 执行单条命令：
python -c 'command'

查看python依赖的module路径,包含当前目录
import sys
print sys.path

添加本地模块到sys.path
1.sys.path 就是个普通 list，所以 += 或者 append就够用了。
    sys.path.append(myPath)
    sys.path += ['path1', 'path2']
2.设置PYTHONPATH变量，
export PYTHONPATH=$PYTHONPATH:'myPath'

3.pycharm
pycharm中点击File->settings->project XXX->Project Interpreter,点击右上角的设置->more
然后点击最后一个，查看path
点击加号，添加目录


查看依赖包版本
python -c "import pymysql; print pymysql.__version__;"

查看依赖包版本，位置
python -c "import pymysql; print pymysql.__file__;"
```

```
bool类型: True False
逻辑运算: and or not ，没有 && || !
三目运算符  1 if true else 0
if 后面可省略括号
```

```
# 判断类型
isinstance(object, classinfo) 
如:isinstance(2, int)  返回True

# type(object) 返回变量类型，本身返回值为type类型。
如 type(1) 返回type类型<type 'int'>
```

### [WTF Python：有趣且鲜为人知的Python特性](https://github.com/leisurelicht/wtfpython-cn)