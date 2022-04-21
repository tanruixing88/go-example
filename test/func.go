package main

import "fmt"

func valCopy(s string) { s = "abc" }

func main() {
	s := "aaa"
	valCopy(s)
	fmt.Printf("s:%s\r\n", s)
}
