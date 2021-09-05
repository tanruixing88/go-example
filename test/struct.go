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

func main() {
	//tag 的操作
	s := Server{}
	st := reflect.TypeOf(s)
	field1 := st.Field(0)
	fmt.Printf("key1:%v key11:%v\r\n", field1.Tag.Get("key1"), field1.Tag.Get("key11"))
	field2 := st.Field(1)
	fmt.Printf("key2:%v\r\n", field2.Tag.Get("key2"))
}