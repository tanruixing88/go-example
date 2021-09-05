package main

import "fmt"

func main() {
	s := "123456"
	//s[1] = '9' // 编译不过：cannot assign to s[1] (strings are immutable)
	fmt.Printf("s:%s", s)
}
