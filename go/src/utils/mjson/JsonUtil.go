package JsonUtil

import (
	"encoding/json"
	"strconv"
)

func Encode(obj interface{}) (string, error) {
	b, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// json中int会转成 float64类型，需要注意
func DecodeMap(s string) (map[string]interface{}, error) {
	var f map[string]interface{}
	err := json.Unmarshal([]byte(s), &f)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func DecodeArray(s string) ([]interface{}, error) {
	var f []interface{}
	err := json.Unmarshal([]byte(s), &f)
	if err != nil {
		return nil, err
	}
	return f, nil
}

// ------------------------------------------------------------------

// JSON操作
type Js struct {
	Data interface{}
}

// 定义Json; obj := JsonUtil.Json(s)
func Json(data string) Js {
	j := Js{}
	var f interface{}
	err := json.Unmarshal([]byte(data), &f)
	if err != nil {
		return j
	}
	j.Data = f
	return j
}

// key不存在则报错
func (j Js) Get(key string) Js {
	m := j.ToMapData()
	if v, ok := m[key]; ok {
		j.Data = v
		return j
	}
	j.Data = nil
	return j
}

// key不存在则报错
func (j Js) GetString(key string) string {
	v := j.Get(key)
	return v.Data.(string)
}

func toInt(v interface{}) int {
	switch v.(type) {
	case float64:
		return int(v.(float64))
	case int:
		return v.(int)
	case string:
		val := v.(string)
		intval, _ := strconv.Atoi(val)
		return intval
	}
	return 0
}

// key不存在则报错
func (j Js) GetInt(key string) int {
	v := j.Get(key)
	return toInt(v.Data)
}

// key不存在则报错
func (j Js) GetFloat(key string) float64 {
	v := j.Get(key)
	return v.Data.(float64)
}

// key不存在则报错
func (j Js) GetMap(key string) map[string]interface{} {
	v := j.Get(key)
	return v.Data.(map[string]interface{})
}

// key不存在则报错
func (j Js) GetArray(key string) []interface{} {
	v := j.Get(key)
	return v.Data.([]interface{})
}

// key不存在则报错
func (j Js) GetArraySize() int {
	return len(j.Data.([]interface{}))
}

// key不存在则报错
func (j Js) GetArrayIndex(i int) Js {
	obj := j.Data.([]interface{})[i]
	return Js{obj}
}

// key不存在则报错
func (j Js) GetArrayIndexString(i int) string {
	return j.Data.([]interface{})[i].(string)
}

// key不存在则报错
func (j Js) GetArrayIndexFloat(i int) float64 {
	return j.Data.([]interface{})[i].(float64)
}

// key不存在则报错
func (j Js) GetArrayIndexInt(i int) int {
	v := j.Data.([]interface{})[i]

	return toInt(v)
}

//判断是否有郊
func (j Js) IsValid() bool {
	if nil == j.Data {
		return false
	} else {
		return true
	}
}

func (j Js) ContainKey(key string) bool {
	m := j.ToMapData()
	if _, ok := m[key]; ok {
		return true
	}
	return false
}

//return json data
func (j Js) ToMapData() map[string]interface{} {
	if m, ok := (j.Data).(map[string]interface{}); ok {
		return m
	}
	return nil
}

func (j Js) ToArray() []interface{} {
	if m, ok := (j.Data).([]interface{}); ok {
		return m
	}
	return nil
}

func (j Js) ToJsonString() (string, error) {
	return Encode(j.Data)
}

func (j Js) ToString() string {
	return j.Data.(string)
}

func (j Js) ToFloat() float64 {
	return j.Data.(float64)
}


func (j Js) ToInt() int {
	return toInt(j.Data)
}
