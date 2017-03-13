
```
#定义dict
dt = dict()
dt = {}

#字典的添加、删除、修改操作
dict = {"a" : "apple", "b" : "banana", "g" : "grape", "o" : "orange"}
dict["w"] = "watermelon"
del(dict["a"])
dict["g"] = "grapefruit"
print dict.pop("b") #pop必须有参数，返回是对应的value
print dict
dict.clear()
print dict
#使用dict[key]方式访问，如果key不存在则抛出异常，使用get不会
#字典的遍历
dict = {"a" : "apple", "b" : "banana", "g" : "grape", "o" : "orange"}
for k in dict:
   print "dict[%s] =" % k,dict[k] #dict的遍历和list不一样，返回的key，而list返回的是value

for k,v in dict.iteritems():
  print k,v

#字典items()的使用
dict = {"a" : "apple", "b" : "banana", "c" : "grape", "d" : "orange"}
#每个元素是一个key和value组成的元组，以列表的方式输出
print dict.items()
#调用items()实现字典的遍历
dict = {"a" : "apple", "b" : "banana", "g" : "grape", "o" : "orange"}
for (k, v) in dict.items():
   print "dict[%s] =" % k, v


#元祖列表转换为字典
d = dict([('a', 11), ('b', 22), ('c', 33), ('d', 44)])
#使用列表、字典作为字典的值
dict = {"a" : ("apple",), "bo" : {"b" : "banana", "o" : "orange"}, "g" : ["grape","grapefruit"]}
print dict["a"]
print dict["a"][0]
print dict["bo"]
print dict["bo"]["o"]
print dict["g"]
print dict["g"][1]

dict = {"a" : "apple", "b" : "banana", "c" : "grape", "d" : "orange"}
#输出key的列表
print dict.keys()
print list(dict.keys())
#输出value的列表
print dict.values()
#每个元素是一个key和value组成的元组，以列表的方式输出
print dict.items()

#字典中元素的获取方法
dict = {"a" : "apple", "b" : "banana", "c" : "grape", "d" : "orange"}
print dict
print dict.get("c", "apple")
print dict.get("e", "apple")
#get()的等价语句
D = {"key1" : "value1", "key2" : "value2"}
if "key1" in D:
   print D["key1"]
else:
   print "None"

#字典的更新
dict = {"a" : "apple", "b" : "banana"}
print dict
dict2 = {"c" : "grape", "d" : "orange"}
dict.update(dict2)
print dict
#udpate()的等价语句
D = {"key1" : "value1", "key2" : "value2"}
E = {"key3" : "value3", "key4" : "value4"}
for k in E:
   D[k] = E[k]
print D

#调用sorted()排序
dict = {"a" : "apple", "b" : "grape", "c" : "orange", "d" : "banana"}
print dict
#按照key排序
print sorted(dict.items(), key=lambda d: d[0])
#按照value排序
print sorted(dict.items(), key=lambda d: d[1])

#字典的浅拷贝
dict = {"a" : "apple", "b" : "grape"}
dict2 = {"c" : "orange", "d" : "banana"}
dict2 = dict.copy()
print dict2

popitem()
删除并返回dict中任意的一个(key,value)队，如果字典为空会抛出KeyError
Python代码
d = {"name":"nico", "age":23}
print d.popitem()       #('age', 23)
print d.popitem()       #('name', 'nico')
#此时字典d已为空
print d.popitem()      #此处会抛出KeyError

字典长度
print(len(dict))

has_key() 函数用于判断键是否存在于字典中，如果键在字典dict里返回true，否则返回false。


dict 排序

# -*- coding: utf-8 -*-

#最简单的方法，这个是按照key值排序：
def sortedDictValues1(adict):
    items = adict.items()
    items.sort()
    return [(key,value) for key, value in items]

#又一个按照key值排序
def sortedDictValues2(adict):
    keys = adict.keys()
    keys.sort()
    return [(key,adict[key]) for key in keys]

#还是按key值排序，只返回value列表
def sortedDictValues3(adict):
    keys = adict.keys()
    keys.sort()
    return map(adict.get, keys)

#一行语句搞定：，返回列表,每个元素为tuple,按照key排序
def sortedDictValues4(adict):
    return [(k,dit[k]) for k in sorted(dit.keys())]

#来一个根据value排序的，先把item的key和value交换位置放入一个list中，再根据list每个元素的第一个值，即原来的value值，排序：，返回的是(value,key)列表,按照value排序
def sort_by_value(d):
    items=d.items()
    backitems=[[v[1],v[0]] for v in items]
    backitems.sort()
    #return [ backitems[i][1] for i in range(0,len(backitems))]
    return backitems

#得到value排序，只返回values
def sortedDictByValues2(adict):
    return [ v for v in sorted(adict.values())]

#用lambda表达式来排序，按照value来排序
def sortedDict2(d):
    return sorted(d.items(), lambda x, y: cmp(x[1], y[1]))

#用lambda表达式来排序，按照value来排序 逆序
def sortedDict3(d):
    return  sorted(d.items(), lambda x, y: cmp(x[1], y[1]), reverse=True)

#用sorted函数的key= 参数排序：
# 按照key进行排序
#print sorted(dict1.items(), key=lambda d: d[0])
# 按照value进行排序
#print sorted(dict1.items(), key=lambda d: d[1])

```