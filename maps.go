package main

import (
	"fmt"
)

type Vertex struct {  //定义结构
	Lat, Long float64
}
var m map[int8]Vertex //定义一个map,[]里面表示其键值类型
//var m map[string]Vertex //定义一个map,[]里面表示其键值类型

func main() {
	m = make(map[int8]Vertex) //map必须使用make函数生成
	//m = make(map[string]Vertex) //map必须使用make函数生成

	m[1] = Vertex{
		40.23234, -63.124343,
	}

	fmt.Println(m[1].Long)
}