package main

import (
	"fmt"
)

type Vertex struct{
	Lot, Long float64
}

var m = map[string]Vertex{ //map的使用必须要有键值
	"zhou" : Vertex{
		30.121234,34.566787,
	},
	"zhou1" : Vertex{
		56.12344,78.23456,
	},
}

func main(){
	fmt.Println(m["zhou"])
}