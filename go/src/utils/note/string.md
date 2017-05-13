
## go字符串操作
* [strings](https://golang.org/pkg/strings/)
* [strings中文版](https://github.com/astaxie/gopkg/tree/master/strings)

#### go语言中，字符串内容在初始化后不可修改，编码为UTF-8，注意文件编码格式改成UTF-8。

### 初始化
* var s string   // 默认初始值为  ""

```
var str string = "hello"

var ch byte = str[0]
println("ch=",ch) // ch=104

len := len(str)
println("len=",len)
```

### format
```
value := fmt.Sprintf(format,args...)

如：value := fmt.Sprintf("i am %s,age=%d","kk",21)
```

### 字符串--引号
```
1.字符串用双引号定义，不能使用单引号

2.折行字符串展示  使用：``，如：
strFormat := `
	Cannot proceed, the divider is zero.
	dividee: %d
	divider: 0
`
```

### 字节和字符串互转
```
bytes := []byte(str)

str = string(bytes)  // string函数为 字节或者 rune类型转字符串
```

### rune(int32位的别名)  字符
```
r := []rune(str)

fmt.Println(string(r))

修改字符串内容：
r[0]='孔'  // 注意，单引号，  但是只能修改r数组中内容，不能修改str内容。

var r rune = '孔'
如果是ascii码 可以用byte: var b byte  = 'a'
```

### 识别中文
* Go中需要将字符串转换成rune数组，runne数组中就可以通过数组下标获取一个汉字所标识的Unicode码，再将Unicode码按创建成字符串即可
* runne 

```
str := "kk孔"

r := []rune(str)
fmt.Println("rune=", r)
for i := 0; i < len(r); i++ {
	fmt.Print(string(r[i]), ",")
}
fmt.Println()
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

### 长度

* len(str)  对应的字节(byte)长度
    * 中文 utf-8编码，1个中文占用3个字节
    * 如：len("kk孔") // 5
* utf8.RuneCountInString(str) 返回unicode字符长度，1个中文占1个长度。
    * import "unicode/utf8"
    


## 常见操作

### int和string互转
```
import "strconv"

#int到string  推荐  
string:=strconv.Itoa(int)  

#string到int   推荐
int,err:=strconv.Atoi(string)  

#int64到string  
string:=strconv.FormatInt(int64,10)  

#string到int64  
int64, err := strconv.ParseInt(string, 10, 64) 

如：
i,_ := strconv.Atoi("1234")
```

### 进制转换
```
a:=strconv.FormatInt(i64,2) // 10进制转成2进制，返回字符串

i64,_  := strconv.ParseInt(str,2,64)  //2进制转成10进制，返回int64
```

### 字符串比较
```
s1 == s2    // 字符串比较

strings.EqualFold(s1,s2)  // 忽略大小写
```

### 字符串拼接
* str = str + substr
* str += substr
* strings.Repeat(s string, count int) // 字符串s 重复count次

### 大小写转换
```
strings.ToLower(s)  // 转小写，支持中文
strings.ToUpper(s)  // 转大写，支持中文
```

### 判断是否是空串
```
方式1：s == ""
方式2：len(s) == 0
```

### 字符串查找
* string.Index 采用 Rabin-Karp 算法，通过hash的方式来检索

```
import "strings"

strings.Contains(s,substr)  // true表示包含子串
string.Index(s,substr)   // 返回子串在s中位置，-1表示没有找到

go原因处理时候当字节处理，1个utf-8中文当3个字节处理，如strings.Index("孔1", "1") 返回3

strings.Contains(str, "")  // 返回true
strings.Index(str,"") // 返回0

string.ContainsRune(s,r rune) // true标示包含此rune字符。
strings.IndexRune(s,r rune) // 返回字符换包含rune字符个数

strings.Count(str,substr)  // 返回str中子串的个数
```

* 考虑中文的Index

```
func UnicodeIndex(str, substr string) int {
	// 子串在字符串的字节位置
	result := strings.Index(str, substr)
	if result >= 0 {
		// 获得子串之前的字符串并转换成[]byte
		prefix := []byte(str)[0:result]
		// 将子串之前的字符串转换成[]rune
		rs := []rune(string(prefix))
		// 获得子串之前的字符串的长度，便是子串在字符串的字符位置
		result = len(rs)
	}

	return result
}

```

### 字符串截取
```
str[st:ed]    // 从st开始，不包含ed
ed 不能取-1，不能超过数组长度。
截取的时候 1个中文当3个字节处理，故处理中文可能会乱码。
```
* 考虑中文的字符串截取

```
// begin~end，不包含end
func SubString(str string, begin, end int) (substr string) {
	// 将字符串的转换成[]rune
	rs := []rune(str)
	lth := len(rs)

	// 简单的越界判断
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	if end > lth {
		end = lth
	}

	// 返回子串
	return string(rs[begin:end])
}

// 从begin开始后length字符，不包含begin+length
func SubString0(str string, begin, length int) (substr string) {
	// 将字符串的转换成[]rune
	rs := []rune(str)
	lth := len(rs)

	// 简单的越界判断
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length
	if end > lth {
		end = lth
	}

	// 返回子串
	return string(rs[begin:end])
}
```

### 获取字符串sep在字符串s中出现的次数 string.Count(s,sep string)
* 注意：如果sep=""，则无论s为何字符串都会返回 len(s)+1

### 判断字符串以prefix开头  strings.HasPrefix(s,prefix)

### split 切分字符串
* strings.Fields(str)   按照空白符切分
* strings.Split(s,seq)  按照seq切分字符串，seq为连续子串

### 拼接字符串
* strings.Join(a []string, sep string)

### 字符串替换
* strings.Replace(s, old, new string, n int) // n表示替换次数，-1表示全部

### trim操作
* strings.Trim(s string, cutset string) // 在s中左右删除cutset中所有字符
* strings.TrimLeft(s,c),  strings.TrimRight(s,c)  // 左右trim
* string.TrimSpace(s)  //去除空白符


### reverse 字符串
```
// 反转字符串
func reverseString(s string) string {
    runes := []rune(s)
    for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
        runes[from], runes[to] = runes[to], runes[from]
    }
    return string(runes)
}
```