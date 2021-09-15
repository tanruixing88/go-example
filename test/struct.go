package main

import (
	"fmt"
	"reflect"
)

//tag 使用
type Server struct {
	ServerName string `key1:"value1" key11:"value11"`
	ServerIP   string `key2:"value2"`
}

type AI interface {
	Show()
}

//判断是否为空struct
type A struct {
	name string
	age int
	list []int
	kv  map[int]int
}

func (a A) IsEmpty() bool {
	return reflect.DeepEqual(a, A{})
}

func (a *A) Show() {
	fmt.Printf("name:%s\r\n", a.name)
}

//声明了String方法，直接打印方法返回的信息
/*
func (a A) String() string {
	return fmt.Sprintf("print: A")
}
*/

func main() {
	//tag 的操作
	s := Server{}
	st := reflect.TypeOf(s)
	field1 := st.Field(0)
	fmt.Printf("key1:%v key11:%v\r\n", field1.Tag.Get("key1"), field1.Tag.Get("key11"))
	field2 := st.Field(1)
	fmt.Printf("key2:%v\r\n", field2.Tag.Get("key2"))


	//判断是否为空struct
	var a A
	//invalid operation: a == A{} (struct containing []int cannot be compared)
	/*
	if a == (A{}) {
		fmt.Printf("a == A{} empty\r\n")
	}
	*/

	if a.IsEmpty() {
		fmt.Printf("reflect deep empty\r\n")
	}

	newA := new(A)
	if newA.list == nil {
		fmt.Printf("newA list is nil \r\n")
	} else {
		fmt.Printf("newA list not nil \r\n")
	}

	if newA.kv == nil {
		fmt.Printf("newA kv is nil \r\n")
	} else {
		fmt.Printf("newA kv not nil \r\n")
	}

	fmt.Printf("newA:%+v\r\n", newA)
	newA.list = append(newA.list, 1)
	fmt.Printf("newA:%+v\r\n", newA)
	//newA.kv[1] = 1 // newA.kv是nil，不能直接赋值
	fmt.Printf("newA:%+v\r\n", newA)

	m := map[string]*Server{"server": {"ServerName", "IP"}}
	//要想修改serverName需要定义指针map[string]*Server，因为map的value会扩容会变化，不能直接修改value内部的数据
	m["server"].ServerName = "server"

	var a1 *A
	if a1 == nil {
		fmt.Printf("a1 is nil\r\n")
	} else {
		fmt.Printf("a1 not nil\r\n")
	}

	fmt.Printf("pointer cmp:%t\r\n", (&Server{"server1", "ip1"}) == (&Server{"server1", "ip1"}));
	fmt.Printf("struct  cmp:%t\r\n", (Server{"server1", "ip1"}) == (Server{"server1", "ip1"}));
	fmt.Printf("[...]string{'1'} == [...]string{'1'} cmp:%t\r\n", [...]string{"1"} == [...]string{"1"})
	//切片不能比较//fmt.Printf("[]string{'1'} == []string{'1'} cmp:%t\r\n", []string{"1"} == []string{"1"})
}