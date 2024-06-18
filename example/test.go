package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	fmt.Printf("len(s):%d cap(s):%d\r\n", len(s), cap(s))
	s = append(s, 4)
	fmt.Printf("len(s):%d cap(s):%d\r\n", len(s), cap(s))

}
