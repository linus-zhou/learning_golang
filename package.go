package main //main包是程序运行的入口

import (
	"fmt"
	"math/rand" //包名与导入路径的最后一个目录一致，也就是说这里引入的rand包
)

func main() {
	fmt.Println("My favorite number is ", rand.Intn(10))
}