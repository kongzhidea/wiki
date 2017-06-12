package main

import (
	"fmt"
	"os/exec"
)

// 执行系统命令
// Run运行命令
// Output 运行命令，并返回结果

func main() {
	{
		// 第一个参数是命令， 后面是参数
		cmd := exec.Command("touch", "/Users/kongzhihui/test/ttt")

		err := cmd.Run()

		fmt.Println(err)
	}

	{

		cmd := exec.Command("pwd")

		out,err := cmd.Output()
		fmt.Println(string(out), err)

	}

}
