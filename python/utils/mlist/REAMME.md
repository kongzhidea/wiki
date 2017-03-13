```
help(list)

python中直接打印list会中文乱码，最好写一个方法来打印list:
def showlist(l):
    print "[",
    for i in l:
        print i,",",
    print "]"

列表成员复制:
a[:] = b[:]

定义一个list:
li = ["a","b","c","d"]
#或者  li = list()  或者  li = list(("a","b","c","d")) 或者   li = list(["a","b","c","d"])
print(li)
l = list(li)
print(l)

list索引：
li = ["a","b","c","d"]
print(len(li))
print(li[0])
print(li[-1]) #倒数第一个
print(li[1:2]) #从位置1到位置2-1，b
print(li[-2:-1]) #从位置-2到位置-1-1，c

增加元素：

li = ["a","b","c","d"]
li.append("e") #在后面增加元素
print(li)
li.insert(2,"new") #在位置2前面插入元素
print(li)
li.extend(["e","f"]) #在后面增加一个list的元素
print(li)

搜索list：
li = ["a","b","c","d"]
print(li.index("c")) #2
若搜索的内容不在list中，则出现fetal级别错误

删除元素：
li.remove("b") #删除一个值的首次出现
print(li)
若删除的内容不在list中，则出现fetal级别错误
print(li.pop())#删除最后一个元素并返回其value
print(li)

list运算符：
li = li + ["e","f"] #在后面加上
print(li)
li = li *2 #在后面加上1倍的自身的元素
print(li)

list转成string
print("-".join(li))

分割字符串：
li = ["a","b","c","d","a"]
print("-".join(li))
str = "-".join(li)
print(str.split("-"))

列表遍历：
for i in li:
   print(i)
输出的是value

判断元素是否在list中：
if "a" in li:
   print("a")


高级操作：
listone = [2, 3, 4]
listtwo = [2*i for i in listone if i > 2]
print listtwo
这里我们为满足条件（if i > 2）的数指定了一个操作（2*i），从而导出一个新的列表。注意原来的列表并没有发生变化。

求和
li = [1,2,3,4,5]
print sum(li)


计算list的长度，最大值，最小值
li = [1,2,3,4,5]
print len(li)
print max(li)
print min(li)

列表反转:
li.reverse()   #li本身反转
li = list(reversed(l)) #内置 reversed函数，l本身不变


enumerate()
A new built-in function, enumerate() , will make certain loops a bit clearer. enumerate(thing) , where thing is either an iterator or a sequence, returns a iterator that will return (0, thing [0]) , (1, thing [1]) , (2, thing [2]) , and so forth.
A common idiom to change every element of a list looks like this:

for i in range(len(L)):
item = L[i]
# ... compute some result based on item ...
L[i] = result
This can be rewritten using enumerate() as:

for i, item in enumerate(L):
# ... compute some result based on item ...
L[i] = result

example:

#!/usr/bin/env python
string = 'mytest'

for i in range(len(string)):
print string[i],'(%d)' %i

for i,item in enumerate(string):
print item,'(%d)' %i





```