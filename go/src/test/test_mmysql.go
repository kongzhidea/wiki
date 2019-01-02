package main

import (
	"fmt"
	"reflect"
	"utils/mmysql"
	_ "utils/mmysql"
)

func main() {
	dataSource := mmysql.NewDataSourceDef("root", "123456", "127.0.0.1", "kk" )

	client, err := mmysql.Open(dataSource)

	if err != nil {
		fmt.Println(err)
	}

	{
		fmt.Println("QueryAllMap")
		result, err := client.QueryAllMap("select * from user where id < ?", 10)

		if err != nil {
			fmt.Println(err)
		}

		// [map, map]
		client.Show(result)
	}

	{
		fmt.Println("QueryAllValue")
		result, err := client.QueryAllValue("select id from user where id < ?", 10)

		if err != nil {
			fmt.Println(err)
		}

		// [val1, val2]
		client.Show(result)

		item0 := result[0]
		fmt.Println(item0)
		fmt.Println(reflect.TypeOf(item0))
	}

	{
		fmt.Println("QueryColumns")
		result, err := client.QueryColumns("select * from user where id < ?", 10)

		if err != nil {
			fmt.Println(err)
		}

		// 返回查询的列名 [id, username, addtime]
		client.Show(result)
	}

	{
		fmt.Println("QueryMap")
		result, err := client.QueryMap("select * from user where id = ?", 10)

		if err != nil {
			fmt.Println(err)
		}

		// 返回查询的列名 [id, username, addtime]
		client.Show(result)
	}

}
