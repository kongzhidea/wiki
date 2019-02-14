package main

// 安装：go get github.com/satori/go.uuid
import (
	"fmt"
	"github.com/satori/go.uuid"
	"strings"
)

func main() {
	u, err := uuid.NewV4()
	if err == nil {
		fmt.Println(u.String())
	}

	fmt.Println(UUID())
}

func UUID() string {
	u, err := uuid.NewV4()
	if err != nil {
		fmt.Println("generate uuid err:" + err.Error())
	}

	uuid := u.String()
	return strings.Replace(uuid, "-", "", -1)
}
