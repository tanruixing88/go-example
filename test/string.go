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

type integer int

func (i integer) String() string {
	return "hello"
}

//输出hello
func printInt() {
	fmt.Printf("%s\r\n", integer(5))
}

func longStr() {
	str := `hello
world
v2.0`
	fmt.Println(str)
}

func copyStr() {
	str1 := "1234"
	str2 := str1

	fmt.Printf("str1Addr:%p str2Addr:%p\r\n", &str1, &str2)
}

func main() {
	longStr()
	s := "123456"
	//s[1] = '9' // 编译不过：cannot assign to s[1] (strings are immutable)
	fmt.Printf("s:%s", s)
	string2byteListNoMemCopy()
	printInt()
	copyStr()
}
