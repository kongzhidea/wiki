package mjson_util
//
//import (
//	"github.com/buger/jsonparser"
//	logs "github.com/golang/glog"
//	"strconv"
//)
//
//// 封装jsonparser，如果是string则自动转int
//func getInt64(data []byte, key ...string) (int64, error) {
//	val, err0 := jsonparser.GetInt(data, key...)
//
//	if err0 == nil {
//		return val, nil
//	}
//
//	vals, err := jsonparser.GetString(data, key...)
//
//	if err == nil {
//		val, err := strconv.ParseInt(vals, 10, 64)
//		if err == nil {
//			return val, nil
//		}
//	}
//
//	return 0, err0
//}
//
//// 封装jsonparser，如果是int则自动转string
//func GetString(data []byte, key ...string) (string, error) {
//	val, err0 := jsonparser.GetString(data, key...)
//
//	if err0 == nil {
//		return val, nil
//	}
//
//	vali, err := jsonparser.GetInt(data, key...)
//
//	if err == nil {
//		return strconv.FormatInt(vali, 10), nil
//	}
//
//	return "", err0
//}
//
//// 封装jsonparser， 获取JsonObject对应的字符串
//func GetJsonObject(data []byte, key ...string) (string, error) {
//	extra, t, _, err := jsonparser.Get(data, key...)
//	if t == jsonparser.Object && err == nil {
//		return string(extra), nil
//	}
//	return "{}", err
//}
//
//// 封装jsonparser，返回jsonarray的数组形式。
//func GetArrayString(data []byte, key ...string) ([]string, error) {
//	var res []string
//	f := func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
//		if dataType == jsonparser.String {
//			val, err := jsonparser.ParseString(value)
//			if err != nil {
//				logs.Warn("parse array string error:%v, data=%v", err.Error(), string(data))
//			} else {
//				res = append(res, val)
//			}
//		} else if dataType == jsonparser.Number {
//			val, err := jsonparser.ParseInt(value)
//			if err != nil {
//				logs.Warn("parse array string error:%v, data=%v", err.Error(), string(data))
//			} else {
//				res = append(res, strconv.FormatInt(val, 10))
//			}
//		} else if dataType == jsonparser.Object {
//			res = append(res, string(value))
//		}
//
//	}
//	_, err := jsonparser.ArrayEach(data, f, key...)
//	return res, err
//}
//
//// 封装jsonparser，返回jsonarray的数组形式。
//func GetArrayInt64(data []byte, key ...string) ([]int64, error) {
//	var res []int64
//	f := func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
//		if dataType == jsonparser.String {
//			val, err := jsonparser.ParseString(value)
//			if err != nil {
//				logs.Warn("parse array string error:%v, data=%v", err.Error(), string(data))
//			} else {
//				vali, erri := strconv.ParseInt(val, 10, 64)
//				if erri != nil {
//					logs.Warn("parse array string error:%v, data=%v", erri.Error(), string(data))
//				} else {
//					res = append(res, vali)
//				}
//			}
//		} else if dataType == jsonparser.Number {
//			val, err := jsonparser.ParseInt(value)
//			if err != nil {
//				logs.Warn("parse array string error:%v, data=%v", err.Error(), string(data))
//			} else {
//				res = append(res, val)
//			}
//		}
//
//	}
//	_, err := jsonparser.ArrayEach(data, f, key...)
//	return res, err
//}
