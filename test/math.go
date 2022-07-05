package main

import "fmt"

//输出NaN
func divisionFloatZero() {
	var a = 0.0
	const b = 0.0
	fmt.Printf("a/b=%+v\r\n", a/b)
}

func main() {
	divisionFloatZero()
}
