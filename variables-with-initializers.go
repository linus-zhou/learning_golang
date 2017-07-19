package main

import (
	"fmt"
)

var i, j int = 1, 2

func main() {
	//如果你在命名变量时给了初始化的值，那么go会自动给变量确定变量的类型
	var c, python, java = true, false, "no!"  

	fmt.Println(i,j,c, python, java)
}