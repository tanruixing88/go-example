package main

import "fmt"

func add_slice(s []int) {

}

func main() {
	s := []int{1, 2, 3}
	fmt.Printf("len(s):%d cap(s):%d", len(s), cap(s))
	add_slice(s)
	s = append(s, 4)
	fmt.Printf("len(s):%d cap(s):%d", len(s), cap(s))

}
