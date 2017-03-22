
## 接收main参数
```
在groovy脚本中，不写main方法，获取命令行参数。

println this.args.size()

println this.args

如： groovy test.groovy a b
结果为：
2
[a,b]

main参数中不包含 文件名。
```

## groovy脚本
```
当做脚本可以直接写：
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

groovy 依赖第三方jar包:
groovy -cp .:lib/* test.groovy

```

```
Groovy会自动导入java.lang.*, java.util.*, java.net.*, java.io.*, java.math.BigInteger, java.math.BigDecimal,   groovy.lang.*, groovy.util.*，

而Java则只自动导入java.lang.*
```