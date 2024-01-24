package main

import "fmt"

const (
	a, b = iota + 10, iota + 12
	e, f = iota, iota + 1
)

//iota 按照声明次序顺次自增
const (
	name1 = "menglu"
	name2 = "menglu"
	name  = "menglu"
	c     = iota
	d     = iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
