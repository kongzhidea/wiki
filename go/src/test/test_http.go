package main

import (
	"fmt"
	"utils/mhttp"
)

func main() {
	{
		url := "http://localhost:9999/json2?id=" + HttpUtils.UrlEncode("控")
		fmt.Println(url)
		fmt.Println(HttpUtils.Get(url))
	}
	{
		url := "http://localhost:9999/json2"
		fmt.Println(HttpUtils.PostWithParamMap(url, map[string]string{"id": "控制0"}))
	}
	{
		url0 := "http://localhost:9999/json2"
		fmt.Println(HttpUtils.PostWithParamValues(url0, HttpUtils.BuildUrlValuesFromMap(map[string]string{"id": "控制1"})))
	}
	{
		url := "http://localhost:9999/postdata"
		fmt.Println(HttpUtils.PostRawData(url, "控制"))
	}
	{
		url := "http://localhost:9999/json2"
		fmt.Println(HttpUtils.GetWithHeaderAndCookie(url,
			map[string]string{"user-agent":"kk","referer":"http://www.kk.com"},
			map[string]string{"token":"123456","t":"45678"}))
	}

}
