# -*- coding: utf-8 -*-

'''
print locals() 打印本地变量，在函数第一行可以打印函数的参数

python函数不支持重载，函数支持参数默认值。
在调用设置有默认值的函数时候，可以指定arg=value的形式来执行参数值，也可以顺序来设置参数。
不建议函数默认值为[]，一般可以设置为None，在函数内再判断是否是None，然后再赋值为[]。
'''

def conn(ip="127.0.0.1",port=80):
    print ip,port

conn(port=801)
conn("198.168.1.1",8080)


'''
函数可变参数, type(args)=tuple
'''
def calcsum(*args):
    # print args #(1, 2, 3, 4, 5)
    return sum(args)


print calcsum(1,2,3,4,5)


'''
函数可变参数, type(args)=dict
'''
def calcValueSum(**args):
    # print args  #{'a': 1, 'b': 2}
    t = 0
    for key,value in args.items():
        t += value
    return t

print calcValueSum(a=1,b=2)


'''
可更改(mutable)与不可更改(immutable)对象
在 python 中，strings, tuples, 和 numbers 是不可更改的对象，而 list,dict 等则是可以修改的对象。
不可变类型：变量赋值 a=5 后再赋值 a=10，这里实际是新生成一个 int 值对象 10，再让 a 指向它，而 5 被丢弃，不是改变a的值，相当于新生成了a。
可变类型：变量赋值 la=[1,2,3,4] 后再赋值 la[2]=5 则是将 list la 的第二个元素值更改，本身la没有动，只是其内部的一部分值被修改了。
python 函数的参数传递：
不可变类型：类似 c++ 的值传递，如 整数、字符串、元组。如fun（a），传递的只是a的值，没有影响a对象本身。比如在 fun（a）内部修改 a 的值，只是修改另一个复制的对象，不会影响 a 本身。
可变类型：类似 c++ 的引用传递，如 列表，字典。如 fun（la），则是将 la 真正的传过去，修改后fun外部的la也会受影响
'''


#全局变量
_min =  10
def testGlobal0():
    # 没有声明全部变量，也没有声明局部变量，可以直接使用全局变量的值
    print "testGlobal",_min  #20

testGlobal0();



#全局变量
_min =  10
def testGlobal1():
    _min = 20  #没有声明全局变量，则覆盖全局变量
    print "testGlobal",_min  #20

testGlobal1();
print _min  #10


#全局变量
_min =  10
def testGlobal2():
    global _min   # 声明全局变量
    _min = 20
    print "testGlobal2",_min  #20

testGlobal2();
print _min  #20

