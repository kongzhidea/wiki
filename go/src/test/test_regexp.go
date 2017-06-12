package main

import (
	"bytes"
	"fmt"
	"regexp"
)

// 正则表达式：http://www.cnblogs.com/golove/p/3269099.html
// 正则表达式：http://studygolang.com/articles/2882

func main() {
	// FindAllString 返回所有匹配项
	// ReplaceAllString 字符串替换
	{
		text := "Hello 世界！123 Go."

		// 查找连续的小写字母
		reg := regexp.MustCompile(`[a-z]+`)
		fmt.Printf("%q\n", reg.FindAllString(text, -1))
		// ["ello" "o"]

		// 查找连续的非小写字母
		reg = regexp.MustCompile(`[^a-z]+`)
		fmt.Printf("%q\n", reg.FindAllString(text, -1))
		// ["H" " 世界！123 G" "."]

		// 查找连续的单词字母
		reg = regexp.MustCompile(`[\w]+`)
		fmt.Printf("%q\n", reg.FindAllString(text, -1))
		// ["Hello" "123" "Go"]

		// 查找连续的汉字
		reg = regexp.MustCompile(`[\p{Han}]+`)
		fmt.Printf("汉字：%q\n", reg.FindAllString(text, -1))
		// ["世界"]

		// 查找连续的非汉字字符
		reg = regexp.MustCompile(`[\P{Han}]+`)
		fmt.Printf("非汉字：%q\n", reg.FindAllString(text, -1))
		// ["Hello " "！123 Go."]

		// 交换 Hello 和 Go
		reg = regexp.MustCompile(`(Hello)(.*)(Go)`)
		fmt.Printf("%q\n", reg.ReplaceAllString(text, "$3$2$1"))
		// "Go 世界！123 Hello."

		// 交换 Hello 和 Go
		reg = regexp.MustCompile(`(Hello)(.*)(Go)`)
		fmt.Printf("%q\n", reg.ReplaceAllString(text, "$3$2$1"))
		// "Go 世界！123 Hello."
	}

	fmt.Println("...........................")

	{
		//这个测试一个字符串是否符合一个表达式。
		match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
		fmt.Println(match)  // true

		//上面我们是直接使用字符串，但是对于一些其他的正则任务，你需要使用 Compile 一个优化的 Regexp 结构体。
		r, _ := regexp.Compile("p([a-z]+)ch")
		//这个结构体有很多方法。这里是类似我们前面看到的一个匹配测试。
		fmt.Println(r.MatchString("peach")) // true


		//这是查找匹配字符串的。
		fmt.Println(r.FindString("peach punch")) // peach

		//这个也是查找第一次匹配的字符串的，但是返回的匹配开始和结束位置索引，而不是匹配的内容。
		fmt.Println(r.FindStringIndex("peach punch")) // [0 5]

		//Submatch 返回完全匹配和局部匹配的字符串。例如，这里会返回 p([a-z]+)ch 和 `([a-z]+) 的信息。
		fmt.Println(r.FindStringSubmatch("peach punch")) // [peach ea]

		//类似的，这个会返回完全匹配和局部匹配的索引位置。
		fmt.Println(r.FindStringSubmatchIndex("peach punch")) // [0 5 1 3]

		//带 All 的这个函数返回所有的匹配项，而不仅仅是首次匹配项。例如查找匹配表达式的所有项。
		fmt.Println(r.FindAllString("peach punch pinch", -1)) // [peach punch pinch]

		//All 同样可以对应到上面的所有函数。
		fmt.Println(r.FindAllStringSubmatchIndex(
			"peach punch pinch", -1)) // [[0 5 1 3] [6 11 7 9] [12 17 13 15]]

		//这个函数提供一个正整数来限制匹配次数。
		fmt.Println(r.FindAllString("peach punch pinch", 2)) // [peach punch]

		//上面的例子中，我们使用了字符串作为参数，并使用了如 MatchString 这样的方法。我们也可以提供 []byte参数并将 String 从函数命中去掉。
		fmt.Println(r.Match([]byte("peach"))) // true

		//创建正则表示式常量时，可以使用 Compile 的变体MustCompile 。因为 Compile 返回两个值，不能用语常量。
		r = regexp.MustCompile("p([a-z]+)ch")
		fmt.Println(r) // p([a-z]+)ch

		//regexp 包也可以用来替换部分字符串为其他值。
		fmt.Println(r.ReplaceAllString("a peach", "<fruit>")) // a <fruit>

		//Func 变量允许传递匹配内容到一个给定的函数中，
		in := []byte("a peach")
		out := r.ReplaceAllFunc(in, bytes.ToUpper)
		fmt.Println(string(out)) // a PEACH

	}
}
