package Md5Util

import (
	"encoding/hex"
	"crypto/md5"
)


// 32位小写加密
func Md5(s string)string  {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(s))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}