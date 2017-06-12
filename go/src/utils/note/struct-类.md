
## struct/类操作

### struct定义
* interface{}类型 和java中的Object类似，标示任意类型。

```
type struct_variable_type struct {
   member definition;
   member definition;
   ...
   member definition;
}
```
### struct声明
```
variable_name := structure_variable_type {value1, value2...valuen}

普通的 struct（非指针类型）的对象不能赋值为 nil，也不能和 nil 进行判等（==）
```

### demo
```
type Book struct {
	id   int
	name string
}

func main() {
	var book = Book{id: 1, name: "k"} // 也可以初始只定义id， 
	// var book Book   // 只声明，不初始化，属性均为其默认值
	// var book = Book{1 ,"k"}   // 此方式初始化需要按顺序传所有参数
	// var book = Book{}  // 空构造函数
	
	book.id = 2  // 赋值
	//book.name = "kk"
	
	println(book.name)  // 获取值
}

```

### struct 指针
```
p := new(Book)

或者
var p *Book = new(Book)
var p *Book = &book

使用结构体指针访问结构体成员，使用 "." 操作符，如：
p.id = 1

获取指针： p := &book
```

### 打印struct（没有String方法）
```
// 不能使用println，需要使用fmt.Println()

fmt.Println(book) // 只打印值，如 {2 kk}

fmt.Printf("%+v\n", book)  // 打印 key:value  {id:2 name:kk}
	
//%#v 形式则输出这个值的 Go 语法表示
fmt.Printf("%#v\n", book)   // 打印： main.Book{id:2, name:"kk"}	
// %T 打印类型
fmt.Printf("%T\n", book)    // 打印： main.Book
```

### struct的String方法
```
当struct没有绑定String方法时候，打印时候输出如上。
当绑定String方法时候，打印的时候 调用String方法，如：

func (book Book) String() string  {
	return "(" + strconv.Itoa(book.id) + "/" + book.name + ")"
}

不要再String中打印本身，否则将循环调用。

```

###  函数中传struct
* 默认函数中传递的为stuct拷贝对象，在函数内修改struct变量不会影响函数外变量的值，可以通过传递struct指针的方式来解决此问题。

```
func show(book Book)  {
		book.id = 1 // 函数外 book值不变
}
func show2(book *Book)  {
		book.id = 1 // 函数外 book值修改
}

调用方式：
show2(p)或者 show2(&book)

```


### tag
```
可以为struct域加说明，这些说明可以通过反射获取

type TagType struct { // tags
   field1 bool “An important answer”
   field2 string “The name of the thing”
   field3 int “How much there are”
}
```

### 反射获取struct结构
```
// 参数p为指针，如 p=&Book{}
// clz := reflect.TypeOf(p).Elem() //通过反射获取type定义
clz := reflect.TypeOf(book) //通过反射获取type定义

for i := 0; i < clz.NumField(); i++ {
	fmt.Printf("name=%s", clz.Field(i).Name)
	fmt.Printf(",type=%s", clz.Field(i).Type)
	fmt.Printf(",tag=%s\n", clz.Field(i).Tag)
}
```
```
clz.Name()  // 返回struct 类名
clz.NumField()  // 返回field数量
field , exists := clz.FieldByName("fieldName")  // 根据名称查找field 
field为StructField类型，包含 Name,Type,Tag等属性
```

```
获取变量的类型：
reflect.TypeOf(1)  // int

// %T 打印类型
fmt.Printf("v is of type %T\n", 1)  // int
```

### struct 定义函数
* **可以定义引用函数，也可以定义值函数**

```
func (recv receiver_type) methodName(parameter_list) (return_value_list) { … }

如：
func (p *Book) setName(name string)   {
	p.name = name
}

func (p *Book) getName() string  {
	return p.name
}

调用样例：
book.setName("hehe")
fmt.Println(book.getName())

p := &book
p.setName("xixi")
fmt.Println(p.getName())
```

### 基本类型，定义别名
```
type Integer int

func (m Integer) Add(i int) Integer {
	return Integer(int(m) + i)
}

使用：
var it Integer = Integer(1) // 定义Integer类型，int转Integer
it2 := it.Add(2)

i3 := int(it2) // Integer类型转int类型
```


### 匿名字段
* 结构体可以包含一个或多个 匿名（或内嵌）字段，即这些字段没有显式的名字，只有字段的类型是必须的，此时类型就是字段的名字。匿名字段本身可以是一个结构体类型，即 结构体可以包含内嵌结构体。
* 在一个结构体中对于每一种数据类型只能有一个匿名字段。

