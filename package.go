package main //main包是程序运行的入口

/**
*
* 包名与导入路径的最后一个目录一致，也就是说这里引入的rand包
* 当然这里还可以写成多个import语句，但用这中是最好的
*/
import (
	"fmt"
	"math/rand" 
)

func main() {
	fmt.Println("My favorite number is ", rand.Intn(10))
}