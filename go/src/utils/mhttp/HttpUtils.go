package HttpUtils

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// resp.Cookies() 可以获取返回的cookie
func Get(url0 string) (string, error) {
	resp, err := http.Get(url0)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// Get请求，带上header和cookie
func GetWithHeaderAndCookie(url0 string, header, cookie map[string]string) (string, error) {
	req, err := http.NewRequest("GET", url0, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if header != nil {
		for key, value := range header {
			req.Header.Set(key, value)
		}
	}

	if cookie != nil {
		cookieStr := ""
		for key, value := range cookie {
			cookieStr += key + "=" + value + ";"
		}
		req.Header.Set("cookie", cookieStr)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return string(body), nil
}

func Post(url0 string) (string, error) {
	return PostWithParamMap(url0, map[string]string{})
}

// post form data形式
func PostWithParamMap(url0 string, param map[string]string) (string, error) {
	paramStr := UrlEncodeMap0(param)

	resp, err := http.Post(url0,
		"application/x-www-form-urlencoded",
		strings.NewReader(paramStr))
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// post  form data形式。
func PostWithParamValues(url0 string, values url.Values) (string, error) {
	resp, err := http.PostForm(url0, values)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// post form data形式，支持header, params 不能为空
func PostForm(url string, params map[string]string, header map[string]string) (string, error) {
	paramList := []string{}
	for key, val := range params {
		paramList = append(paramList, key+"="+val)
	}

	req, err := http.NewRequest("POST", url, strings.NewReader(strings.Join(paramList, "&")))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if header != nil {
		for key, value := range header {
			req.Header.Set(key, value)
		}
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return string(body), nil
}

// post原始内容， header中不设置Content-Type，
func PostRawData(url0 string, data string) (string, error) {
	req, err := http.NewRequest("POST", url0, strings.NewReader(data))
	if err != nil {
		return "", err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return string(body), nil
}
