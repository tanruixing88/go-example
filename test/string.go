package main

import (
	"fmt"
	"reflect"
	"unsafe"
)


/*
* 字符串转成切片，会产生拷贝。严格来说，只要是发生类型强转都会发生内存拷贝
*/
//字符串转[]byte类型
func string2byteListNoMemCopy() {
	a := "abc"
	sh := *(*reflect.StringHeader)(unsafe.Pointer(&a))
	byteList := *(*[]byte)(unsafe.Pointer(&sh))
	fmt.Printf("string2byteListNoMemCopy a:%s byteList:%+v\r\n", a, byteList)
}

func main() {
	s := "123456"
	//s[1] = '9' // 编译不过：cannot assign to s[1] (strings are immutable)
	fmt.Printf("s:%s", s)
	string2byteListNoMemCopy()
}
