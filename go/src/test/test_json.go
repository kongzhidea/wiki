package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"utils/mjson"
)

// 如果是struct类型，字段需要是public类型，即首字母大写，否则无法转成json。
type Response1 struct {
	Page   int
	Fruits []string
}

// 使用tag来自定义编码后JSON键的名称
type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	dict := map[string]interface{}{}
	dict["int"] = 1
	dict["int2"] = "1"
	dict["float"] = 1.2
	dict["str"] = "b"
	dict["cn"] = "中文"
	dict["array"] = []string{"1", "中"}
	dict["bool"] = true

	// encode
	jsonstr, _ := JsonUtil.Encode(dict)
	fmt.Println(jsonstr)
	fmt.Println(JsonUtil.EncodeIgnoreErr(dict))

	fmt.Println(JsonUtil.Encode(Response1{1, []string{"s1", "s2"}}))
	fmt.Println(JsonUtil.Encode(&Response1{1, []string{"s1", "s2"}}))

	fmt.Println(JsonUtil.Encode(Response2{1, []string{"s3", "s4"}}))
	fmt.Println(JsonUtil.Encode(&Response2{1, []string{"s3", "s4"}}))

	// 对象反序列化
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := &Response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res) // &{1 [apple peach]}
	fmt.Println(res.Fruits[0])

	// decode
	s := `{"array":["1","中",{"o":2}],"bool":true,"cn":"中文","float":1.2,"int":1,"int2":"1","str":"b","obj":{"o":1}}`
	dict, err := JsonUtil.DecodeMap(s)
	fmt.Println(dict, err)
	fmt.Println(dict["int"].(float64))
	fmt.Println(dict["float"].(float64))
	fmt.Println(dict["str"].(string))
	fmt.Println(dict["cn"].(string))

	fmt.Println(dict["array"].([]interface{}))
	fmt.Println(dict["array"].([]interface{})[0].(string))

	fmt.Println(dict["bool"].(bool))

	fmt.Println(dict["obj"].(map[string]interface{}))
	fmt.Println(dict["obj"].(map[string]interface{})["o"])

	// jsonobject
	obj := JsonUtil.Json(s)
	fmt.Println(obj.Get("int").Data)           // interface{}
	fmt.Println(obj.Get("int").Data.(float64)) // int--float64
	fmt.Println(obj.ContainKey("str"))
	fmt.Println(obj.GetString("str"))
	fmt.Println(obj.GetString("cn"))
	fmt.Println(obj.GetInt("int"))
	fmt.Println(obj.GetInt("int2"))
	fmt.Println(obj.GetFloat("float"))
	fmt.Println(obj.GetMap("obj"))
	fmt.Println(obj.GetArray("array")[0].(string))

	arr := obj.Get("array")
	fmt.Println(",,,,,", arr.GetArrayIndex(1).Data.(string))
	fmt.Println(",,,,,", arr.GetArrayIndex(1).ToString())
	fmt.Println(",,,,,", arr.GetArrayIndex(0).ToInt())
	fmt.Println(arr.GetArrayIndex(2).ToMapData())
	fmt.Println(arr.GetArrayIndexString(1))
	fmt.Println(arr.GetArrayIndexInt(0))
	fmt.Println(obj.Get("obj").GetInt("o"))

	fmt.Println(obj.ToMapData())
	fmt.Println(arr.ToArray())

	fmt.Println(arr.ToJsonString()) // ["1","中",{"o":2}] <nil>
	fmt.Println(obj.ToJsonString()) //{"array":["1","中",{"o":2}],"bool":true,"cn":"中文","float":1.2,"int":1,"int2":"1","obj":{"o":1},"str":"b"} <nil>

	fmt.Println(reflect.TypeOf(arr)) // JsonUtil.Js

	fmt.Println(obj) // {map[array:[1 中 map[o:2]] bool:true cn:中文 float:1.2 int:1 int2:1 str:b obj:map[o:1]]}
	fmt.Println(arr) // {[1 中 map[o:2]]}
}
