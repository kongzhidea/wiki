package FileUtils

import (
	"io"
	"io/ioutil"
	"os"
)

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func Exists(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func IsFile(filename string) bool {
	finfo, err := os.Stat(filename)
	if err != nil {
		return false
	}
	if finfo.IsDir() {
		return false
	} else {
		return true
	}

}
func IsDir(filename string) bool {
	finfo, err := os.Stat(filename)
	if err != nil {
		return false
	}
	if finfo.IsDir() {
		return true
	} else {
		return false
	}

}

// 返回文件和目录，不包含路径，只有名字
func ListDir(pathname string) []string {
	files := []string{}
	dir, err := ioutil.ReadDir(pathname)
	if err != nil {
		return nil
	}
	for _, fi := range dir {
		//if fi.IsDir() { // 忽略目录
		//	continue
		//}
		files = append(files, fi.Name())
	}
	return files
}

// 要求父目录必须存在
//func Touch(filename string) error{
//}


// defer 声明的操作在return之前 逆序执行。在go语言里，记得open一个文件做了错误判断后要紧跟一个defer close延迟调用

// 复制文件，dest如果存在会被覆盖。
func CopyFile(src, dest string) error {
	sp, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sp.Close()

	dp, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer dp.Close()

	_, err = io.Copy(dp, sp)
	if err != nil {
		return err
	}
	return nil
}

// 移动文件
func MoveFile(src, dest string) error {
	return os.Rename(src, dest)
}

// 删除文件、目录（要求目录内不能有文件存在）
func DeleteFile(filename string) error  {
	return os.Remove(filename)
}