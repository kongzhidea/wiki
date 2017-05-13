package main

import (
	"fmt"
	"utils/mfile"
)

func main() {
	pathname := "/Users/kongzhihui/test"
	fileName := "/Users/kongzhihui/test/test"
	fmt.Println(FileUtils.Exists(fileName))
	fmt.Println(FileUtils.IsFile(fileName))
	fmt.Println(FileUtils.IsDir(fileName))
	fmt.Println(FileUtils.IsDir(pathname))

	fmt.Println(FileUtils.ListDir(pathname))

	fmt.Println(FileUtils.CopyFile(fileName, pathname+"/"+"dest_test"))

	fmt.Println(FileUtils.MoveFile(pathname+"/"+"dest_test", pathname+"/"+"dest_test2"))

	fmt.Println(FileUtils.DeleteFile("/Users/kongzhihui/test/tt/test"))
	fmt.Println(FileUtils.DeleteFile("/Users/kongzhihui/test"))
}
