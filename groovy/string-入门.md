
```

Groovy不但涵盖了Java的语法，而且还有增强部分


HelloWorld:
package com.kk.groovy

class VarTest {
     static void main(args) {
        println "hello world"
    }
}

也可以直接写：
println "Hello World!"
然后执行 ： groovy VarTest.groovy
如果不写 class：
    如果别的类没有import自己，可以忽略package，否则package必须。 同级目录可以不填
    package后面的分号 可省略
    如果要执行本类，可以不写main方法，不写class。

    相比起类来说，在groovy脚本中，变量不需要申明（def），在脚本中变量的引用将自动创建并放入Binding中。
    在groovy脚本中，为了编写一个方法，没必要像java一样必须先定义一个类，
    但是和传统的基于class的groovy而言，在类外定义函数需要使用def关键字。
    但是如果你需要一些比如static或者实例变量等等的东西的时候，最好写一个类。

也可以直接执行：
groovy -e "println 'Hello World!'"

在一个groovy文件内，不应该有定义多个static main方法的类，如果有定义多个类，那么应该将有main方法的类放到第一位 

00.首行注释，一般不写。
#! 首行注释，使Unix shell能够定位Groovy启动程序以运行Groovy代码，例如
#!/usr/bin/groovy


000.Groovy会自动导入java.lang.*, java.util.*, java.net.*, java.io.*, java.math.BigInteger, java.math.BigDecimal,   groovy.lang.*, groovy.util.*，
而Java则只自动导入java.lang.*


1、  没有类型的java
作为动态语言，groovy中所有的变量都是对象(类似于.net framework，所有对象继承自java.lang.Object)，在声明一个变量时，groovy不要求强制类型声明，仅仅要求变量名前使用关键字def（从groovy jsr 1开始，在以前的版本中，甚至连def都不需要）。

修改main 方法中的代码：

def var="hello world"
println var
println var.class

你可以看到程序最后输出了var的实际类型为：java.lang.String
作为例外，方法参数和循环变量的声明不需要def。

2、  不需要的public
你可以把main方法前面的public去掉，实际上，groovy中默认的修饰符就是public，所以public修饰符你根本就不需要写，这点跟java不一样。

3、  不需要的语句结束符
Groovy中没有语句结束符，当然为了与java保持一致性，你也可以使用;号作为语句结束符。在前面的每一句代码后面加上;号结束，程序同样正常运行(为了接受java程序员的顽固习惯)。

4.groovy对于对象是什么类型并不关心，一个变量的类型在运行中随时可以改变，一切根据需要而定

5.打印内容

print 输出内容不回车，println输出内容后回车。

println var  或者 println(var)

java方式:System.out.println()


6.boolean类型

true／false
可以使用assert，  如   assert 1>0， 如果断言失败则抛出异常。

7.asType

String a = '78'
int b = a as int
print b



8.字符串
Java中的equals方法对应Groovy中的== , 而Java中的==（判断是否引用同一对象）对应Groovy中的is方法 
'abc"d' 引号中的引号，大中放小，小中放大都可以，类似js
    a.单引号字符串
            无法将变量转义
    b.双引号字符串
            支持变量，"gmtstring ${var}"，还可以支持函数，如${getId()}。
            ${}解析为null，要转义，使用\$即可。
    c.三引号字符串
            支持换行

//int compareToIgnoreCase(String str) 按字典大小比较两个字符串，忽略大小写，返回他们的顺序差值
def str = "a"assert str.compareToIgnoreCase("a") == 0    //相同返回0
assert str.compareToIgnoreCase("A") == 0    //忽略大小写
assert str.compareToIgnoreCase("c") == -2   //返回差值

//Boolean equalsIgnoreCase(String str) 判断两个字符串是否相等，忽略大小写/

/String getAt(int index)   字符串的下标运算符
assert "abcdefg".getAt(2) == "c"
assert "abcdefg"[2] == "c"
assert "abcdefg".getAt(1..2) == "bc"    

//String getAt(Range range)
assert "abcdefg"[1..2] == "bc"

//Int indexOf(String str) 返回给定子字符串在当前字符串中首次出现的索引值
assert "abcdefg".indexOf("b") == 1
assert "abcd".indexOf("g") == -1   

 //如果原字符串中不存在给定子字符串就返回-1

//int length() / int size()   返回字符串的长度
assert "abcd".length() == 4
assert "abcd".size() == 4

//String concat(String str) 在字符串后添加str字符串
assert "ab".concat("12") == "ab12"

//Boolean endsWith(String suffix)  测试字符串是否以给定的后缀结尾
assert "demo1".endsWith("1") == true
// 对应的还有startsWith(String suffix) 测试字符串是否以给定的前缀开头

//String minus(Object value) 删除字符串中value部分
assert "abcd".minus("bc") == "ad"
//String plus(Object valus) 字符串相加
assert "abcd".plus("123") == "abcd123"

//String reverse() 创建当前字符串的逆序字符串
assert "abcd".reverse() == "dcba"

//String substring(int beginIndex) 返回一个当前字符串的指定索引开始的子字符串
assert "abcd".substring(1) == "bcd"

//String substring(int beginIndex,int endIndex) 返回一个当前字符串的指定索引开始的子字符串
assert "abcd".substring(1,2) == "bc"

//Character toCharacter()//Double toDouble()

//Float toFloat()
//Integer toInteger()
//Long toLong() 字符串类型转换

//List toList() 将指定的字符串转换成一个由单个字符组成的字符串列表
assert "abcd".toList() == ["a","b","c","d"]

//String toUpperCase() 将当前字符串对象的所有字符转换为大写
assert "abcd".toUpperCase() == "ABCD"

//String toLowerCase() 将当前字符串对象的所有字符转换为小写
 assert "ABCD".toUpperCase() == "abcd"

//List tokenize()  使用空格作为字符串的分隔符

//List tokenize(String token) 使用指定的token参数作为字符串的分隔符

//String[] split(String regex) 使用与给定的正则表达式相匹配的子字符串将字符串分隔为多个字符串，默认按照空白符

// Boolean matches(String regex) 测试字符串是否匹配给定子字符串
split()返回string[]， tokenize()返回list 

assert 'hi'.matches('hi')assert '111'.matches('\\d{3}')       //'\d':表示匹配一位数字assert '11as11'.matches('\\d{2}.*\\d{2}')

str.trim()  过滤字符串前后空格






```