# -*- coding: utf-8 -*-


'''
try:
<语句>        #运行别的代码
except <名字>：
<语句>        #如果在try部份引发了'name'异常
except <名字>，<数据>:
<语句>        #如果引发了'name'异常，获得附加的数据
else:
<语句>        #如果没有异常发生
'''

'''
常见异常
BaseException	所有异常的基类
Exception	常规错误的基类
IOError	输入/输出操作失败
IndexError	序列中没有此索引(index)
KeyError	映射中没有这个键
RuntimeError	一般的运行时错误
'''


'''
1.使用except而不带任何异常类型

你可以不带任何异常类型使用except，如下实例：
try:
    正常的操作
   ......................
except:
    发生异常，执行这块代码
   ......................
else:
    如果没有异常执行这块代码

以上方式try-except语句捕获所有发生的异常。但这不是一个很好的方式，我们不能通过该程序识别出具体的异常信息。因为它捕获所有的异常。


2.使用except而带多种异常类型

你也可以使用相同的except语句来处理多个异常信息，如下所示：

try:
    正常的操作
   ......................
except(Exception1[, Exception2[,...ExceptionN]]]):
   发生以上多个异常中的一个，执行这块代码
   ......................
else:
    如果没有异常执行这块代码

3.try-finally 语句

try-finally 语句无论是否发生异常都将执行最后的代码。

try:
<语句>
finally:
<语句>    #退出try时总会执行

'''


'''

4.触发异常
我们可以使用raise语句自己触发异常
raise语法格式如下：
raise [Exception [, args [, traceback]]]

样例:raise Exception("Invalid level!", level)
'''


'''
5.用户自定义异常
通过创建一个新的异常类，程序可以命名它们自己的异常。异常应该是典型的继承自Exception类，通过直接或间接的方式。
以下为与RuntimeError相关的实例,实例中创建了一个类，基类为RuntimeError，用于在异常触发时输出更多的信息。
在try语句块中，用户自定义的异常后执行except块语句，变量 e 是用于创建Networkerror类的实例。

class Networkerror(RuntimeError):
    def __init__(self, arg):
        self.args = arg

在你定义以上类后，你可以触发该异常，如下所示：

try:
    raise Networkerror("Bad hostname")
except Networkerror,e:
    print e.args
'''


try:
    fh = open("no_this_file", "r")
    cont = fh.read()
except IOError,e:
    print "Error: 没有找到文件或读取文件失败",e.message,e.filename
else:
    print "读取文件成功"
    fh.close()