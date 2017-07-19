package main

import (
	"fmt"
)

func swap(x, y string) (string, string) {
	return y,x
} 

func main() {
	a, b := swap("hello", "world")  //返回对各值的时候需要用 := 匹配返回值
	fmt.Println(a,b)
}