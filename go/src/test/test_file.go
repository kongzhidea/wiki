package main

import (
	"fmt"
	"reflect"
	"utils/mfile"
)

// https://github.com/racklin/go-fileutil/blob/master/fileutil.go

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

	fmt.Println(FileUtils.Remove("/Users/kongzhihui/test/tt/test"))
	fmt.Println(FileUtils.Remove("/Users/kongzhihui/test/tt"))

	fmt.Println(FileUtils.MkdirAll("/Users/kongzhihui/test/a/b"))
	//fmt.Println(FileUtils.RemoveAll("/Users/kongzhihui/test/a"))

	cont, _ := FileUtils.ReadString("/Users/kongzhihui/test/cate")
	fmt.Println(cont)
	fmt.Println(reflect.TypeOf(cont))

	conts, _ := FileUtils.ReadLines("/Users/kongzhihui/test/cate")

	for _, cont := range conts{
		fmt.Println(cont)
	}

	FileUtils.WriteString("/Users/kongzhihui/test/c1","你是谁啊？\nhello\n")
	FileUtils.AppendString("/Users/kongzhihui/test/c1","world\n")
	//fmt.Println(FileUtils.Touch("/Users/kongzhihui/test/c2"))

	FileUtils.WriteLines("/Users/kongzhihui/test/clist", conts);
	FileUtils.AppendLines("/Users/kongzhihui/test/clist", conts);
}
