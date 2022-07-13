package main

import "fmt"

// defer 压栈时参数已经固定，输出1
func deferFuncParam1() {
	a := 1

	defer fmt.Println(a)

	a = 2
	return
}

// 返回结果值是0,因为返回值没有声明为i，是一个匿名变量，在返回前固定为0, 然后再执行defer，虽然将i值改为1
// 但是匿名变量的值还为0，返回结果仍为0
func foo1() int {
	i := 0
	defer func() { i++ } ()
	return i
}

//defer 处理已经固定下来是10,传入了函数后输出10和11,但返回值声明不带参数名，还是一个匿名变量，直接输出为4
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

//返回值声明带参数名i说明返回值没有匿名变量直接为i，虽然后面i值为4,但是全部处理完返回前还得执行一下defer流程
//但是defer执行的是值拷贝，改的不是返回的i，是拷贝的i，故结果还是4
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

// panic 先输出defer倒序，然后输出panic
func deferPanic() {
	defer func() { fmt.Printf("deferPanic 1 \r\n") }()
	defer func() { fmt.Printf("deferPanic 2 \r\n") }()
	defer func() { fmt.Printf("deferPanic 3 \r\n") }()
	panic("panic \r\n")
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
	//deferPanic()
}
