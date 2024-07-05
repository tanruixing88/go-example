package main

import "fmt"

type P struct {
	Name string
}

func (p *P) String() string {
	return fmt.Sprintf("print: %v", p)
}

/*
 在golang中String() string 方法实际上是实现了String的接口的，该接口定义在fmt/print.go 中：

type Stringer interface {
	String() string
}
在使用 fmt 包中的打印方法时，如果类型实现了这个接口，会直接调用。而题目中打印 p 的时候会直接调用 p 实现的 String() 方法，然后就产生了循环调用。
*/
func RecursiveCall() {
	p := &P{}
	p.String()
}

func main() {
}