```
package main

import "fmt"

type innerS struct {
    in1 int
    in2 int
}

type outerS struct {
    b    int
    c    float32
    int  // anonymous field
    innerS //anonymous field
}

func main() {
    outer := new(outerS)
    outer.b = 6
    outer.c = 7.5
    outer.int = 60
    outer.in1 = 5
    outer.in2 = 10

    fmt.Printf("outer.b is: %d\n", outer.b)
    fmt.Printf("outer.c is: %f\n", outer.c)
    fmt.Printf("outer.int is: %d\n", outer.int)
    fmt.Printf("outer.in1 is: %d\n", outer.in1)
    fmt.Printf("outer.in2 is: %d\n", outer.in2)

    // 使用结构体字面量
    outer2 := outerS{6, 7.5, 60, innerS{5, 10}}
    // outer2 := outerS{b:6, c:7.5, int:60, innerS:innerS{5, 10}}
    fmt.Println("outer2 is:", outer2)
}
```

### interface 接口/多态
* 接口定义了一组方法（方法集），但是这些方法不包含（实现）代码：它们没有被实现（它们是抽象的）。接口里也不能包含变量。
* Go 语言中的接口都很简短，通常它们会包含 0 个、最多 3 个方法。
* 类型不需要显式声明它实现了某个接口：接口被隐式地实现。多个类型可以实现同一个接口。
* 实现某个接口的类型（除了实现接口方法外）可以有其他的方法。
* 一个类型可以实现多个接口。
* **接口可以接受引用（指针），也可以接受值，主要看定义函数时候 定义的是引用还是值。**
* **一个类型认为实现某接口，需要实现其接口所有方法**

```
type Namer interface {
    Method1(param_list) return_type
    Method2(param_list) return_type
    ...
}
```


#### interface.demo

```
package main

import "fmt"

type Shaper interface {
    Area() float32
}

type Square struct {
    side float32
}

func (sq *Square) Area() float32 {
    return sq.side * sq.side
}

type Rectangle struct {
    length, width float32
}

func (r Rectangle) Area() float32 {
    return r.length * r.width
}

func main() {

    r := Rectangle{5, 3} // Area() of Rectangle needs a value
    q := &Square{5}      // Area() of Square needs a pointer
    // shapes := []Shaper{Shaper(r), Shaper(q)}
    shapes := []Shaper{r, q}
    fmt.Println("Looping through shapes for area ...")
    for _, shape := range shapes {
        fmt.Println("Shape details: ", shape)
        fmt.Println("Area of this shape is: ", shape.Area())
    }
}

```

### interface在函数中作为参数
```
package main

import "fmt"

type stockPosition struct {
    ticker     string
    sharePrice float32
    count      float32
}

/* method to determine the value of a stock position */
func (s stockPosition) getValue() float32 {
    return s.sharePrice * s.count
}

type car struct {
    make  string
    model string
    price float32
}

/* method to determine the value of a car */
func (c car) getValue() float32 {
    return c.price
}

/* contract that defines different things that have value */
type valuable interface {
    getValue() float32
}

func showValue(asset valuable) {
    fmt.Printf("Value of the asset is %f\n", asset.getValue())
}

func main() {
    var o valuable = stockPosition{"GOOG", 577.20, 4}
    showValue(o)
    o = car{"BMW", "M3", 66500}
    showValue(o)
}
```

### 接口嵌套
* 一个接口可以包含一个或多个其他的接口，这相当于直接将这些内嵌接口的方法列举在外层接口中一样。
* 如下，接口 File 包含了 ReadWrite 和 Lock 的所有方法，它还额外有一个 Close() 方法。

```
type ReadWrite interface {
    Read(b Buffer) bool
    Write(b Buffer) bool
}

type Lock interface {
    Lock()
    Unlock()
}

type File interface {
    ReadWrite
    Lock
    Close()
}
```

### [接口类型判断，强制转为原来类型](http://wiki.jikexueyuan.com/project/the-way-to-go/11.3.html)
```
转换形式：v, ok := varI.(T)

样例如下：

if t, ok := shape.(Square); ok {
	fmt.Printf("The type of shape is: %T\n", t)
	fmt.Println("square.name=",t.name)
}else {
	fmt.Println("areaIntf does not contain a variable of type Circle")
}

```
### 接口: type-swtich
```
switch t := areaIntf.(type) {
case *Square:
    fmt.Printf("Type Square %T with value %v\n", t, t)
case *Circle:
    fmt.Printf("Type Circle %T with value %v\n", t, t)
case nil:
    fmt.Printf("nil value: nothing to check?\n")
default:
    fmt.Printf("Unexpected type %T\n", t)
}
```

```

interface{} 转为string类型：
    var obj interface{} = "abcd"
    val, ok = obj.(string)    // ok为true/false，表示转换成功或失败

interface{} 转为string类型：
    var obj interface{} = 11
    obj.(int)


obj.(type) 只能在switch中使用
```


### 判断interface{} 是否实现某接口
```
var obj interface{} = shape

if sv, ok := obj.(Shaper); ok {
	fmt.Printf("v implements String(): %+v\n", sv) // note: sv(Shape类型), not v
}
```
