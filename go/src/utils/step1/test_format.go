package main

import "fmt"

func main() {
	s := fmt.Sprintf("i am %s,age is %d", "kk", 20)
	fmt.Println(s) // i am kk,age is 20

	// 使用`%T`来输出一个值的数据类型
	fmt.Printf("%T\n", 1)  // int

	fmt.Printf("%t\n",true)
}
