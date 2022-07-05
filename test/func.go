package main

import "fmt"

func valCopy(s string) { s = "abc" }

//string类型肯定是值拷贝
//结果输出aaa
func testValCopy() {
	s := "aaa"
	valCopy(s)
	fmt.Printf("s:%s\r\n", s)
}

//空函数是否相等, 此类都是编译错误
func emptyFuncEqual() {
	//fmt.Printf("func is equal:%t\r\n", (func(){} == func(){}))
	//fmt.Println(func(){}== func(){})
}


func main() {
	testValCopy()
}
