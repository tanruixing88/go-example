package main

import "fmt"

const  (
	a = iota
	b = iota
)

//iota 按照声明次序顺次自增
const (
	name1 = "menglu"
	name2 = "menglu"
	name = "menglu"
	c    = iota
	d    = iota
)

func main()  {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
