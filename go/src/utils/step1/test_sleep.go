package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Second)

	// sleep 时候，写法最好如下：
	time.Sleep(3 * time.Second)

	fmt.Println("end")
}
