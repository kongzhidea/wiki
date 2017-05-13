
## map

### map声明和定义
* map 无序
* interface{}类型 和java中的Object类似，标示任意类型。
    * dict := map[string]interface{}{} 

```
/* 声明变量，默认 map 是 nil，无法赋值 */
var map_variable map[key_data_type]value_data_type

/* 使用 make 函数 */
map_variable = make(map[key_data_type]value_data_type)

// 声明时候同时赋值
map_variable := map[key_data_type]value_data_type{key1:val1,key2:val2}
```

### [dict用法，对map封装](https://github.com/gosimple/listdict/blob/master/dict.go)
* 参考 listdict/Dict

### demo
```
var mat = map[string]string{}
fmt.Println(mat)

// 先声明map
var m1 map[string]string
// 再使用make函数创建一个非nil的map，nil map不能赋值
m1 = make(map[string]string)
// 最后给已声明的map赋值
m1["a"] = "aa"
m1["b"] = "bb"

// 直接创建
m2 := make(map[string]string)
// 然后赋值
m2["a"] = "aa"
m2["b"] = "bb"

// 初始化 + 赋值一体化
m3 := map[string]string{
	"a": "aa",
	"b": "bb",
}


// ==========================================
// 查找键值是否存在
if v, ok := m1["a"]; ok {
	fmt.Println(v)
} else {
	fmt.Println("Key Not Found")
}

// 遍历map， key/value
for k, v := range m1 {
	fmt.Println(k, v)
}

// 遍历map， key
for k := range m1{
	fmt.Println(k, m1[k])
}
	
fmt.Println(m3)
```

### 打印map
```
1.fmt.Println(dict)

2.fmt.Printf("%v\n",dict)
```


### 获取值
```
1.value,exists = mat[key]   //key不存在时候，exists为false，value为其类型默认值。

2.val = mat[key]   // key不存在时候返回其类型默认值。
```

### 删除值 
```
delete() 函数用于删除集合的元素, 参数为 map 和其对应的 key。

如：
delete(dict, key)
```

### map当 函数参数传递，在函数内修改map值，函数外可以生效
```
func func_name(dict map[string]string){

}
```

### 获取map的key列表
```
func mapKeys(dict map[string]string) []string  {
	keys := make([]string, len(dict))
	key_idx := 0
	for k, _ := range dict {
		keys[key_idx] = k
		key_idx++
	}
	return keys
}
```

### 对string数组排序
```
sort.Strings(keys)
```

### 获取key排序后的map
```
keys := mapKeys(dict)
sort.Strings(keys)
for _, k := range keys {
	fmt.Printf("Key: %v, Value: %v \n ", k, dict[k])
}
```


### 空接口--interface{}
```
package main
import "fmt"

var i = 5
var str = "ABC"

type Person struct {
    name string
    age  int
}

type Any interface{}

func main() {
    var val Any
    val = 5
    fmt.Printf("val has the value: %v\n", val)
    val = str
    fmt.Printf("val has the value: %v\n", val)
    pers1 := new(Person)
    pers1.name = "Rob Pike"
    pers1.age = 55
    val = pers1
    fmt.Printf("val has the value: %v\n", val)
    switch t := val.(type) {
    case int:
        fmt.Printf("Type int %T\n", t)
    case string:
        fmt.Printf("Type string %T\n", t)
    case bool:
        fmt.Printf("Type boolean %T\n", t)
    case *Person:
        fmt.Printf("Type pointer to Person %T\n", t)
    default:
        fmt.Printf("Unexpected type %T", t)
    }
}
```