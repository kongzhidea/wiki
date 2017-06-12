package HttpUtils

import "net/url"

// 返回map，val为数组， 参数不是一个完整的url，只是?后面的参数
func ParseQuery(query string) (map[string][]string, error) {
	return url.ParseQuery(query)
}

// 返回map，value为string， 参数不是一个完整的url，只是?后面的参数
func ParseQuery0(query string) (map[string]string, error) {
	ret, err := url.ParseQuery(query)
	if err != nil {
		return nil, err
	}

	result := map[string]string{}
	for key, value := range ret {
		if len(value) == 0 {
			continue
		}
		result[key] = value[len(value)-1] //同key，取最后一个
	}
	return result, nil
}

// 解析Url， 参数为完成的url，
// ret.Query() 返回 url.Values类型(map[string][]string)
// ret.RawQuery 返回原始参数
// ret.Path 返回路径
// ret.Host 返回域名:端口号
// ret.Scheme 返回http，https
func ParseUrl(u string) (*url.URL, error) {
	return url.Parse(u)
}

// 获取url中的参数，?后面的参数
func GetQueryByUrl(u string) (string, error) {
	ret, err := url.Parse(u)
	if err != nil {
		return "", err
	}
	return ret.RawQuery, nil
}

// url.Values.Encode()  将map(value为string数组)转为 param格式，同时进行html转义。
func UrlEncodeMap(param map[string][]string) string {
	p := url.Values(param)
	return p.Encode()
}

// url.Values.Encode()  将map(value为string)转为 param格式，同时进行html转义。
func UrlEncodeMap0(param map[string]string) string {
	p := url.Values{}
	for key, value := range param {
		p.Set(key, value)
	}
	return p.Encode()
}

func UrlEncode(query string) string {
	return url.QueryEscape(query)
}

func UrlDecode(query string) (string, error) {
	return url.QueryUnescape(query)
}

// map -> url.Values
func BuildUrlValuesFromMap(param map[string]string) url.Values {
	p := url.Values{}
	for key, value := range param {
		p.Set(key, value)
	}
	return p
}