package main

import (
	"fmt"
	"utils/mrandon"
	"math/rand"
)

func main()  {
	for i:=0;i<100;i++ {
		fmt.Print(rand.Int31n(20),",")
	}
	fmt.Println()

	for i:=0;i<100;i++ {
		fmt.Print(RandomUtil.RandInt(10,20),",")
	}
	fmt.Println()

	for i:=0;i<100;i++ {
		fmt.Print(RandomUtil.RandFloat(),",")
	}
	fmt.Println()


	for i:=0;i<10;i++ {
		fmt.Print(RandomUtil.RandomNumeric(10),",")
	}
	fmt.Println()

	for i:=0;i<10;i++ {
		fmt.Print(RandomUtil.RandomAlpha(10),",")
	}
	fmt.Println()

	for i:=0;i<10;i++ {
		fmt.Print(RandomUtil.RandomAlphaNumeric(10),",")
	}
	fmt.Println()

}
