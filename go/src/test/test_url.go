package main

import (
	"fmt"
	"reflect"
	"utils/mhttp"
	"net/url"
)

func main() {
	{
		query := "a=1&b=2&c=控&a=11&d="
		val, _ := HttpUtils.ParseQuery(query)
		fmt.Println(reflect.TypeOf(val))
		fmt.Println(val)
	}

	{
		query := "a=1&b=2&c=控&a=11&d="
		val, _ := HttpUtils.ParseQuery0(query)
		fmt.Println(reflect.TypeOf(val))
		fmt.Println(val)
	}

	{
		url := "http://localhost:9999/json12?a=1&b=2&c=控&a=11&d="
		ret ,_ := HttpUtils.ParseUrl(url)
		fmt.Println(ret.Query()) // map[a:[1 11] b:[2] c:[控] d:[]]

		fmt.Println(ret.RawQuery)
		fmt.Println(ret.Host)
		fmt.Println(ret.Scheme)

		fmt.Println(HttpUtils.GetQueryByUrl(url))
	}

	{
		param := url.Values{}
		param.Set("a","1")
		param.Set("a","11")
		param.Set("b","控")

		fmt.Println(param.Encode())
	}

	{
		param := map[string][]string{}
		param["a"] = []string{"1"}
		param["b"] = []string{"控制"}
		fmt.Println(HttpUtils.UrlEncodeMap(param))
	}

	{
		param := map[string]string{}
		param["a"] = "1"
		param["b"] = "控制"
		fmt.Println(HttpUtils.UrlEncodeMap0(param))
	}
}
