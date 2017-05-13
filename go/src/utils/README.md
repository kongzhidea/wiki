
# [go语言入门](http://www.runoob.com/go/go-tutorial.html)

### hello world

```
package main

func main() {
	println("hello world")
}

```

### go环境配置
```
http://www.runoob.com/go/go-environment.html

安装包下载地址为：
	https://golang.org/dl/。
	http://golangtc.com/download

mac下go位置：GOROOT="/usr/local/go"

linux下设置go：
	export GOROOT="/usr/local/go"
    export GOPATH=$GOPATH:"/Users/kongzhihui/workspace/go"
	export PATH=$PATH:$GOROOT/bin

	GOROOT为系统自带包，GOPATH为自定义依赖包。

```
### 推荐IDE：[pycharm](http://www.cnblogs.com/NerdWill/p/5479551.html)、idea。
### 推荐编码：utf-8
* mac下 pycharm 返回到上/下一个页面快捷键: alt+command+left/right

### 运行
```
go run hello.go

package main就必须包含一个main函数，程序的初始化和执行都起始于main包。

go version 查看当前go版本
go env 查看当前go的环境变量
go fmt [file] 格式化代码，不带file参数时候为当前目录下go文件。

go build hello.go  编译go文件到可执行文件
./hello 执行

从github上安装依赖，安装在GOPATH目录下，如：
go get github.com/nu7hatch/gouuid
```

### 需要注意地方
```
1.go语言要求 '{' 一定要在行末，不能另起一行。
2.局部变量声明后一定要使用，变量使用 _ 不强制要求使用。
3.import 包引入后一定要使用， 如果不想使用可以采用：import _ "math"，不强制要求使用。
4.不能用nil初始化无类型变量
```

### 打印到控制台
#### fmt.Println 可以打印数组，println无法打印数组
* 方案1

```
args可以使多个参数

print(args)  不换行
println(args) 换行
```

* 方案2 

```
args可以使多个参数

import "fmt"

fmt.Print(args)  不换行
fmt.Println(args)  换行
```

* 方案3
```
fmt.Printf，可以使用占位符，如%d，%s。 
%d使用int，%s使用string。 int无法使用%s占位符，如果类型错误则无法正确显示。  
```

### 注释
```
单行注释   //
多行注释   /* */
```

### 分段问题
```
每行最后 分号可省略，如果一行有多个语句则可以用分号隔开。
代码块 用{}分隔。
```

### 变量定义，Go 语言程序中全局变量与局部变量名称可以相同，但是函数内的局部变量会被优先考虑。
```
1.变量类型：byte, int32, int64, string, bool，
int 默认为int64。

8进制数字  0
16进制数字  0x
2进制数字 暂不支持， java8支持 0b写法。

go语言中局部变量声明后必须使用，否则编译失败。全部变量无此限制。
go语言变量可以用 _ 代替， 表示占位，不受局部变量必须使用的限制。

2.定义变量 
    a. 指定类型：var identifier type = value  
       如 var age int = 10， 有type时候 value也缺省，为默认值。
	 b. 自行判定变量类型：var identifier = value
	    如 var age = 10，无type时候， value不也缺省。
	 c. 自行判定变量类型，省略var, 注意 :=左侧的变量不应该是已经声明过的，否则会导致编译错误
	 	v_name := value，如 age:=10
	 	此类声明 只能在函数体内，不能用于全局变量生命。
	 d.不能用nil初始化无类型变量
3.多变量声明

//类型相同多个变量, 非全局变量
var vname1, vname2, vname3 type
vname1, vname2, vname3 = v1, v2, v3

var vname1, vname2, vname3 = v1, v2, v3 //和python很像,不需要显示声明类型，自动推断

vname1, vname2, vname3 := v1, v2, v3 //出现在:=左侧的变量不应该是已经被声明过的，否则会导致编译错误


// 这种因式分解关键字的写法一般用于声明全局变量
var (
    vname1 v_type1
    vname2 v_type2
    vname3 v_type3 = v3
)	 	
```

### 常量
```
1.常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。

2.常量的定义格式：
const identifier [type] = value

你可以省略类型说明符 [type]，因为编译器可以根据变量的值来推断其类型。
	显式类型定义： const b string = "abc"
	隐式类型定义： const b = "abc"
	相同类型合并声明：var a,b int
	
3.多个相同类型的声明可以简写为：
	const c_name1, c_name2 [type] = value1, value2

4.因式分解关键字的写法一般用于声明全局变量
const (
    vname1 v_type1
    vname2 v_type2
    vname3 v_type3 = v3
)	
```

### 运算符
```
1.支持 ++,-- 自增操作，只支持后++，如 i++，不支持++i
2.支持 +=，-= 等操作
3.逻辑运算符  && || !
4.布尔运算符  bool： true，false。
```

### 结构分支
```
1.if
if (condition){      
	...
} else if (condition){
	...
} else{
	...
}

if condition{    // 括号可省略
	...
}else if condition{
	...
} else{
	...
}

2.switch 语句执行的过程从上至下，直到找到匹配项，匹配项后面也不需要再加break

switch var1 {
    case val1:
        ...
    case val2:
        ...
    default:
        ...
}
```

### for
```
1.for init; condition; post { }  没有括号
如下，for中重新定义变量i 在for内会覆盖外部变量，for外则为原变量。

for i := 1; i < 4; i++ {
	println(i)
}

2.for condition { }   类似C语言中while， condition外可以加括号，也可以缺省。
如下：
for a < b {
	println(a, b)
	a++
}

3.for { }   类似C语言中 for(;;){}
for {
	if i >= 10 {
		break
	}
	i++
}

4.遍历数组
numbers := [6]int{1, 2, 3, 5}
for i, x := range numbers {
	println("第 %d 位 x 的值 = %d\n", i, x)   // i从0开始
}

5.如果想忽略 index，则可以采用以下写法：
x := []string{"a","b","c"}

for _, v := range x {
    fmt.Println(v) //prints a, b, c
}
	
```