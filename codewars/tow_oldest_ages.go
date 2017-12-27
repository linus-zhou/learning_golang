package main

import "fmt"

func two_oldest_age(ages []int) [2]int {
		lenth := len(ages)
		// 冒泡排序
		for i := 0; i < lenth; i++ {
			for j := 1; j < lenth - i; j++ {
				if ages[j] < ages[j-1] {
					temp := ages[j]
					ages[j] = ages[j-1]
					ages[j-1] = temp
				}
			}
		}
		res := [2]int{ ages[lenth-2], ages[lenth-1]}
		return res
}

func main() {
	ages := []int{6,98,54,3}
	res := two_oldest_age(ages)
	fmt.Printf("%v", res)
}