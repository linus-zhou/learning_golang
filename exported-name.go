package main

import (
	"fmt"
	"math"
)
/**
* 在go中包的变量名和方法名首字母是大写就是暴露出来可以被调用的方法
*/
func main() {
	fmt.Println(math.Pi)  //如果你写math.pi是报错的
}