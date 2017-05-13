package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"reflect"
	"utils/mencode/base64"
	"utils/mencode/md5"
	"utils/mencode/sha1"
)

func main() {
	fmt.Println(Md5Util.Md5("123456"))

	{
		input := []byte("孔1234")

		// 演示base64编码
		encodeString := base64.StdEncoding.EncodeToString(input)
		fmt.Println(encodeString)

		// 对上面的编码结果进行base64解码
		decodeBytes, err := base64.StdEncoding.DecodeString(encodeString)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(string(decodeBytes))

		fmt.Println()

		// 如果要用在url中，需要使用URLEncoding
		uEnc := base64.URLEncoding.EncodeToString([]byte(input))
		fmt.Println(uEnc)

		uDec, err := base64.URLEncoding.DecodeString(uEnc)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(string(uDec))
	}

	{
		fmt.Println(Base64Util.Encode("孔1234"))
		fmt.Println(reflect.TypeOf(Base64Util.Decode("5a2UMTIzNA==")))
		fmt.Println(Base64Util.Decode("5a2UMTIzNA=="))
	}

	{
		fmt.Println(SHA1Util.Sha1("kk"))
	}

}
