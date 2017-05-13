package SHA1Util

import (
	"encoding/hex"
	"crypto/sha1"
)

func Sha1(s string) string  {
	md5Ctx := sha1.New()
	md5Ctx.Write([]byte(s))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}