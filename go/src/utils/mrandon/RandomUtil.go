package RandomUtil

import (
	"math/rand"
	"strings"
)

// rand.Intn(max) 范围：[0,max)
// rand.Int31() 范围：[0, Integer.MaxValue]
// rand.Int() 范围：[0, Long.MaxValue]

// randInt(min,max) 范围：[min,max]
func RandInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

// 范围为0-1
func RandFloat() float64 {
	return rand.Float64()
}

// 随机n位数字的 字符串
func RandomNumeric(n int) string {
	ret := []string{}
	for i := 0; i < n; i++ {
		c := string(RandInt('0', '9'))
		ret = append(ret, c)
	}
	return strings.Join(ret, "")
}

// 随机n位字母(大小写)的 字符串
func RandomAlpha(n int) string {
	ret := []string{}
	for i := 0; i < n; i++ {
		var c string

		t := rand.Intn(2)
		if t == 0 {
			c = string(RandInt('a', 'z'))
		} else {
			c = string(RandInt('A', 'Z'))
		}

		ret = append(ret, c)
	}
	return strings.Join(ret, "")
}

// 随机n位字母(大小写)或数字的 字符串
func RandomAlphaNumeric(n int) string {
	ret := []string{}
	for i := 0; i < n; i++ {
		var c string

		t := rand.Intn(3)
		if t == 0 {
			c = string(RandInt('0', '9'))
		} else if t == 1 {
			c = string(RandInt('a', 'z'))
		} else {
			c = string(RandInt('A', 'Z'))
		}
		ret = append(ret, c)
	}

	return strings.Join(ret, "")
}
