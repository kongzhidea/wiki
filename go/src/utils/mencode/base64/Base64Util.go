package Base64Util

import "encoding/base64"

func Encode(s string) string {
	bytes := []byte(s)
	return base64.StdEncoding.EncodeToString(bytes)
}

func Decode(s string) string  {
	decodeBytes, err := base64.StdEncoding.DecodeString(s)
	if err != nil{
		return ""
	}
	return string(decodeBytes)
}