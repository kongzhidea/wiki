package main

import (
	"fmt"
	"reflect"
	"utils/list"
)

func main() {

	list := goArrayList.ArrayListNew(10)

	list.Append(1)
	list.Append("3")
	list.Insert(1,"11")
	fmt.Println("list,", list)
	fmt.Println(reflect.TypeOf(list.Get(0)))
	fmt.Println(list.Size())
	fmt.Println(list.Get(1))

	// sublist 返回引用， 修改sublist中元素会同步修改 原list中元素。
	sublist := list.SubList(0, 1)
	fmt.Println("sublist,", sublist)

	sublist.Set(0, 100)
	fmt.Println("sublist2,", sublist)
	fmt.Println("list2,", list)

	var obj goArrayList.Obj = 1
	fmt.Println(obj)

	obj = "100"
	fmt.Println(obj)

	var arr1 = []int{1, 2, 3}
	arr2 := arr1
	arr2[1] = 100
	arr1[0] = 0
	fmt.Println(arr1, arr2) // [0 100 3] [0 100 3]

}
