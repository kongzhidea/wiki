
# 1.文件头编码:<br/>
  # -*- coding: utf-8 -*- <br/>


# 2.python中  unicode字符串不能和普通字符串 join，例如：u"aa" + "空" 会报错， <br/>
可以加上以下代码解决此问题。<br/>
```
import sys
reload(sys)
sys.setdefaultencoding('utf8')
```

# 3.字符串常见操作

#### python字符串可以用  ''，""定义 ，''中可包含"，""中可包含'，不用转义。
#### python字符串可以用  '''      '''来表示，可跨行，不允许修改

#### int 转 str， str(i)
#### str 转 int， int(s)


#### 复制字符串：
```
st = "kongzhihui"
s = st
print(s)
del st
print(st) # s还存在
```

#### 遍历字符串:
```
string="kongzhidea"
for i,item in enumerate(string):
    print i,item  #index value
```

#### 连接字符串：
```
st = "kongzhihui"
s = "kk"
print(s+st)
也可以用+=运算符
```

#### 查找字符串：
```
st = "kongzhihui"
try:
   print(st.index("j")) #如果查找字符串在则返回第一次出现的位置，不过不在则出错，需要用try来捕捉，否则出现fetal级别error
except:
   print("err")
推荐使用find或者用in查询:
方法1：使用 in 方法实现contains的功能：
site = 'http://www.sharejs.com/'
if "sharejs"  in site:
print('site contains sharejs')
输出结果：site contains sharejs 方法2：使用find函数实现contains的功能
方法 2:利用find:
s = "This be a string"
if s.find("is") == -1:
    print "No 'is' here!"
else:
    print "Found 'is' in the string."


s = "kk"
S.find(substr, [start, [end]])
返回S中出现substr的第一个字母的标号，如果S中没有substr则返回-1。
start和end作用就相当于在S[start:end]中搜索
```

#### 字符串比较：
```
st = "kongzhihui"
print(s>st) #返回值为False
3.2中删除了cmp函数
```

#### 字符串长度：
* 如果是非unicode类型，则返回byte数量， utf-8编码1个中文占3个字节
* u"" 字符串表示 unicode类型字符串

```
print(len(st))

len("孔") // 3
len(u"孔") // 1
```

#### 大小写转换：
```
st = "Kongzhihui"
print(st.upper())
print(st.lower())
```

#### 首字母大写：
```
print("kong".capitalize())
ss = "kong zhi dea"
print ss.capitalize()#Kong zhi dea
print ss.title()#Kong Zhi Dea
```

#### 子串：
```
st = "Kongzhihui"
print(st[0:3])
print st[0]
```

#### 翻转字符串：
```
st = "Kongzhihui"
print(st[::-1])

st[start:end]   end取不到
st[start:end:step]
st[pos]
```

#### 分割字符串：
```
st = "Kongzhihui"
print(st.split("h"))#返回值是list
s.split()  #默认按照空白符分割

"a ".split(" ")   #['a', '']，如果末尾有空格，则最后加一个''
```

#### 连接list到字符串：
```
li = ["a","b","c"]
print("-".join(li))
print("".join(li))
```

#### 替换字符串：
```
st = "Kongzhihui"
print(st.replace("h","_"))
直接替换了所有的
```

#### 其他函数：
```
print("13".isdigit())#判断是否全是数字
print("as".isalpha())#判断是否全是字母
print("as2".isalnum())#判断是否全是数字和字母
print("as2".islower())#判断字母是否全是小写
print("as2".isupper()))#判断字母是否全是大写
print("kong".startswith("k")))#判断是否以。。开头
print("kong".endswith("g"))#判断是否以。。结尾

isalpha对unicode字符无法识别，全部返回True，可以再转成str,如  str(u"孔").isalpha() = False
```

#### 字符串对齐
```
ss = "kong"
print ss.ljust(10,".")#kong......
print ss.rjust(10,".")#......kong
print ss.center(10,".")#...kong...
```

#### 计算substr在S中出现的次数
```
print ss.count("k")
```

```
str = "abcd"
for i in str:
    print i
print list(str)  #['a', 'b', 'c', 'd']
print str*2   #abcdabcd
```

## python中string的操作函数

在python有各种各样的string操作函数。在历史上string类在python中经历了一段轮回的历史。在最开始的时候，python有一个专门的string的module，要使用string的方法要先import，但后来由于众多的python使用者的建议，从python2.0开始， string方法改为用S.method()的形式调用，只要S是一个字符串对象就可以这样使用，而不用import。同时为了保持向后兼容，现在的 python中仍然保留了一个string的module，其中定义的方法与S.method()是相同的，这些方法都最后都指向了用S.method ()调用的函数。要注意，S.method()能调用的方法比string的module中的多，比如isdigit()、istitle()等就只能用 S.method()的方式调用。

