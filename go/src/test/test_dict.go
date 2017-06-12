package main

import (
	"fmt"
	"utils/listdict"
)

func main() {
	// 初始化 dict
	mat := listdict.NewDict()

	mat.PutIfAbsent("name", "kk")
	mat.Put("id", 2)

	// interface{} 转string类型
	var name string = mat.Get("name", "").(string)

	fmt.Println(name)

	var id int = mat.Get("id", 0).(int)
	fmt.Println(id)

	// 初始化 dict
	d := listdict.Dict{"one": 1, "two": 2}
	d["three"] = 3

	fmt.Println(d)

	// keys
	keys := d.Keys()

	for _, _key := range keys {
		// 转成需要的类型
		var key string = _key.(string)
		fmt.Println(key)
	}
}
