
## 函数

#### Go 语言不支持函数重载，不支持参数默认值，即无法定义两个名称一样的函数，即使它们的参数不一致。
#### go语言支持闭包

### 函数定义
```
1.Go 语言函数定义格式如下：
func function_name( [parameter list] ) [return_types] {
   函数体
}
函数定义解析：
	func：函数由 func 开始声明
	function_name：函数名称，函数名和参数列表一起构成了函数签名。
	parameter list：参数列表，不用var声明，需要指定类型，相同类型可以合并，如 a,b int
	return_types：返回类型，函数返回一列值。return_types 不是必须的。
```

函数样例1：

```
func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

```

### 不定长参数，需要放在参数列表的最后面，无法直接传数组，传数组格式：sum(nums...)。
```
func sum(nums ...int) int {
	var total = 0
	for _, n := range nums {
		total += n
	}
	return total
}

调用样例：
fmt.Println(sum(1, 2, 3, 4, 5))
```

### 函数返回多个值， 返回类型用括号括起
```
func swap(a, b int) (int, int) {
	return b, a
}

调用时候：
a,b = swap(a,b)

如果有的值没有使用到，则可以使用 _ 代替，如：
i,_ = swap(1,2)
```

#### Go 语言使用的是值传递，即在调用过程中不会影响到实际参数

#### Go 语言函数引用传递值
```
/* 定义交换值函数*/
func swap(x *int, y *int) {
   var temp int
   temp = *x    /* 保持 x 地址上的值 */
   *x = *y      /* 将 y 值赋给 x */
   *y = temp    /* 将 temp 值赋给 y */
}

调用样例：
swap(&a, &b)
```

### 函数作为值，匿名函数
```
getMax := func(a,b int)int {
	if a >= b{
		return a
	}
	return b
}

println(getMax(34,45))
```

```
func(a,b int)int {
	if a >= b{
		return a
	}
	return b
}(34,35)

```

### 数组作为函数参数
#### go语言在将数组名作为函数参数的时候，参数传递即是对数组的复制。在形参中对数组元素的修改都不会影响到数组元素原来的值。这点和其他语言不一样。


### 函数可以作为函数的参数 ，函数可以赋值给变量。
```
func add(a int,b int)int  {
	return a+b
}

func opera(a int, b int, sum func(int,int) int) int{
	return sum(a,b)
}
// f := add
println("opera =" ,opera(1,2,add))
```

### Go中如果函数名的首字母大写，表示该函数是公有的，可以被其他程序调用，如果首字母小写，该函数是私有的

###### 函数中 返回值定义中可以定义变量，函数中可以直接使用，在return时候如果不指定return的值，则返回return_types中定义的变量
```
// 返回 1100
func getValue() (x int)  {
	x = 1100
	return
}

```