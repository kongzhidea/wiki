package main

import (
       "os"
       "fmt"
)
/**
main参数中包含文件名。

go run args.go  1 2 3

结果：arg_number=4，
 */
func main() {
       arg_num := len(os.Args)
       fmt.Printf("the num of input is %d\n", arg_num)

       fmt.Printf("they are :\n")
       for i := 0; i < arg_num; i++ {
              fmt.Println(os.Args[i])
       }

}