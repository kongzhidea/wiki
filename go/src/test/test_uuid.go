package main


// 安装：go get github.com/nu7hatch/gouuid
import (
	"fmt"
	"github.com/nu7hatch/gouuid"
)

func main() {
	u, err := uuid.NewV4()
	if err == nil {
		fmt.Println(u.String())
		fmt.Println(u.String2())
	}
}
