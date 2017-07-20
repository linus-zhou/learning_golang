package main

import "fmt" 

func main() {
	p := []int{1,2,3,4,5,12}

	fmt.Println("p ==", p)

	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n", i, p[i])//println是不能打印%d代替变量的
	}
}