对一个字符串对象，首先想到的操作可能就是计算它有多少个字符组成，很容易想到用S.len()，但这是错的，应该是len(S)。因为len()是内置函数，包括在__builtin__模块中。python不把len()包含在string类型中，乍看起来好像有点不可理解，其实一切有其合理的逻辑在里头。len()不仅可以计算字符串中的字符数，还可以计算list的成员数，tuple的成员数等等，因此单单把len()算在string里是不合适，因此一是可以把len()作为通用函数，用重载实现对不同类型的操作，还有就是可以在每种有len()运算的类型中都要包含一个len()函数。 python选择的是第一种解决办法。类似的还有str(arg)函数，它把arg用string类型表示出来。

#### 字符串中字符大小写的变换：
```
S.lower() #小写
S.upper() #大写
S.swapcase() #大小写互换
S.capitalize() #首字母大写
String.capwords(S)
```
#### 这是模块中的方法。它把S用split()函数分开，然后用capitalize()把首字母变成大写，最后用join()合并到一起
```
S.title() #只有首字母大写，其余为小写，模块中没有这个方法
```

#### 字符串在输出时的对齐：
```
S.ljust(width,[fillchar])
```

#### 输出width个字符，S左对齐，不足部分用fillchar填充，默认的为空格。
```
S.rjust(width,[fillchar]) #右对齐
S.center(width, [fillchar]) #中间对齐
S.zfill(width) #把S变成width长，并在右对齐，不足部分用0补足
```

#### 字符串中的搜索和替换：
```
S.find(substr, [start, [end]])
```

#### 返回S中出现substr的第一个字母的标号，如果S中没有substr则返回-1。start和end作用就相当于在S[start:end]中搜索
```
S.index(substr, [start, [end]])
```

#### 与find()相同，只是在S中没有substr时，会返回一个运行时错误
```
S.rfind(substr, [start, [end]])
```

#### 返回S中最后出现的substr的第一个字母的标号，如果S中没有substr则返回-1，也就是说从右边算起的第一次出现的substr的首字母标号
```
S.rindex(substr, [start, [end]])
S.count(substr, [start, [end]]) #计算substr在S中出现的次数
S.replace(oldstr, newstr, [count])
```


#### 把S中的oldstar替换为newstr，count为替换次数。这是替换的通用形式，还有一些函数进行特殊字符的替换
```
S.strip([chars])  默认删除空白符（包括'\n', '\r',  '\t',  ' ')
```

#### 把S中前后chars中有的字符全部去掉，可以理解为把S前后chars替换为None
```
S.lstrip([chars])
S.rstrip([chars])
S.expandtabs([tabsize])
```

#### 把S中的tab字符替换没空格，每个tab替换为tabsize个空格，默认是8个
```
字符串的分割和组合：

S.split([sep, [maxsplit]])
```

#### 以sep为分隔符，把S分成一个list。maxsplit表示分割的次数。默认的分割符为空白字符
```
S.rsplit([sep, [maxsplit]])
S.splitlines([keepends])
```

#### 把S按照行分割符分为一个list，keepends是一个bool值，如果为真每行后而会保留行分割符。
```
S.join(seq) #把seq代表的序列──字符串序列，用S连接起来
```


#### 字符串还有一对编码和解码的函数：
```
S.encode([encoding,[errors]])
```

####  其中encoding可以有多种值，比如gb2312 gbk gb18030 bz2 zlib big5 bzse64等都支持。errors默认值为"strict"，意思是UnicodeError。可能的值还有'ignore', 'replace', 'xmlcharrefreplace', 'backslashreplace' 和所有的通过codecs.register_error注册的值。

```
S.decode([encoding,[errors]])
```


#### 编码转换  在python中，使用unicode类型作为编码的基础类型

```
line.decode("gbk").encode("utf-8")
```

#### 是否以prefix开头
S.startswith(prefix[,start[,end]])

#### 以suffix结尾
S.endwith(suffix[,start[,end]])

#### 是否全是字母和数字，并至少有一个字符
S.isalnum()

```
S.isalpha() #是否全是字母，并至少有一个字符
S.isdigit() #是否全是数字，并至少有一个字符
S.isspace() #是否全是空白字符，并至少有一个字符
S.islower() #S中的字母是否全是小写
S.isupper() #S中的字母是否便是大写
S.istitle() #S是否是首字母大写的
```

#### 字符串类型转换函数，这几个函数只在string模块中有：

string.atoi(s[,base])
#### base默认为10，如果为0,那么s就可以是012或0x23这种形式的字符串，如果是16那么s就只能是0x23或0X12这种形式的字符串
```
string.atol(s[,base]) #转成long
string.atof(s[,base]) #转成float
```
