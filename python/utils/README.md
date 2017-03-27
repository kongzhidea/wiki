

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

查看依赖包版本
python -c "import pymysql; print pymysql.__version__;"

查看依赖包版本，位置
python -c "import pymysql; print pymysql.__file__;"
```

```
bool类型: True False
逻辑运算: and or not ，没有 && || !
三目运算符  1 if true else 0
```