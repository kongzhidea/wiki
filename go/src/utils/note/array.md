## 数组

### 数组定义
```
1.声明数组，不初始化，均为默认值
var variable_name [SIZE] variable_type

如：var iar[10] int， iar中值均为默认值0。

2.声明数组同时初始化，指定长度
var variable_name = [SIZE] variable_type{v1,v2,v3}

如：var array3 = [6]int{9, 10, 11, 12}    其他未指定值为默认值。
array3[4] array3[4] 值为0。

3.声明函数同时初始化，根据元素的个数来设置数组的大小。
var variable_name = [] variable_type{v1,v2,v3}
var variable_name = [...] variable_type{v1,v2,v3}
```

### 打印数组
```
1.无法使用 println()打印数组元素
2.fmt.Printf(arr) 打印数组元素
3.fmt.Printf("%v\n",arr)  打印数组元素
4.fmt.Printf("%T\n",arr)  打印数组类型
```

### 数组作为函数参数
#### go语言在将数组名作为函数参数的时候，参数传递即是对数组的复制。在形参中对数组元素的修改都不会影响到数组元素原来的值。这点和其他语言不一样。
```
void myFunction(param [10]int){
}

直接传递数组，需要指定param长度。
如果传递数组切片，则不需要指定param长度。%T
```

### 想达到更新数据x数据目的，传数组指针
```
func sol_array2(iar* [10]int) {
    iar[0] = 10
	//(*iar)[0] = 10  // 此写法也可以
}

sol_array2(&iar)   // 调用方式

```

### 传递数组切片，在函数中不需要执行数组长度，在函数中修改数组，函数外原数组内容也会跟着修改
```
func solarr(arr []int){
}

调用方式:
solarr(arr[:])
```

### 数组赋值，数组切片的本质就是指向数组的指针
```
浅拷贝：
1.arr2 := arr1[:]   // 数组
2.arr2 := &arr1     // 数组指针

如：
var arr1 = []int{1,2,3}
arr2 := arr1[:]
arr2[1] = 100
arr1[0] = 0
fmt.Println(arr1,arr2)

// 打印结果
[0 100 3] [0 100 3]
```
```
如果是数组直接赋值：
arr2 := arr1
此时需要看 arr1声明时候是否指定长度，如果指定长度则是深拷贝，
没有指定长度（切片）则是浅拷贝。
```

### 数组操作
```
1. 数组长度：  len(arr)
   cap() 可以测量切片最长可以达到多少
2. 切片
arr[st:ed]  从st到ed，不包含ed。
ed不能去-1，不能超过数组长度。

arr[:] 全部切片。
arr[st:] end为len(s)。
arr[:end] st为0。

3. 数组追加元素
   arr = append(arr, item1,item2...) 
   如：numbers = append(numbers, 2, 3, 4)
```

```
/* 默认下限为 0*/
fmt.Println("numbers[:3] ==", numbers[:3])

/* 默认上限为 len(s)*/
fmt.Println("numbers[4:] ==", numbers[4:])
```

### 遍历数组
```
1.遍历数组
numbers := [6]int{1, 2, 3, 5}
for i, x := range numbers {
	println("第 %d 位 x 的值 = %d\n", i, x)   // i从0开始
}

2.如果想忽略 index，则可以采用以下写法：
x := []string{"a","b","c"}

for _, v := range x {
    fmt.Println(v) //prints a, b, c
}

```

### 遍历字符串
```
for i, c := range "go语言" {
	fmt.Println(i, string(c))
}

c 为rune类型，i为字节位置。
输出结果：
0 g
1 o
2 语
5 言
```

### 数组反转
```
func ReverseIntArray(numbers []int) []int {
	newNumbers := numbers[:]
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		newNumbers[i], newNumbers[j] = numbers[j], numbers[i]
	}
	return newNumbers
}
```

### make
```
make()函数在golang的代码如下：
    func make(t Type,size IntegerType) Type

使用make来创建slice，map，chanel说明如下：
    var slice_ []int = make([]int,5,10)
    fmt.Println(slice_)

    var slice_1 []int = make([]int,5)
    fmt.Println(slice_1)

    var slice_2 []int = []int{1,2}
    fmt.Println(slice_2)

打印结果：
    [0 0 0 0 0]
    [0 0 0 0 0]
    [1,2]
```