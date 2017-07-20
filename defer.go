package main

import (
	"fmt"
)

func testDefer() {
	defer fmt.Println("world") //defer语句会在当前函数执行结束返回后执行

	fmt.Println("hello")
}

func main() {
	defer fmt.Println("here is mian")
	testDefer()
}
