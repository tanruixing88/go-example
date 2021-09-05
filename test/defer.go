package main

import "fmt"

// defer 压栈时参数已经固定，输出1
func deferFuncParam1() {
	a := 1

	defer fmt.Println(a)

	a = 2
	return
}

func foo1() int {
	i := 0
	defer func() { i++ } ()
	return i
}

//返回值声明不带参数名则并不修改返回值
func foo2() int {
	i := 10
	defer func(i int) {
		fmt.Printf("foo2 defer enter:%d\r\n", i)
		i++
		fmt.Printf("foo2 defer exit:%d\r\n", i)
	} (i)
	i = 4
	return i
}

//返回值声明带参数名同样并不修改返回值
func foo3() (i int) {
	i = 10
	defer func(i int) {
		fmt.Printf("foo3 defer enter:%d\r\n", i)
		i++
		fmt.Printf("foo3 defer exit:%d\r\n", i)
	} (i)
	i = 4
	return i
}

//返回值声明带参数名, defer函数并不锁定i = 10入参，所以会更新参数i = 4 -> i = 5
func foo4() (i int) {
	i = 10
	defer func() {
		fmt.Printf("foo4 defer enter:%d\r\n", i)
		i++
		fmt.Printf("foo4 defer exit:%d\r\n", i)
	} ()
	i = 4
	return i
}

func main() {
	deferFuncParam1()
	r := foo1()
	fmt.Printf("foo1 ret:%d\r\n", r)
	r = foo2()
	fmt.Printf("foo2 ret:%d\r\n", r)
	r = foo3()
	fmt.Printf("foo3 ret:%d\r\n", r)
	r = foo4()
	fmt.Printf("foo4 ret:%d\r\n", r)
}
