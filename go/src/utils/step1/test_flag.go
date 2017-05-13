package main

import "flag"

/**
命令行flag的语法有如下三种形式：
-flag // 只支持bool类型
-flag=x
-flag x // 只支持非bool类型

--flag
--flag=x
--flag x
 */
func main() {

	var ip = flag.String("ip", "127.0.0.1", "Usage:ip")
	var port = flag.Int("port", 80, "Usage:port")

	var f = flag.Bool("flag", false, "Usag:flag")
	flag.Parse()

	println("ip =", *ip, "port=", *port,"flag=", *f)
}
