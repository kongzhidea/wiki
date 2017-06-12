package FileUtils

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

const (
	NEW_FILE_PERM = 0644 // FileMode
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
func Remove(filename string) error {
	return os.Remove(filename)
}

// 删除文件、文件夹，强制删除
func RemoveAll(path string) error {
	return os.RemoveAll(path)
}

// 创建文件夹，不允许级联创建。 path原文件夹存在不会报错，
func Mkdir(path string) error {
	return os.Mkdir(path, os.ModePerm)
}

// 创建文件夹，允许级联创建。  path原文件夹存在不会报错，
func MkdirAll(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

// Read reads the file named by filename and returns the contents.
func Read(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

// ReadString reads the file named by filename and returns the contents as string.
func ReadString(filename string) (string, error) {
	buf, err := Read(filename)
	if err != nil {
		return "", err
	} else {
		return string(buf), nil
	}
}

func ReadLines(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cont []string
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行

		if err != nil || io.EOF == err {
			break
		}
		cont = append(cont, strings.TrimRight(line, "\r\n"))
	}
	return cont, nil
}

// Write writes data to a file named by filename.
func Write(filename string, content []byte) error {
	return ioutil.WriteFile(filename, content, NEW_FILE_PERM)
}

// 会覆盖原文件
func WriteString(filename, content string) error {
	return writeString(filename, content, false)
}

// 向文件中追加内容
func AppendString(filename, content string) error {
	return writeString(filename, content, true)
}

func writeString(filename, content string, append bool) error {
	var flag int
	if append {
		flag = os.O_WRONLY | os.O_CREATE | os.O_APPEND
	} else {
		flag = os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	}
	f, err := os.OpenFile(filename, flag, NEW_FILE_PERM)
	if err != nil {
		return err
	}
	data := []byte(content)
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}

// 同linux touch命令，如果原文件存在则不处理； 要求父目录必须存在
func Touch(filename string) error {
	_, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, NEW_FILE_PERM)
	if err != nil {
		return err
	}
	return err
}

// 会覆盖原文件
func WriteLines(filename string, contents []string) error {
	return writeLines(filename, contents, false)
}

// 向文件中追加内容
func AppendLines(filename string, contents []string) error {
	return writeLines(filename, contents, true)
}

func writeLines(filename string, contents []string, append bool) error {
	var flag int
	if append {
		flag = os.O_WRONLY | os.O_CREATE | os.O_APPEND
	} else {
		flag = os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	}
	f, err := os.OpenFile(filename, flag, NEW_FILE_PERM)
	if err != nil {
		return err
	}
	for _, content := range contents {
		data := []byte(content)
		f.Write(data)
		f.Write([]byte("\n"))
	}

	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}
