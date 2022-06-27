package main

import "fmt"

//数组可以进行对比所有元素，此例输出true
func arrayEqual() {
	type pos [2]int
	a := pos{4,5}
	b := pos{4,5}
	fmt.Printf("array a == b:%t", (a==b))
}

func main() {
	arrayEqual()
}
